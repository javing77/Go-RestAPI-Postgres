package main

import (
	"context"
	"fmt"

	"github.com/javing77/go-rest-postgress/internal/comment"
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

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "2aecc170-04a9-11ee-be56-0242ac120002",
			Slug:   "Manual-test",
			Author: "Javier",
			Body:   "Insert test",
		},
	)
	fmt.Println(cmtService.GetComment(
		context.Background(),
		"2aecc170-04a9-11ee-be56-0242ac120002",
	))
	return nil
}

func main() {
	fmt.Println("Go Rest App")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
