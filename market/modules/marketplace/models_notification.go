package marketplace

import (
	// "fmt"
	"time"

	"argomarket/market/modules/util"

	// "github.com/jinzhu/gorm"
)

type Notification struct {
	Uuid     string     `json:"uuid" gorm:"primary_key"`
	UserUuid string     `json:"user_uuid" gorm:"type:uuid"` // Foreign key field
	User     User       `json:"user" gorm:"foreignkey:UserUuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // Relationship field

	Text      string      `json:"-"`
	Link      string      `json:"-"`
	Read      bool        `json:"status" sql:"index"`
	CreatedAt *time.Time  `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at"`
	DeletedAt *time.Time  `json:"deleted_at"`
}

func CreateNotification(userUuid string, text string, link string) (*Notification) {
	notification := &Notification{
		Uuid:             util.GenerateUuid(),
		UserUuid: userUuid,
		Text: text,
		Link: link,
		Read: false,
	}

	notification.Save()

	return notification
}

/*
	Model Methods
*/

func (n Notification) Save() error {
	// err := n.Validate()
	// if err != nil {
	// 	return err
	// }
	return n.SaveToDatabase()
}

func (n Notification) SaveToDatabase() error {
	if existing, _ := FindNotificationByUuid(n.Uuid); existing == nil {
		return database.Create(&n).Error
	}
	return database.Save(&n).Error
}

func FindNotificationsByUserUuid(UserUuid string) ([]Notification, error) {
	var notifications []Notification

	err := database.
		Where("user_uuid = ?", UserUuid).
		Find(&notifications).
		Error

	if err != nil {
		return nil, err
	}

	return notifications, nil
}


func FindNotificationByUuid(Uuid string) (*Notification, error) {
	var notification Notification

	err := database.
		Where("uuid = ?", Uuid).
		First(&notification).
		Error

	if err != nil {
		return nil, err
	}

	return &notification, nil
}


