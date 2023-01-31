# JWK to PEM

POC for converting a [JSON web key (JWK)](https://www.rfc-editor.org/rfc/rfc7515#section-4.1.3) from a set to PEM, based on a known [`kid`](https://www.rfc-editor.org/rfc/rfc7515#section-4.1.4)

### Setup

Locally you can test this by exporting the following, which will set the JWK set
path and the expected token kid, from within `/test/`:

```bash
# example JWKs from https://auth0.com/docs/secure/tokens/json-web-tokens/json-web-key-set-properties
export JWK_SET_PATH=jwks.json

# the kid from test/jwks.json
export JWT_KID=NjVBRjY5MDlCMUIwNzU4RTA2QzZFMDQ4QzQ2MDAyQjVDNjk1RTM2Qg
```

### Example Output
Run `go run test.go` to start the server, which will expose `/v1/jwt` and use the
`JWK_SET_PATH` and `JWT_KID` to return a sample PEM response on:

```
http://localhost:8000/v1/jwt
```

### Usage
Import and use the package directly:

```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/RedHatInsights/jwk2pem"
    "strings"
)

func main() {
    // the kid you want to match your JWK on
    tokenKID := "some-kid"
    jb, err := os.ReadFile("some-path.json")
    if err != nil {
        fmt.Println(err)
    }
    jwksJson := string(jb)
    var b []byte

    // unmarshal the JWKs
    keys := jwk2pem.JWKeys{}
    json.Unmarshal([]byte(jwksJson), &keys)

    // get the PEM for a JWK set and kid
    b = jwk2pem.JWKsToPem(keys, tokenKID)
    fmt.Println(strings.TrimSuffix(string(b), "\n"))

    // get the PEM for a single JWK
    key := jwk2pem.JWKey{
        Alg: "RS256",
        E:   "AQAB",
        Kid: "NjVBRjY5MDlCMUIwNzU4RTA2QzZFMDQ4QzQ2MDAyQjVDNjk1RTM2Qg",
        Kty: "RSA",
        N:   "yeNlzlub94YgerT030codqEztjfU_S6X4DbDA_iVKkjAWtYfPHDzz_sPCT1Axz6isZdf3lHpq_gYX4Sz-cbe4rjmigxUxr-FgKHQy3HeCdK6hNq9ASQvMK9LBOpXDNn7mei6RZWom4wo3CMvvsY1w8tjtfLb-yQwJPltHxShZq5-ihC9irpLI9xEBTgG12q5lGIFPhTl_7inA1PFK97LuSLnTJzW0bj096v_TMDg7pOWm_zHtF53qbVsI0e3v5nmdKXdFf9BjIARRfVrbxVxiZHjU6zL6jY5QJdh1QCmENoejj_ytspMmGW7yMRxzUqgxcAqOBpVm0b-_mW3HoBdjQ",
        Use: "sig",
    }
    b = jwk2pem.JWKToPem(key)
    fmt.Println(strings.TrimSuffix(string(b), "\n"))
}
```

### TODO:
- Account for tokens with variable `kid` values which we'll need to dynamically check against the JWKs
