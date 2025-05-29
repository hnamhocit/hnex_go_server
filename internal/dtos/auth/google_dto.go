package dtos

type GoogleAuthRequest struct {
	IDToken string `json:"id_token"`
}

type GoogleTokenInfo struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Aud           string `json:"aud"`
	Exp           int64  `json:"exp"`
}
