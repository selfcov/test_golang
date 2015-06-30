package main

import "encoding/base64"
import "encoding/json"
import "fmt"
import "os"
import "strings"

func main() {
    environ_map := make(map[string] string)
    for _, kv := range os.Environ() {
        fmt.Println(kv)
        pair := strings.SplitN(kv, "=", 2)
        environ_map[pair[0]] = pair[1]
    }
    json_thing, _ := json.Marshal(environ_map)
    json_base64 := base64.StdEncoding.EncodeToString(json_thing)
    fmt.Println(json_base64)
}
