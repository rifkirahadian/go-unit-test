package dtos

type CreateOrderDto struct {
	Name   string `form:"name" binding:"required"`
	Email  string `form:"email" binding:"required"`
	BookId int16  `form:"bookId" binding:"required"`
}
