package structures

type AdminStructs struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	Phone string `json:"phone"`
	OTP string `json:"otp"`
	OTPExpiry string `json:"otpExpiry"`
	APIKey string `json:"APIKey"`
	Uuid string `json:"uuid"`
	Jwt string `json:"jwt"`
	EmailVerified bool `json:"emailVerified"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}