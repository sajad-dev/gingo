package validation

type Login struct {
	Email    string `json:"email" validate:"required,email,max=50"`
	Password string `json:"password" validate:"required,max=100"`
}
type Register struct {
	Email           string `json:"email" validate:"required,uniq=accounts:email,email,max=50"`
	Mobile          string `json:"mobile" validate:"max=20,uniq=accounts:mobile"`
	FullName        string `json:"full_name" validate:"max=100"`
	Password        string `json:"password" validate:"required,max=50"`
	PasswordConfirm string `json:"password_confirm" validate:"required,max=50,eqfield=Password"`
}
