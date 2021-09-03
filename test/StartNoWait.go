package main

import (
	"context"
	"gcmd2"
	"log"
)

func main() {

	cmd := gcmd2.NewCommand("php index.php", context.TODO())

	err := cmd.StartNoWait()

	if err != nil {

		log.Println(err)

		return

	}
}
