package BusinessLogic

import (
	"book-manager-service/DataAccess"
	"strings"
)

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

func ConvertBookRequestToBook(br BookRequestAndResponse) DataAccess.Book {
	return DataAccess.Book{
		Name: br.Name,
		Author: DataAccess.Author{
			FirstName:   br.Author.FirstName,
			LastName:    br.Author.LastName,
			Birthday:    br.Author.Birthday,
			Nationality: br.Author.Nationality,
		},
		Category:        br.Category,
		Volume:          br.Volume,
		PublishedAt:     br.PublishedAt,
		Summary:         br.Summary,
		TableOfContents: strings.Join(br.TableOfContents[:], ","),
		Publisher:       br.Publisher,
	}
}
func ConvertBookToBookResponse(book DataAccess.Book) BookRequestAndResponse {
	return BookRequestAndResponse{
		Name: book.Name,
		Author: AuthorRequestAndResponse{
			FirstName:   book.Author.FirstName,
			LastName:    book.Author.LastName,
			Birthday:    book.Author.Birthday,
			Nationality: book.Author.Nationality,
		},
		Category:        book.Category,
		Volume:          book.Volume,
		PublishedAt:     book.PublishedAt,
		Summary:         book.Summary,
		TableOfContents: strings.Split(book.TableOfContents, ","),
		Publisher:       book.Publisher,
	}
}
