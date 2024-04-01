package main

import (
	"log"
    "strconv"
    "os"
    "encoding/json"

    "github.com/tcnksm/go-input"

    "vpnclient/internal/adduser"
    "vpnclient/internal/encryptkey"
    "vpnclient/internal/addprivkey"
    "vpnclient/internal/writeconfig"
)

func main() {
    ui := &input.UI {
        Writer: os.Stdout,
        Reader: os.Stdin,
    }
    options := &input.Options{
        Required: true,
        Loop:     true,
    }

    query := "signup or login?"
    initial, err := ui.Ask(query, options)

    if err != nil { log.Fatal(err) }

    query = "UserID?"
    readid, err := ui.Ask(query, options)
    if err != nil { log.Fatal(err) }
    id, err := strconv.Atoi(readid)
    if err != nil { log.Fatal(err) }
    query = "Password?"
    password, err := ui.Ask(query, options)
    if err != nil { log.Fatal(err) }

    switch initial {
    case "signup":
        key := adduser.AddUser(id)
        x, err := json.Marshal(key)
        if err != nil { log.Fatal(err) }
        a, b, c := encryptkey.EncryptKey([]byte(string(x)), []byte(password))

        addprivkey.AddPrivKey(id, a, b, c)
    case "login":
        err = writeconfig.GenerateConfig(id, password)
        if err != nil { log.Fatal(err) }
    }
}
