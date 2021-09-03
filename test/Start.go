package main

import (
	"context"
	"github.com/PeterYangs/gcmd2"
	"log"
)

func main() {

	cmd := gcmd2.NewCommand("php index.php", context.TODO())

	err := cmd.Start()

	if err != nil {

		log.Println(err)

		return

	}
}
