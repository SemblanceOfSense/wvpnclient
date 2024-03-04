package main

import (
	"encoding/json"
	"fmt"
	"vpnclient/internal/addprivkey"
	"vpnclient/internal/adduser"
	"vpnclient/internal/encryptkey"
)

func main() {
    id := 1
    key := adduser.AddUser(1)
    fmt.Printf("%d %d\n", id, key.E)

    x, err := json.Marshal(key)
    if err != nil {
        fmt.Println(err)
        return
    }

    password := "password"
    a, b, c := encryptkey.EncryptKey([]byte(string(x)), []byte(password))

    addprivkey.AddPrivKey(id, a, b, c)
}
