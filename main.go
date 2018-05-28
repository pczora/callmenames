package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/docker/docker/pkg/namesgenerator"
)

func main() {
	argsWithoutApp := os.Args[1:]
	if len(argsWithoutApp) > 0 {
		i := 0
		nameCount, err := strconv.Atoi(argsWithoutApp[0])

		if err != nil {
			fmt.Println("Invalid argument")
			return
		}
		for i < nameCount {
			fmt.Println(namesgenerator.GetRandomName(0))
			i = i + 1
		}
	} else {
		fmt.Println(namesgenerator.GetRandomName(0))
	}

}
