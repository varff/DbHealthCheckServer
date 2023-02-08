package repo

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Server/internal/settings"

	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	PingMongo() error
	PingPostgres() error
}

func (db *Repository) PingMongo() error {
	err := db.mongo.Ping(context.Background(), nil)
	return err
}

func (db *Repository) PingPostgres() error {
	err := db.postgres.Ping()
	return err
}

type Repository struct {
	postgres *sqlx.DB
	mongo    *mongo.Client
}

func NewRepository(dbSetting *settings.DBSetting) (*Repository, error) {
	connStr := fmt.Sprintf("user=" + dbSetting.DBUser + " password=" +
		dbSetting.DBPassword + " host=" + dbSetting.DBHost + " port=" +
		dbSetting.DBPort + " database=" + dbSetting.DBName)
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	err = godotenv.Load("configs/mongo.env")
	if err != nil {
		return nil, err
	}
	user, err := settings.GetEnvDefault("MONGO_INITDB_ROOT_USERNAME", "root")
	if err != nil {
		return nil, err
	}
	pass, err := settings.GetEnvDefault("MONGO_INITDB_ROOT_PASSWORD", "")
	if err != nil {
		return nil, err
	}
	credential := options.Credential{
		Username: user,
		Password: pass,
	}
	mongoHost, err := settings.GetEnvDefault("MONGO_HOST", "localhost")
	if err != nil {
		return nil, err
	}
	mongoPort, err := settings.GetEnvDefault("MONGO_PORT", "27017")
	if err != nil {
		return nil, err
	}
	uri := fmt.Sprintf("mongodb://%s:%s", mongoHost, mongoPort)
	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}
	return &Repository{postgres: db, mongo: client}, nil
}

func (db *Repository) StopRepository() error {
	err := db.postgres.Close()
	if err != nil {
		return err
	}
	err = db.mongo.Disconnect(context.Background())
	return nil
}
