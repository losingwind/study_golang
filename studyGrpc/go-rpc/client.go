package main

import (
	"fmt"
	"log"
	"study_golang/studyGrpc/go-rpc/pkg"
)

func main() {
	var rsp pkg.String

	c, err := pkg.DialServer()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := c.Request(&rsp); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(rsp.Value)
}
