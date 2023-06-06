package main

import (
	"fmt"

	"github.com/javing77/go-rest-postgress/internal/db"
)

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func Run() error {
	fmt.Println("Starting up out application")

	db, err := db.NewDataBase()
	if err != nil {
		fmt.Println("failed to connect to database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	fmt.Println("succefully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Go Rest App")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
