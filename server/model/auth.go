package model

type Auth struct {
	ID        int    `json:id,omitempty`
	Uuid      string `json:"uuid",omitempty`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type User struct {
	Uuid      string `json:"uuid",omitempty`
	Email     string `json:"email,omitempty"`
	Password  string `json:-,omitempty`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
