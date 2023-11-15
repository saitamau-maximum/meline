package entity

type OAuthUserResponse struct {
	OAuthUserID   string  `json:"user_id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}
