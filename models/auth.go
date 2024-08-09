package models

type RegisterReq struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Fullname       string `json:"full_name"`
	NativeLanguage string `json:"native_language"`
	Role           string `json:"role"`
}

type RegisterResp struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	Fullname       string `json:"full_name"`
	NativeLanguage string `json:"native_language"`
	CreatedAt      string `json:"created_at"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	ExpiresAt    string `json:"expires_at"` // expiry date of access token
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
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
