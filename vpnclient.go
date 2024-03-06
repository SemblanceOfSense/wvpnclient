package main

import (
	"fmt"
	"log"
	"vpnclient/internal/getwgkey"
)

func main() {
    x, err := getwgkey.GetWgKey(1, "password")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(x)
}
