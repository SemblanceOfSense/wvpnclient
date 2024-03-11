package main

import (
	"fmt"
	"github.com/webui-dev/go-webui/v2"
)

const script = "return [document.getElementById('userid'), document.getElementById('password')]"

func signup(e webui.Event) string {
    var Timeout uint = 10
    var opts = webui.ScriptOptions{Timeout: Timeout}
    resp, err := e.Window.Script(script, opts)
    if err != nil {
        e.Window.Show("<p>" + err.Error() + "</p>")
    }
    fmt.Println(resp)
    return ""
}

func main() {
    index := webui.NewWindow()
    index.Show("webui/index.html")

    webui.Bind(index, "signup", signup)
}
