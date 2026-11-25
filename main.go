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

	// The platform passes the port in the PORT variable; 8080 by default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	beego.Run("0.0.0.0:" + port)
}
