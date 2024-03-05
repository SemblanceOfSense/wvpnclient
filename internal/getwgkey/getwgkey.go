package getwgkey

import (
	"crypto"
	"crypto/rsa"
	"io"
	"net/http"
	"vpnclient/internal/getkeys"
)

func GetWgKey(id int, password string) (string, error) {
    privKey, err := getkeys.GetPrivateKey(id, password)

    req, err := http.NewRequest("GET", "http://140.82.19.210:8080/privkey", nil)
    if err != nil { return "", err }

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return "", err }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)

    decryptedBytes, err := privKey.Decrypt(nil, body, &rsa.OAEPOptions{Hash: crypto.SHA256})
    if err != nil { return "", err }

    return string(decryptedBytes), nil
}
