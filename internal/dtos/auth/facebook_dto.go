package dtos

type FacebookAuthRequest struct {
	AccessToken string `json:"accessToken" binding:"required"`
}

type FacebookDebugToken struct {
	Data struct {
		AppID     string `json:"app_id"`
		IsValid   bool   `json:"is_valid"`
		UserID    string `json:"user_id"`
		ExpiresAt int64  `json:"expires_at"`
	} `json:"data"`
}

type FacebookUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	} `json:"picture"`
}
