package controllers

import (
	"book-management-api/commons"
	"book-management-api/internals/models"
	"book-management-api/internals/services"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(bookService services.BookService) *BookController {
	return &BookController{bookService: bookService}
}

func (b *BookController) GetBooks(c *gin.Context) {
	fmt.Println("GetBooks called")
	books, err := b.bookService.GetBooks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commons.ApiErrorResponse("Failed to fetch books", nil))
		return
	}
	fmt.Printf("%d records fetched from database %v\n", len(books), books)
	c.JSON(http.StatusOK, map[string]interface{}{
		"total": len(books),
		"books": books,
	})
}

func (b *BookController) GetBookById(c *gin.Context) {
	bookId := c.Param("id")
	fmt.Printf("GetBookById for id: %s\n", bookId)
	if len(strings.TrimSpace(bookId)) == 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Book ID is required", nil))
		return
	}

	book, err := b.bookService.GetBookById(c, bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commons.ApiErrorResponse("Failed to fetch book", nil))
		return
	}
	c.JSON(http.StatusOK, book)
	fmt.Printf("GetBookById for id: %s, %v\n", bookId, book)
}

func (b *BookController) CreateBook(c *gin.Context) {
	var book *models.BookSchema
	if err := c.ShouldBindJSON(&book); err != nil || book == nil {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Invalid request payload", nil))
		return
	}

	if len(strings.TrimSpace(book.Title)) == 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Title is required", nil))
		return
	}

	if len(strings.TrimSpace(book.Author)) == 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Author is required", nil))
		return
	}

	if book.Year <= 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Year is required and must be valid", nil))
		return
	}

	// Create the book using the service
	bookId, err := b.bookService.CreateBook(c, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, commons.ApiErrorResponse("Failed to create book", nil))
		return
	}
	c.JSON(http.StatusCreated, map[string]string{"id": bookId})
}

func (b *BookController) UpdateBook(c *gin.Context) {
	bookId := c.Param("id")
	if len(strings.TrimSpace(bookId)) == 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Book ID is required", nil))
		return
	}

	var book *models.BookSchema
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Invalid request payload", nil))
		return
	}

	if len(strings.TrimSpace(book.Title)) == 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Title is required", nil))
		return
	}

	if len(strings.TrimSpace(book.Author)) == 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Author is required", nil))
		return
	}

	if book.Year <= 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Year is required and must be valid", nil))
		return
	}

	// Update the book using the service
	if err := b.bookService.UpdateBook(c, book, bookId); err != nil {
		c.JSON(http.StatusInternalServerError, commons.ApiErrorResponse("Failed to update book", nil))
		return
	}
	c.Status(http.StatusOK)
}

func (b *BookController) DeleteBook(c *gin.Context) {
	bookId := c.Param("id")
	if len(strings.TrimSpace(bookId)) == 0 {
		c.JSON(http.StatusBadRequest, commons.ApiErrorResponse("Book ID is required", nil))
		return
	}

	if err := b.bookService.DeleteBookById(c, bookId); err != nil {
		c.JSON(http.StatusInternalServerError, commons.ApiErrorResponse("Failed to delete book", nil))
		return
	}

	c.Status(http.StatusNoContent)
}
