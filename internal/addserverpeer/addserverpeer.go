package addserverpeer

import (
	"bytes"
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"vpnclient/internal/getkeys"
)

type AddServerPeerStruct struct {
    id int
    pubkey string
    signature []byte
}

func AddServerPeer(pubkey string, id int, password string) error {
    msgHash := sha256.New()
    _, err := msgHash.Write([]byte(pubkey))
    if err != nil {
        panic(err)
    }
    msgHashSum := msgHash.Sum(nil)

    privKey, err := getkeys.GetPrivateKey(id, password)
    if err != nil { return err }
    signature, err := rsa.SignPSS(crand.Reader, &privKey, crypto.SHA256, msgHashSum, nil)
    if err != nil {
	    panic(err)
    }

    requestJson, err := json.Marshal(AddServerPeerStruct{id, pubkey, signature})
    if err != nil { return err }

    req, err := http.NewRequest("POST", "http://149.248.7.39:8080/addpeer", bytes.NewReader(requestJson))
    if err != nil { return err }

    client := http.Client{Timeout: 10 * time.Second}
    res, err := client.Do(req)
    if err != nil {
        return err
    }
    log.Printf("status Code: %d", res.StatusCode)

    return nil
}
