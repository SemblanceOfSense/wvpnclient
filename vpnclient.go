package main

import (
	"encoding/json"
	"os"
	"vpnclient/internal/getkeys"
    "fmt"
)

func main() {
    id := 1

    pubkey, err := getkeys.GetPublicKey(id)
    if err != nil {
        fmt.Println("1 " + err.Error())
        return
    }
    jsonPubKey, err := json.Marshal(pubkey)
    if err != nil {
        fmt.Println("2 " + err.Error())
        return
    }
    os.Remove("/etc/wvpn/pubkey")
    os.WriteFile("/etc/wvpn/pubkey", jsonPubKey, 1411)

    privkey, err := getkeys.GetPrivateKey(id, "password")
    if err != nil {
        fmt.Println("3 " + err.Error())
        return
    }
    jsonPrivKey, err := json.Marshal(privkey)
    if err != nil {
        fmt.Println("4 " + err.Error())
        return
    }
    os.Remove("/etc/wvpn/privkey")
    os.WriteFile("/etc/wvpn/privkey", jsonPrivKey, 1411)
}
