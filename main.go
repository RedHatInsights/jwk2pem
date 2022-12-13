package main

import (
    "encoding/json"
    "fmt"
    "jwk2pem/jwk2pem"
    "os"
    "strings"
)

func main() {
    tokenKID := os.Getenv("JWT_KID")
    jb, err := os.ReadFile(os.Getenv("JWK_SET_PATH"))
    if err != nil {
        fmt.Println(err)
    }
    jwksJson := string(jb)
    var b []byte

    keys := jwk2pem.JWKeys{}
    json.Unmarshal([]byte(jwksJson), &keys)

    b = jwk2pem.JWKsToPem(keys, tokenKID)
    fmt.Println(strings.TrimSuffix(string(b), "\n"))
}
