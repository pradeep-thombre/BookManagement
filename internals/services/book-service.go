package services

import (
	"book-management-api/internals/dbservice"
	"book-management-api/internals/models"
	"context"
	"encoding/json"
	"fmt"
)

type BookService interface {
	GetBookById(ctx context.Context, bookId string) (*models.BookSchema, error)
	DeleteBookById(ctx context.Context, bookId string) error
	GetBooks(ctx context.Context) ([]*models.BookSchema, error)
	CreateBook(ctx context.Context, book *models.BookSchema) (string, error)
	UpdateBook(ctx context.Context, book *models.BookSchema, bookId string) error
}

type bookService struct {
	dbservice dbservice.DbService
}

func NewBookService(dbservice dbservice.DbService) BookService {
	return &bookService{dbservice: dbservice}
}

func (s *bookService) GetBookById(ctx context.Context, bookId string) (*models.BookSchema, error) {
	fmt.Println("GetBookById called")
	bookSchema, err := s.dbservice.GetBookById(ctx, bookId)
	if err != nil {
		fmt.Println("Error fetching book:", err)
		return nil, err
	}
	return bookSchema, nil
}

func (s *bookService) GetBooks(ctx context.Context) ([]*models.BookSchema, error) {
	fmt.Println("GetBooks called")
	books, err := s.dbservice.GetBooks(ctx)
	if err != nil {
		fmt.Println("Error fetching books:", err)
		return nil, err
	}

	return books, nil
}

func (s *bookService) DeleteBookById(ctx context.Context, bookId string) error {
	fmt.Println("DeleteBookById called")
	if err := s.dbservice.DeleteBookById(ctx, bookId); err != nil {
		fmt.Println("Error deleting book:", err)
		return err
	}
	return nil
}

func (s *bookService) CreateBook(ctx context.Context, book *models.BookSchema) (string, error) {
	fmt.Println("CreateBook called")
	var bookSchema *models.BookSchema
	pbyes, _ := json.Marshal(book)
	err := json.Unmarshal(pbyes, &bookSchema)
	if err != nil {
		fmt.Println("Error marshalling/unmarshalling book:", err)
		return "", err
	}
	bookId, err := s.dbservice.SaveBook(ctx, bookSchema)
	if err != nil {
		fmt.Println("Error saving book:", err)
		return "", err
	}
	return bookId, nil
}

func (s *bookService) UpdateBook(ctx context.Context, book *models.BookSchema, bookId string) error {
	fmt.Println("UpdateBook called")
	var bookSchema models.BookSchema
	pbyes, _ := json.Marshal(book)
	err := json.Unmarshal(pbyes, &bookSchema)
	if err != nil {
		fmt.Println("Error marshalling/unmarshalling book:", err)
		return err
	}
	return s.dbservice.UpdateBook(ctx, &bookSchema, bookId)
}
