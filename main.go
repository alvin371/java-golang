package main

import (
	"java-agro/driver"
	routes "java-agro/route"
)

func main() {
	driver.InitDB()
	e := routes.New()

	e.Start(":8000")
}
