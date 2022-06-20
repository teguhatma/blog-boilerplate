package response

type EntryResponse struct {
	Owner    string `json:"owner"`
	TagName  string `json:"tag_name"`
	Blog     string `json:"blog"`
	Title    string `json:"title"`
	ReadTime string `json:"read_time"`
}
