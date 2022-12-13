package jwk2pem

import (
    "crypto/x509"
    "encoding/json"
    "encoding/pem"
    "fmt"
    "github.com/lestrrat-go/jwx/jwk"
)

func JWKsToPem(keys JWKeys, kid string) []byte {
    var b []byte
    for _, key := range keys.Keys {
        if key.Kid == kid {
            b = JWKToPem(key)
        }
    }
    return b
}

func JWKToPem(key JWKey) []byte {
    var b []byte
    jKey, err := json.Marshal(key)
    if err != nil {
        fmt.Println(err)
    }

    k, err := jwk.ParseKey([]byte(jKey))
    if err != nil {
        fmt.Println(err)
    }

    var rawKey interface{}
    if err := k.Raw(&rawKey); err != nil {
        fmt.Println(err)
    }

    pubData, err := x509.MarshalPKIXPublicKey(rawKey)
    if err != nil {
        fmt.Println(err)
    }

    b = pem.EncodeToMemory(&pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: pubData,
    })
    if err != nil {
        fmt.Println(err)
    }

    return b
}
