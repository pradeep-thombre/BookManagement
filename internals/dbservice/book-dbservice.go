package dbservice

import (
	"book-management-api/commons/appdb"
	"book-management-api/configs"
	"book-management-api/internals/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbService struct {
	collection appdb.DatabaseCollection
}

type DbService interface {
	GetBookById(ctx context.Context, bookId string) (*models.BookSchema, error)
	SaveBook(ctx context.Context, book *models.BookSchema) (string, error)
	UpdateBook(ctx context.Context, book *models.BookSchema, bookId string) error
	DeleteBookById(ctx context.Context, bookId string) error
	GetBooks(ctx context.Context) ([]*models.BookSchema, error)
}

func NewDbService(dbclient appdb.DatabaseClient) DbService {
	return &dbService{
		collection: dbclient.Collection(configs.MONGO_BOOK_COLLECTION),
	}
}

func (d *dbService) GetBookById(ctx context.Context, bookId string) (*models.BookSchema, error) {
	fmt.Println("GetBookById called")
	var book models.BookSchema
	id, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return nil, fmt.Errorf("invalid bookId: %v", err)
	}
	err = d.collection.FindOne(ctx, bson.M{"_id": id}, &book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("book not found: %v", err)
		}
		return nil, err
	}
	return &book, nil
}

func (d *dbService) GetBooks(ctx context.Context) ([]*models.BookSchema, error) {
	fmt.Println("GetBooks called")
	var books []*models.BookSchema
	err := d.collection.Find(ctx, bson.M{}, &options.FindOptions{}, &books)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch books: %v", err)
	}

	return books, nil
}

func (d *dbService) SaveBook(ctx context.Context, book *models.BookSchema) (string, error) {
	fmt.Println("SaveBook called")
	book.ID = primitive.NewObjectID()
	currentTime := time.Now()
	book.CreatedAt = currentTime
	book.UpdatedAt = currentTime

	_, err := d.collection.InsertOne(ctx, book)
	if err != nil {
		return "", err
	}
	return book.ID.Hex(), nil
}

func (d *dbService) UpdateBook(ctx context.Context, book *models.BookSchema, bookId string) error {
	fmt.Println("UpdateBook called")
	id, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return fmt.Errorf("invalid bookId: %v", err)
	}

	update := bson.M{
		"$set": bson.M{
			"title":     book.Title,
			"author":    book.Author,
			"year":      book.Year,
			"updatedAt": time.Now(),
		},
	}

	_, err = d.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return err
	}
	return nil
}

func (d *dbService) DeleteBookById(ctx context.Context, bookId string) error {
	fmt.Println("DeleteBookById called")
	id, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return fmt.Errorf("invalid bookId: %v", err)
	}

	_, err = d.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}
