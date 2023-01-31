package main

import (
    "encoding/json"
    "github.com/coderbydesign/jwk2pem"
    "log"
    "net/http"
    "os"
)

func main() {
    http.Handle("/v1/jwt", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenKID := os.Getenv("JWT_KID")
        jb, err := os.ReadFile(os.Getenv("JWK_SET_PATH"))
        if err != nil {
            log.Fatal(err)
        }
        jwksJson := string(jb)
        var b []byte

        keys := jwk2pem.JWKeys{}
        json.Unmarshal([]byte(jwksJson), &keys)

        b = jwk2pem.JWKsToPem(keys, tokenKID)
        w.Write(b)
    }))

    log.Printf("Starting jwk2pem server on :8000")
    http.ListenAndServe(":8000", nil)
}
