package main

import (
	"fmt"
	"image_processing/repository/psql"
)

func main() {
	p := psql.NewPgxPool()
	fmt.Println(p)
}