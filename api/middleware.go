package api

import (
	"errors"
	"fmt"
	"os"

	"github.com/anthdm/weavebox"
	"github.com/golang-jwt/jwt/v4"
)

var ErrUnAuthenticated = errors.New("unAuthenticated")

// // sample token string taken from the New example
// tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

// // Parse takes the token string and a function for looking up the key. The latter is especially
// // useful if you use multiple keys for your application.  The standard is to use 'kid' in the
// // head of the token to identify which key to use, but the parsed token (head and claims) is provided
// // to the callback, providing flexibility.
// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 	// Don't forget to validate the alg is what you expect:
// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 	}

// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
// 	return hmacSampleSecret, nil
// })

type AdminAuthMiddleware struct{}

func (mw *AdminAuthMiddleware) Authenticate(ctx *weavebox.Context) error {
	tokenString := ctx.Header("x-api-token")
	if len(tokenString) == 0 {
		return ErrUnAuthenticated
	}
	token, err := parseJWT(tokenString)
	if err != nil {
		return ErrUnAuthenticated
	}
	if !token.Valid {
		return ErrUnAuthenticated
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrUnAuthenticated
	}

	// id := claims["userID"]

	fmt.Println(claims)
	fmt.Println("guarding the admin routes")
	return nil
}

func parseJWT(tokenString string) (*jwt.Token, error) {
	secret := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
}
