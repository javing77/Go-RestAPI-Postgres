package main

import (
	"fmt"

	"github.com/javing77/go-rest-postgress/internal/comment"
	"github.com/javing77/go-rest-postgress/internal/db"
	transportHttp "github.com/javing77/go-rest-postgress/internal/transport/http"
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

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go Rest App")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
