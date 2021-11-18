package mailjet

type Email struct {
	FromName string `json:"from_name"`
	To       string
	Subject  string
	Content  string
}
