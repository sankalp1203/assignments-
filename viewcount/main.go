package main

import (
	"viewcount/routes"
)

const portNumber = ":8080"

func main() {

	r := routes.Routes()

	r.Run(portNumber)

}
