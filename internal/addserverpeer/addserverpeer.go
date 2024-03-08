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
    Id int
    Pubkey string
    Signature []byte
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

    requestBody := &AddServerPeerStruct{Id: id, Pubkey: pubkey, Signature: signature}

    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        panic(err)
    }

    req, err := http.NewRequest("POST", "http://149.248.7.39:8080/addpeer", bytes.NewReader(jsonData))
    if err != nil {
        log.Fatalf("impossible to build request: %s", err)
    }

    client := http.Client{Timeout: 10 * time.Second}
    res, err := client.Do(req)
    if err != nil {
        log.Fatalf("impossible to send request: %s", err)
    }
    defer res.Body.Close()
    log.Printf("status Code: %d", res.StatusCode)


    return nil
}
