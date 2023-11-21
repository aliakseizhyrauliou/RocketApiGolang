package db

import (
	"fmt"
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/rocket"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
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

func (s Store) GetRocketByID(ctx context.Context, id string) (rocket.Rocket, error) {
	var rkt rocket.Rocket
	err := s.db.Get(&rkt, "SELECT * FROM rockets WHERE id=$1", id)
	if err != nil {
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

func (s Store) InsertRocket(ctx context.Context, rkt rocket.Rocket) (rocket.Rocket, error) {
	rows, err := s.db.NamedQuery(
		`INSERT INTO rockets (name, type) VALUES (:name, :type) RETURNING id`,
		rkt,
	)

	if err != nil {
		return rocket.Rocket{}, err
	}

	var newId string

	if rows.Next() {
		err := rows.Scan(&newId)

		if err != nil {
			return rocket.Rocket{}, err
		}
	}

	rkt.ID = newId

	return rkt, nil
}

func (s Store) DeleteRocket(ctx context.Context, id string) error {
	_, err := s.db.Exec("DELETE FROM rockets WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s Store) GetRocketList(ctx context.Context, take int32, skip int32) ([]rocket.Rocket, error) {

	var rockets []rocket.Rocket

	err := s.db.Select(&rockets, "SELECT * FROM rockets ORDER BY id LIMIT $1 OFFSET $2", take, skip)

	if err != nil {
		return nil, err
	}

	return rockets, nil
}
