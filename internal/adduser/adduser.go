package adduser

import (
	"bytes"
	crand "crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"log"
	mrand "math/rand/v2"
	"net/http"
	"time"
)

type PublicKeyRequest struct {
    Publickey rsa.PublicKey
    UserID int
}

func AddUser() (int, rsa.PrivateKey) {
    privateKey, err := rsa.GenerateKey(crand.Reader, 2048)
    if err != nil {
        panic(err)
    }

    requestBody := &PublicKeyRequest{privateKey.PublicKey, mrand.IntN(1000000000)}

    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        panic(err)
    }

    req, err := http.NewRequest("POST", "http://localhost:8080/addpublickey", bytes.NewReader(jsonData))
    if err != nil {
        log.Fatalf("impossible to build request: %s", err)
    }

    client := http.Client{Timeout: 10 * time.Second}
    res, err := client.Do(req)
    if err != nil {
        log.Fatalf("impossible to send request: %s", err)
    }
    log.Printf("status Code: %d", res.StatusCode)

    return requestBody.UserID, *privateKey
}
