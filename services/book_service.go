package services

// Book contains an author, a block of text and a title
type Book struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Title  string `json:"title"`
}

// BookService is type responsible for Book CRUD
type BookService struct{}

// Get returns a Book
func (service *BookService) Get() *Book {
	book := Book{
		Author: "TEST AUTHOR",
		Text:   "TEST TEXT",
		Title:  "TEST TITLE",
	}
	return &book
}
