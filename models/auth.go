package models

type Auth struct {
	Email 				 string `json:"email" binding:"omitempty,email"`
	Password 			 string `json:"password" binding:"omitempty,gt=6"`
	ConfirmationPassword string `json:"confirmation_password" binding:"omitempty,gt=6"`
	Name 				 string `json:"name" binding:"omitempty"`
}