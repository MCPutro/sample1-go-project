package main

import (
	"fmt"

	"github.com/MCPutro/sample1-go-project/service"
)

func main() {
	fmt.Println("ahaha, ini dari main")
	sh := service.SayHi()
	fmt.Println(sh)
}
