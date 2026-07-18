package psqlimage

import "image_processing/repository/psql"

type DB struct {
	conn *psql.PsqlDB
}

func New(conn *psql.PsqlDB) *DB {
	return &DB{conn: conn}
}