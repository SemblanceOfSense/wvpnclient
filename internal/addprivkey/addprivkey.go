package addprivkey

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type PrivateKeyRequest struct {
    Ciphertext []byte
    Iv []byte
    Salt []byte
    UserID int
}

func AddPrivKey(id int, ciphertext []byte, iv []byte, salt []byte) {
    requestBody := &PrivateKeyRequest{ciphertext, iv, salt, id}

    jsonData, err := json.Marshal(requestBody)
    if err != nil {
        panic(err)
    }

    req, err := http.NewRequest("POST", "http://149.248.7.39:8080/addprivatekey", bytes.NewReader(jsonData))
    if err != nil {
        log.Fatalf("impossible to build request: %s", err)
    }

    client := http.Client{Timeout: 10 * time.Second}
    res, err := client.Do(req)
    if err != nil {
        log.Fatalf("impossible to send request: %s", err)
    }
    log.Printf("status Code: %d", res.StatusCode)
}
