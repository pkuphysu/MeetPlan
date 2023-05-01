package jwt

import (
	"crypto/rsa"
	"strconv"
	"time"

	"meetplan/config"

	"github.com/golang-jwt/jwt/v5"

	"meetplan/biz/gorm_gen"
)

var privateKey *rsa.PrivateKey

func init() {
	pvk, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.GetConf().Hertz.PrivateKey))
	if err != nil {
		panic(err)
	}
	privateKey = pvk
}

type Claims struct {
	jwt.RegisteredClaims
	PkuID     string `json:"pku_id"`
	IsTeacher bool   `json:"is_teacher"`
	IsAdmin   bool   `json:"is_admin"`
}

func NewJwt(user *gorm_gen.User) (string, error) {
	now := time.Now()
	claim := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "PKU PHY MeetPlan WebSite",
			Subject:   strconv.FormatInt(user.ID, 10),
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * 7)),
			NotBefore: jwt.NewNumericDate(now),
			IssuedAt:  jwt.NewNumericDate(now),
			//ID:        "",
		},
		PkuID:     user.PkuID,
		IsTeacher: user.IsTeacher,
		IsAdmin:   user.IsAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	return token.SignedString(privateKey)
}

func VerifyJwt(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(_ *jwt.Token) (interface{}, error) {
		return privateKey.Public(), nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*Claims), nil
}
