package main

import (
	"crud/pkg/database"
	"crud/pkg/handlers"
	"crud/pkg/repos"
	"crud/pkg/usecase"
	"log"
	"os"
)

func main() {

	collection, err := database.ConnectToDB(os.Getenv("MONGO"))

	if err != nil {
		panic(err)
	}

	repository := repos.NewRepository(collection)

	useCase := usecase.NewUsecase(repository)

	router := handlers.NewRouter(useCase)

	router.Run()
	log.Println("server started in 8080")
}
