
package main

import (
	"github.com/csalg/carpooling/src/presentation/rest"
	"os"
)


func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 { 
		port = "9091"
	 }

	server := rest.NewServer()
	server.Run(":"+port)

}
