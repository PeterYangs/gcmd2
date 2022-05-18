package main

import (
	"context"
	"github.com/PeterYangs/gcmd2"
	"log"
	"time"
)

func main() {

	cmd := gcmd2.NewCommand("php index.php", context.TODO())

	err := cmd.Start()

	if err != nil {

		log.Println(err)

		return

	}

	time.Sleep(10 * time.Second)
}
