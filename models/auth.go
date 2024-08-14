package models

type RegisterReq struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	Role        string `json:"role"`
}

type RegisterResp struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Gender      string `json:"gender"`
	CreatedAt   string `json:"created_at"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	ExpiresAt    string `json:"expires_at"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type ChangePasswordReq struct {
	Email           string `json:"email"`
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type ForgotPasswordReq struct {
	Email string `json:"email"`
}

type ResetPasswordReq struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

type UserDetails struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
