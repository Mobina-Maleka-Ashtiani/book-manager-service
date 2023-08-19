package BusinessLogic

type AuthorRequestAndResponse struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Birthday    string `json:"birthday"`
	Nationality string `json:"nationality"`
}

type BookRequestAndResponse struct {
	Name            string                   `json:"name"`
	Author          AuthorRequestAndResponse `json:"author"`
	Category        string                   `json:"category"`
	Volume          int                      `json:"volume"`
	PublishedAt     string                   `json:"published_at"`
	Summary         string                   `json:"summary"`
	TableOfContents []string                 `json:"table_of_contents"`
	Publisher       string                   `json:"publisher"`
}
