package marketplace

import (
	// "fmt"
	"github.com/gocraft/web"
	"net/http"

	"argomarket/market/modules/util"
)

func (c *Context) Index(w web.ResponseWriter, r *web.Request) {
	redirectUrl := "/marketplace"
	http.Redirect(w, r.Request, redirectUrl, 302)
}

func (c *Context) ListSerpItems(w web.ResponseWriter, r *web.Request) {
	if(r.PathParams["location"] != "") {
		c.ViewUser.User.Location = r.PathParams["location"]
		c.ViewUser.User.Save()
	}

	c.ViewPackages = GetAllPackages()
	util.RenderTemplate(w, "home", c)
}