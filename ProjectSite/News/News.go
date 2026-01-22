package News

type News struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Categories struct {
	newsid     int `json:"newsid"`
	categoryid int `json:"categoryid"`
}
