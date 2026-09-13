package main

import (
	"fmt"
	simpleconnection "study/postgres/simple_connection"
)

func main() {
	fmt.Println("Start Main!")
	simpleconnection.CheckConnection()
}
