package main

import (
	"context"
	"fmt"
	"gcmd2"
	"log"
)

func main() {

	cmd := gcmd2.NewCommand("dir", context.TODO())

	out, err := cmd.CombinedOutput()

	if err != nil {

		log.Println(err)

		return

	}

	fmt.Println(string(out))

}
