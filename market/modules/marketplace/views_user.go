package marketplace

import (
	"os"
	"io"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	// "time"

	"github.com/gocraft/web"

	"argomarket/market/modules/util"
)

func (c *Context) ViewUserNotificationsGET(w web.ResponseWriter, r *web.Request) {
	notifications, _ := FindNotificationsByUserUuid(c.ViewUser.Uuid)

	c.ViewNotifications = notifications

	util.RenderTemplate(w, "notifications", c)
}

func (c *Context) ViewNotificationDetail(w web.ResponseWriter, r *web.Request) {
	notification, _ := FindNotificationByUuid(r.PathParams["uuid"])

	notification.Read = true
	notification.Save()

	http.Redirect(w, r.Request, notification.Link, 302)
}

func (c *Context) ViewFavoritesGET(w web.ResponseWriter, r *web.Request) {
	secretText := util.GenerateUuid()
	session := sessionStore.Load(r.Request)
	session.PutString(w, "secrettext", secretText)

	c.SecretText = secretText

	util.RenderTemplate(w, "favorites", c)
}

func (c *Context) ViewUserSettingsGET(w web.ResponseWriter, r *web.Request) {
	if len(r.URL.Query()["section"]) > 0 {
		section := r.URL.Query()["section"][0]
		c.SelectedSection = section
	} else {
		c.SelectedSection = "profile"
	}
	secretText := util.GenerateUuid()
	session := sessionStore.Load(r.Request)
	session.PutString(w, "secrettext", secretText)

	c.SecretText = secretText

	// c.UserSettingsHistory = SettingsChangeHistoryByUser(c.ViewUser.User.Uuid)

	util.RenderTemplate(w, "user-settings", c)
}

func (c *Context) ViewPackageDetailGet(w web.ResponseWriter, r *web.Request) {
	package_, err := FindPackageByUuid(r.PathParams["packageUuid"])
	if err != nil {
		return 
	} 
	c.ViewPackage = package_

	util.RenderTemplate(w, "package_detail", c)
}

func (c *Context) ViewUserSettingsPOST(w web.ResponseWriter, r *web.Request) {

	success := false
	// var (
	// 	previousBTCAddress = c.ViewUser.User.Bitcoin
	// 	previousBCHAddress = c.ViewUser.User.BitcoinCash
	// 	previousETHAddress = c.ViewUser.User.Ethereum
	// 	btcAddress         = r.FormValue("bitcoin")
	// 	bchAddress         = r.FormValue("bitcoin_cash")
	// 	ethereumAddress    = r.FormValue("ethereum")
	// )

	// Retrieve the uploaded file
	
	file, handler, err := r.FormFile("avatar")
	if err != nil && err != http.ErrMissingFile {
		c.Error = "Failed to retrieve the uploaded file"
		c.ViewUserSettingsGET(w, r)
		return
	}
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	if err == nil && err != http.ErrMissingFile {
		// Generate a unique filename
		filename := c.ViewUser.User.Uuid
		fmt.Println("filename", handler.Filename)

		// Create a file to store the avatar
		avatarPath := filepath.Join("public/avatars", filename)
		avatarFile, err := os.Create(avatarPath)
		if err != nil {
			c.Error = "Failed to create the avatar file"
			c.ViewUserSettingsGET(w, r)
			return
		}
		defer avatarFile.Close()

		// Save the uploaded file to the avatar file
		_, err = io.Copy(avatarFile, file)
		if err != nil {
			c.Error = "Failed to save the avatar file"
			c.ViewUserSettingsGET(w, r)
			return
		}

		dbPath := strings.Replace(avatarPath, "public", "", 1)
		c.ViewUser.User.Avatar = dbPath
		c.ViewUser.User.Save()
		success = true
	}

	if r.FormValue("description") != "" {
		c.ViewUser.User.Description = r.FormValue("description")
		c.ViewUser.User.Save()
		success = true
	}

	// if btcAddress != "" && !bitcoinRegexp.MatchString(btcAddress) {
	// 	c.Error = "Wrong Bitcoin Address"
	// } else {
	// 	c.ViewUser.User.Bitcoin = btcAddress
	// }

	// if bchAddress != "" && !bitcoinRegexp.MatchString(bchAddress) {
	// 	c.Error = "Wrong Bitcoin Cash Address"
	// } else {
	// 	c.ViewUser.User.BitcoinCash = bchAddress
	// }

	// if ethereumAddress != "" && !ethereumRegexp.MatchString(ethereumAddress) {
	// 	c.Error = "Wrong Ethereum"
	// } else {
	// 	c.ViewUser.User.Ethereum = ethereumAddress
	// }

	// if r.FormValue("bitcoin_multisig") != "" {
	// 	c.ViewUser.User.BitcoinMultisigPublicKey = r.FormValue("bitcoin_multisig")
	// 	success = true
	// }
	// if r.FormValue("bitmessage") != "" {
	// 	c.ViewUser.User.Bitmessage = r.FormValue("bitmessage")
	// 	success = true
	// }
	// if r.FormValue("tox") != "" {
	// 	c.ViewUser.User.Tox = r.FormValue("tox")
	// 	success = true
	// }
	// if r.FormValue("email") != "" {
	// 	c.ViewUser.User.Email = r.FormValue("email")
	// 	success = true
	// }
	// if r.FormValue("2fa") != "" {
	// 	success = true
	// 	if r.FormValue("2fa") == "1" {
	// 		c.ViewUser.User.TwoFactorAuthentication = true
	// 	} else if r.FormValue("2fa") == "0" {
	// 		c.ViewUser.User.TwoFactorAuthentication = false
	// 	}
	// }

	if r.FormValue("password") != "" {
		oldPassword := r.FormValue("old_password")
		hashV1 := util.PasswordHashV1(c.ViewUser.User.Username, oldPassword)
		if c.ViewUser.User.PasswordHash != hashV1 {
			c.Error = "Invalid worng password"
			c.ViewUserSettingsGET(w, r)
			return
		}

		newPassword := r.FormValue("password")
		repeatNewPassword := r.FormValue("confirm_password")

		if newPassword != repeatNewPassword {
			c.Error = "New password and repeat new password does not match"
			c.ViewUserSettingsGET(w, r)
			return
		}

		newHash := util.PasswordHashV1(c.ViewUser.User.Username, newPassword)
		c.ViewUser.User.PasswordHash = newHash
		c.ViewUser.User.Save()
		success = true
	}
	if validationError := c.ViewUser.User.Validate(); validationError != nil {
		c.Error = validationError.Error()
		c.ViewUserSettingsGET(w, r)
		return
	}

	// avatarError := util.SaveImage(r, "avatar_image", 300, c.ViewUser.User.Uuid+"_av")
	// if avatarError == nil {
	// 	c.ViewUser.User.HasAvatar = true
	// 	success = true
	// }

	c.ViewUser.User.Save()

	// if previousBTCAddress != c.ViewUser.User.Bitcoin {
	// 	historyEvent := UserSettingsHistory{
	// 		UserUuid: c.ViewUser.User.Uuid,
	// 		Action:   "Bitcoin address changed to " + c.ViewUser.User.Bitcoin,
	// 		Datetime: time.Now(),
	// 		Type:     "bitcoin",
	// 	}
	// 	if c.ViewUser.User.Bitcoin == "" {
	// 		historyEvent.Action = "Bitcoin address deleted"
	// 	}
	// 	historyEvent.Save()
	// }

	// if previousBCHAddress != c.ViewUser.User.BitcoinCash {
	// 	historyEvent := UserSettingsHistory{
	// 		UserUuid: c.ViewUser.User.Uuid,
	// 		Action:   "BitcoinCash address changed to " + c.ViewUser.User.BitcoinCash,
	// 		Datetime: time.Now(),
	// 		Type:     "bitcoin_cash",
	// 	}
	// 	if c.ViewUser.User.BitcoinCash == "" {
	// 		historyEvent.Action = "Bitcoin Cash address deleted"
	// 	}
	// 	historyEvent.Save()
	// }

	// if previousETHAddress != c.ViewUser.User.Ethereum {
	// 	historyEvent := UserSettingsHistory{
	// 		UserUuid: c.ViewUser.User.Uuid,
	// 		Action:   "Ethereum address changed to " + c.ViewUser.User.Ethereum,
	// 		Datetime: time.Now(),
	// 		Type:     "ethereum",
	// 	}
	// 	if c.ViewUser.User.Ethereum == "" {
	// 		historyEvent.Action = "Ethereum address deleted"
	// 	}
	// 	historyEvent.Save()
	// }

	if(!success) {
		c.Error = "No changes"
	} else {
		c.Success = "You have successfully changed your settings"
	}
	c.ViewUserSettingsGET(w, r)
}

func (c *Context) ViewUserProfileGET(w web.ResponseWriter, r *web.Request) {
	user, err := FindUserByUsername(r.PathParams["username"])
	if(err != nil) {
		http.NotFound(w, r.Request)
		return
	}
	c.ViewProfileUser = user
	
	util.RenderTemplate(w, "profile", c)
}

func (c *Context) ViewUserOrdersGET(w web.ResponseWriter, r *web.Request) {

	util.RenderTemplate(w, "orders", c)
}
