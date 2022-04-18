package request

type ContactRequest struct {
	Owner   string `json:"owner"`
	Github  string `json:"github"`
	Twitter string `json:"twitter"`
}
