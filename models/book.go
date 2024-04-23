package models

type Book struct {
	Title      string `json:"title"`
	Writer     string `json:"writer"`
	CoverImage string `json:"cover_image"`
	Point      int64  `json:"point"`
	Tag        string `json:"tag"`
}
