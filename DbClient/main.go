package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {

	dbUrl := os.Getenv("db_url")
	dbURI := fmt.Sprintf("host=%s user=CacheClient dbname=postgres sslmode=disable password=ChangeIt2", dbUrl)
	fmt.Println(dbURI)
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return "err", err
	}

	db.AutoMigrate(&Entity{})
	var ent = &Entity{}
	ent.Text = name.Name
	db.Save(&ent)

	return fmt.Sprint(&ent.ID), nil
}

func main() {
	lambda.Start(HandleRequest)
}
