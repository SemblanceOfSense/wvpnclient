package main

import (
	"fmt"
	"vpnclient/internal/writeconfig"
)

func main() {
    err := writeconfig.GenerateConfig()
    if err != nil { fmt.Println(err) }
}
