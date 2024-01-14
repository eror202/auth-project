package response

type Otp struct {
	SecretCode  string `json:"secretCode"`
	OtpInBase64 string `json:"otpInBase64"`
}
