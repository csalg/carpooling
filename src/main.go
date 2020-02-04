// TODO
// * Could add a yaml with parameters. Some parameters: car sizes, group sizes.
// * Do a writeup on this thing:
//   * SOLID principles.
//   * Time complexity, scalability and latency.
//   * Running match asynchronously.
// * Running match asynchronously in another thread 

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

}