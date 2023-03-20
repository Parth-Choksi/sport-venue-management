package models

type UserDetails struct {
	User_Id    string `json:"user_id"`
	Name       string `json:"name" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password" validate: "required"`
	Contact_No string `json:"contact_no" validate: "required"`
}
