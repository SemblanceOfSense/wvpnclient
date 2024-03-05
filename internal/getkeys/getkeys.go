package getkeys

import (
	"crypto/rsa"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"vpnclient/internal/addprivkey"
	"vpnclient/internal/adduser"
	"vpnclient/internal/encryptkey"
    "fmt"
)

func GetPublicKey(id int) (rsa.PublicKey, error) {
    req, err := http.NewRequest("GET", "http://149.248.7.39:8080/getpublickey", nil)
    if err != nil {
        fmt.Println("1")
        return rsa.PublicKey{}, err
    }
    req.Header.Set("UserID", strconv.Itoa(id))

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println("2")
        return rsa.PublicKey{}, err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)

    responseStruct := &adduser.PublicKeyRequest{}
    err = json.Unmarshal(body, responseStruct)
    if err != nil {
        fmt.Println(string(body))
        return rsa.PublicKey{}, err
    }

    return responseStruct.Publickey, nil
}

func GetPrivateKey(id int, password string) (rsa.PrivateKey, error) {
    req, err := http.NewRequest("GET", "http://149.248.7.39:8080/getprivatekey", nil)
    if err != nil { return rsa.PrivateKey{}, err }
    req.Header.Set("UserID", strconv.Itoa(id))

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return rsa.PrivateKey{}, err }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)

    responseStruct := &addprivkey.PrivateKeyRequest{}
    err = json.Unmarshal(body, responseStruct)

    decryptedJson := encryptkey.DecryptKey(responseStruct.Ciphertext, []byte(password), responseStruct.Iv, responseStruct.Salt)

    returnStruct := &rsa.PrivateKey{}
    json.Unmarshal(decryptedJson, returnStruct)

    return *returnStruct, nil
}
