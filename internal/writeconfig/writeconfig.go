package writeconfig

import (
	"fmt"
	"os"
	"os/exec"
)

func GenerateConfig() error {
    cmd := exec.Command("wg", "genkey")

    output, err := cmd.Output()
    if err != nil {
        fmt.Println("1")
        return err
    }
    privKey := string(output)
    privKey = privKey[:len(privKey)-1]


    pubkeyCmd := "echo " + privKey + " | wg pubkey"
    cmd = exec.Command("bash", "-c", pubkeyCmd)

    output, err = cmd.Output()
    if err != nil {
        fmt.Println("2")
        return err
    }
    pubKey := string(output)
    pubKey = pubKey[:len(pubKey)-1]

    _ = os.Remove("/etc/wireguard/wg0.conf")

    newpeer := "[Interface]\nPrivateKey = " + privKey + "\nAddress = 10.0.0.2/24\n\n[Peer]\nPublicKey = " + pubKey + "\nEndpoint = 140.82.19.210:443\nAllowedIPs = 0.0.0.0/0\nPersistentKeepalive = 25"
    err = os.WriteFile("/etc/wireguard/wg0.conf", []byte(newpeer), 1204)
    if err != nil {
        fmt.Println("3")
        return err
    }

    return nil
}
