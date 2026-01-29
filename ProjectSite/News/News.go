package News

type News struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Categories struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type GetResults struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Categories string `json:"Categories"`
}
