package main

import (
	"fmt"

	utils "github.com/kadekchresna/playground/logging/lumberjack/other"
)

func main() {
	for i := 0; i < 100000; i++ {
		utils.L.Write([]byte(fmt.Sprintf("%d\n",i)))
	}

}