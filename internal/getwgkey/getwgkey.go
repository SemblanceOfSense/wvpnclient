package getwgkey

import (
	"crypto/rsa"
	"crypto/sha256"
	"hash"
	"io"
	"net/http"
	"strconv"
	"vpnclient/internal/getkeys"
    crand "crypto/rand"
)

type EncryptedWgKey struct {
    sha hash.Hash
    rand io.Reader
    cipherText []byte
}

func GetWgKey(id int, password string) (string, error) {
    privKey, err := getkeys.GetPrivateKey(id, password)
    if err != nil { return "", err }

    req, err := http.NewRequest("GET", "http://149.248.7.39:8080/getwgkey", nil)
    req.Header.Set("UserID", strconv.Itoa(id))
    if err != nil { return "", err }

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return "", err }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)

    privateKey := &privKey
    decryptedBytes, err := rsa.DecryptOAEP(sha256.New(), crand.Reader, privateKey, body, nil)
    if err != nil { return "", err }

    returnValue := string(decryptedBytes)
    return returnValue[1 : len(returnValue)-1], nil
}
