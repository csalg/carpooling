// TODO
// * Could add a yaml with parameters. Some parameters: car sizes, group sizes.

package main

import (
	"os"
)


func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 { 
		port = "9091"
	 }

	server := NewServer()
	server.Run(":"+port)
	//http.ListenAndServe("0.0.0.0:"+port, server)

}
