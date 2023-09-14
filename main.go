package main

import (
	"blogapi/models"
	"blogapi/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//docker run --name=kubecontrollermysql -e MYSQL_DATABASE=kubecontrollerdb -e MYSQL_USER=kubeuser -e MYSQL_PASSWORD=xrock@971768 -e MYSQL_RANDOM_ROOT_PASSWORD=yes -d mysql:latest

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	models.ConnectDatabase(os.Getenv("DATABASE"), os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSERNAME"), os.Getenv("DBPASSWORD"))

	r := routes.SetupRouter()
	fmt.Println(os.Getenv("HOST"), " : ", os.Getenv("PORT"))
	r.Run("localhost:8000")
	// if err := http.ListenAndServe(":8080", r); err != nil {
	// 	panic(err)
	// }

}
