package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/smockoro/mysql-master-slave/sample-go/adapter/repository"
	"github.com/smockoro/mysql-master-slave/sample-go/component/config"
)

func main() {
	cfg := config.NewConfig()
	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	ur := repository.NewUserRepository(db)

	ctx := context.Background()
	user, err := ur.Find(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user)

	os.Exit(0)
}
