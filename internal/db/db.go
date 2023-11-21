package db

import (
	"fmt"
	"github.com/aliakseizhyrauliou/gRPCApiGo/internal/rocket"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func New() (Store, error) {

	dbUsername := /*os.Getenv("DB_USERNAME")*/ "postgres"
	dbPassword := /*os.Getenv("DB_PASSWORD")*/ "postgres"
	dbHost := /*os.Getenv("DB_HOST")*/ "localhost"
	dbTable := /*os.Getenv("DB_TABLE")*/ "postgres"
	dbPort := /*os.Getenv("DB_PORT")*/ "5432"
	dbSSLMode := /*os.Getenv("DB_SSL_MODE")*/ "disable"

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
	err := s.db.Get(&rkt, "SELECT * FROM rockets WHERE id=$1", id)
	if err != nil {
		return rocket.Rocket{}, err
	}

	return rkt, nil
}

func (s Store) InsertRocket(rkt rocket.Rocket) (rocket.Rocket, error) {
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

func (s Store) DeleteRocket(id string) error {
	_, err := s.db.Exec("DELETE FROM rockets WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
