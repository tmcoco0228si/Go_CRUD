package main

import (
	"Go_CRUD/app/controllers"
	"Go_CRUD/app/models"
	"fmt"
)

func main() {
	fmt.Println(models.Db)
	controllers.StartMainServer()
}
