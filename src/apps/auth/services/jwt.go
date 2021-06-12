package services

import (
	"hciengserver/src/hciengserver"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"

	"github.com/dgrijalva/jwt-go"
)

func MakeJWT(email string) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	signedTkn, err := tkn.SignedString(hciengserver.JWT_SECRET)
	if err != nil {
		return "", err
	}
	return signedTkn, nil
}

func ValidateGoogleJWT(tokenString string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	verifier := googleAuthIDTokenVerifier.Verifier{}
	aud := "835439685490-8j1kg7tk53vhflhp5n9ifmrs164mmbom.apps.googleusercontent.com"
	err := verifier.VerifyIDToken(tokenString, []string{aud})
	if err != nil {
		return nil, err
	}

	claimSet, err := googleAuthIDTokenVerifier.Decode(tokenString)
	if err != nil {
		return nil, err
	}

	return claimSet, nil
}