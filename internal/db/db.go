package db

import (
	"fmt"
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/rocket"
	"github.com/jmoiron/sqlx"
	"os"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		dbPort,
		dbUsername,
		dbTable,
		dbPassword,
		dbSSLMode,
	)

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		return Store{}, err
	}

	return Store{
		db: db,
	}, nil
}

func (s Store) GetRocketByID(id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	err := s.db.Get(&rkt, "SELECT * FROM rockets WHERE id=?", id)
	if err != nil {
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
	_, err := s.db.NamedExec(
		`INSERT INTO rockets (id, name, type) VALUES (:id, :name, :type)`,
		rkt,
	)
	if err != nil {
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

func (s Store) DeleteRocket(id string) error {
	_, err := s.db.Exec("DELETE FROM rockets WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
