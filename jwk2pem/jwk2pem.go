package jwk2pem

import (
    "crypto/x509"
    "encoding/json"
    "encoding/pem"
    "github.com/lestrrat-go/jwx/jwk"
    "log"
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
        log.Fatal(err)
    }

    k, err := jwk.ParseKey([]byte(jKey))
    if err != nil {
        log.Fatal(err)
    }

    var rawKey interface{}
    if err := k.Raw(&rawKey); err != nil {
        log.Fatal(err)
    }

    pubData, err := x509.MarshalPKIXPublicKey(rawKey)
    if err != nil {
        log.Fatal(err)
    }

    b = pem.EncodeToMemory(&pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: pubData,
    })
    if err != nil {
        log.Fatal(err)
    }

    return b
}
