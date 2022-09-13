package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	jwtgo "github.com/golang-jwt/jwt"
)

// GenerateToken creates a jwt.Token out of the provided claims.
func GenerateToken(kid string, claimsJSON jwtgo.MapClaims) *jwtgo.Token {
	t := jwtgo.NewWithClaims(jwtgo.SigningMethodRS256, claimsJSON)
	t.Header["kid"] = kid
	return t
}

// Sign signs the token using the provided *rsa.PrivateKey
func Sign(token *jwtgo.Token, key *rsa.PrivateKey) string {
	s, _ := token.SignedString(key)
	return s
}

// LoadRSAKey parses the provided PEM string into an *rsa.PrivateKey
func LoadRSAKey(pemString string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(pemString))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}
