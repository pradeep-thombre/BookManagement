package configs

import (
	"book-management-api/commons/appdb"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	AppConfig *ApplicationConfig
)

type ApplicationConfig struct {
	HttpPort string
	DbClient appdb.DatabaseClient
}

func NewApplicationConfig(context context.Context) error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv(MONGO_URI)).SetServerAPIOptions(serverAPI)
	client, cerror := mongo.Connect(context, opts)
	if cerror != nil {
		fmt.Println("Error while connecting db, error: ", cerror)
		panic(cerror)
	}

	fmt.Println("You successfully connected to MongoDB!")
	dbClient := appdb.NewDatabaseClient(os.Getenv(MONGO_DATABASE), client)
	AppConfig = &ApplicationConfig{
		HttpPort: os.Getenv(HTTP_PORT),
		DbClient: dbClient,
	}
	return nil
}
