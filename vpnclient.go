package main

import (
	"fmt"
	"vpnclient/internal/writeconfig"
)

func main() {
    err := writeconfig.GenerateConfig(1, "password")
    if err != nil { fmt.Println(err) }
}
