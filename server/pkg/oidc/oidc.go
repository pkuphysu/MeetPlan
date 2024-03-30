package oidc

import (
	"context"
	"sync"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type Claims struct {
	Sub         string `json:"sub"`
	Name        string `json:"name"`
	GivenName   string `json:"given_name"`
	FamilyName  string `json:"family_name"`
	Nickname    string `json:"nickname"`
	Website     string `json:"website"`
	Gender      string `json:"gender"`
	Birthdate   string `json:"birthdate"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     struct {
		Formatted string `json:"formatted"`
	} `json:"address"`
	IsPku      bool     `json:"is_pku"`
	PkuID      string   `json:"pku_id"`
	PkuEmail   string   `json:"pku_email"`
	IsTeacher  bool     `json:"is_teacher"`
	InSchool   bool     `json:"in_school"`
	Department string   `json:"department"`
	Introduce  string   `json:"introduce"`
	Groups     []string `json:"groups"`
}

var (
	provider *oidc.Provider
	conf     *oauth2.Config
	once     sync.Once
)

func initVariable() {
	pvd, err := oidc.NewProvider(context.Background(), "https://auth.phy.pku.edu.cn/oidc")
	if err != nil {
		panic(err)
	}
	provider = pvd

	conf = &oauth2.Config{
		ClientID:     "16302204390022",
		ClientSecret: "c2c51981a1e716d6b5dc32b986a3f923d57c014271cf37d8bd3e8660",
		Endpoint:     pvd.Endpoint(),
		RedirectURL:  "http://localhost:3000/#/login",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "phone", "address", "pku"},
	}
}

func GetProvider() *oidc.Provider {
	once.Do(initVariable)
	return provider
}

func GetVerifier() *oidc.IDTokenVerifier {
	return provider.Verifier(&oidc.Config{ClientID: "16302204390022"})
}

func GetOauth2Config() *oauth2.Config {
	once.Do(initVariable)
	return conf
}
