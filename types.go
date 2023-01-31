package jwk2pem

type JWKeys struct {
    Keys []JWKey `json:keys`
}

type JWKey struct {
    Alg string `json:"alg"`
    E   string `json:"e"`
    Kid string `json:"kid"`
    Kty string `json:"kty"`
    N   string `json:"n"`
    Use string `json:"use"`
}
