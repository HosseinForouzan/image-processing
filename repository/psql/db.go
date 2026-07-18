package psql

import (
	"context"
	"log"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Username string	`koanf:"username"`
	Password string	`koanf:"password"`
	Port int	`koanf:"port"`
	Host string	`koanf:"host"`
	DBName string	`koanf:"db_name"`
} 

type PsqlDB struct {
	config Config
	db *pgx.Pool
}

func (p *PsqlDB) Conn() *pgx.Pool {
	return p.db
}


func NewPgxPool() *pgx.Pool {
	ctx := context.Background()
	// urlExample := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)
	urlExample := "postgres://myuser:secret@localhost:5431/image_db"
	db, err := pgx.New(ctx, urlExample)
	if err != nil {
		log.Fatal(err)
	}


	err = db.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("db connected")

	return db
}