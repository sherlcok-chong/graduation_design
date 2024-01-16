package v1

import (
	"GraduationDesign/src/api/v1/chat"
)

type group struct {
	User        user
	Email       email
	File        file
	Account     account
	Application application
	Notify      notify
	Setting     setting
	MGroup      mGroup
	Message     message
	Chat        chat.Group
}

var Group = new(group)
