package domain

type CreatePasswordRequest struct {
	Url      string `json:"url"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
