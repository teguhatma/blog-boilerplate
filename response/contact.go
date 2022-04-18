package response

import (
	"time"
)

type ContactResponse struct {
	ID        int       `json:"id"`
	Owner     string    `json:"owner"`
	Github    string    `json:"github"`
	Twitter   string    `json:"twitter"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
