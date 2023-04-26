package response

type RegisterUser struct {
	FullName string `json:"full_name" form:"full_name" valid:"required~Your full name is required" example:"Alicia Keys"`
	Email    string `json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format" example:"alicia@mail.com"`
	Password string `json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters" example:"Password123"`
}

type RegisterUserResult struct {
	ID       uint   `json:"id"  example:"Alicia Keys"`
	FullName string `json:"full_name" example:"alicia@mail.com"`
	Email    string `json:"email" example:"Password123"`
}
