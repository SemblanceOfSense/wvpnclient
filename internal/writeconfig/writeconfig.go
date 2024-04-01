package writeconfig

import (
	"fmt"
	"os"
	"os/exec"
	"vpnclient/internal/addserverpeer"
	"vpnclient/internal/getwgkey"
)

func GenerateConfig(id int, password string) error {
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

    err = addserverpeer.AddServerPeer(pubKey, id, password)
    if err != nil { return err }

    _ = os.Remove("/etc/wireguard/wg0.conf")

    wgkey, err := getwgkey.GetWgKey(id, password)
    if err != nil { return err }
    newpeer := "[Interface]\nPrivateKey = " + privKey + "\nAddress = 10.0.0.2/24\n\n[Peer]\nPublicKey = " + wgkey + "\nEndpoint = 140.82.19.210:443\nAllowedIPs = 0.0.0.0/0\nPersistentKeepalive = 25"
    err = os.WriteFile("/etc/wireguard/wg0.conf", []byte(newpeer), 1204)
    if err != nil {
        fmt.Println("3")
        return err
    }
    return nil
}
