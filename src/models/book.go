package models

type Book struct {
	ID         uint    `json:"id"`
	Title      string  `json:"title"`
	Writer     string  `json:"writer"`
	CoverImage string  `json:"cover_image"`
	Point      float64 `json:"point"`
	Tag        string  `json:"tag"`
}
