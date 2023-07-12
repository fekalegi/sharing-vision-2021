package model

type AddArticle struct {
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content"`
	Category string `json:"category"`
	Status   string `json:"status"`
}
