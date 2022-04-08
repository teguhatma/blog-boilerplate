package response

import "time"

type TagResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
