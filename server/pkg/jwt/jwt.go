package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"meetplan/model"
)

//var privateKey *rsa.PrivateKey
//
//func init() {
//	pvk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte())
//	if err != nil {
//		panic(err)
//	}
//	privateKey = pvk
//}

var (
	CtxUserKey = struct{}{}
)

type Claims struct {
	jwt.RegisteredClaims
	PkuID     string `json:"pku_id"`
	IsTeacher bool   `json:"is_teacher"`
	IsAdmin   bool   `json:"is_admin"`
}

func (c *Claims) GetUserID() string {
	return c.Subject
}

func NewJwt(user *model.User, valid time.Duration) (string, error) {
	now := time.Now()
	claim := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "pku-phy-meetplan",
			Subject:   user.ID.Hex(),
			ExpiresAt: jwt.NewNumericDate(now.Add(valid)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		PkuID:     user.PkuID,
		IsTeacher: user.IsTeacher,
		IsAdmin:   user.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodNone, claim)
	return token.SignedString(jwt.UnsafeAllowNoneSignatureType)
}

func VerifyJwt(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return jwt.UnsafeAllowNoneSignatureType, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*Claims), nil
}
