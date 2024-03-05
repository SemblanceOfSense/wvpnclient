package adduser

import (
	"bytes"
	crand "crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type PublicKeyRequest struct {
    Publickey rsa.PublicKey
    UserID int
}

func AddUser(id int) (rsa.PrivateKey) {
    privateKey, err := rsa.GenerateKey(crand.Reader, 2048)
    if err != nil {
        panic(err)
    }

    requestBody := &PublicKeyRequest{privateKey.PublicKey, id}

    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        panic(err)
    }

    req, err := http.NewRequest("POST", "http://149.248.7.39:8080/addpublickey", bytes.NewReader(jsonData))
    if err != nil {
        log.Fatalf("impossible to build request: %s", err)
    }

    client := http.Client{Timeout: 10 * time.Second}
    res, err := client.Do(req)
    if err != nil {
        log.Fatalf("impossible to send request: %s", err)
    }
    log.Printf("status Code: %d", res.StatusCode)

    return *privateKey
}
