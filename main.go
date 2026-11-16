package main

import (
	"os"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("Hello world")
}

func main() {
	beego.Router("/", &MainController{})

	// Порт передаёт платформа в переменной PORT; без неё — 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	beego.Run("0.0.0.0:" + port)
}
