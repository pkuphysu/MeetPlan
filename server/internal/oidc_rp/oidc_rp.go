package oidc_rp

import (
	"context"
	"sync"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/samber/lo"
	"golang.org/x/oauth2"

	"meetplan/config"
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
	pvd, err := oidc.NewProvider(context.Background(), config.GetConf().Oidc.Host)
	if err != nil {
		panic(err)
	}
	provider = pvd

	oidcConf := config.GetConf().Oidc
	scopes := oidcConf.Scope
	scopes = append(scopes, oidc.ScopeOpenID)
	scopes = lo.Uniq(scopes)
	conf = &oauth2.Config{
		ClientID:     oidcConf.ClientID,
		ClientSecret: oidcConf.ClientSecret,
		Endpoint:     pvd.Endpoint(),
		RedirectURL:  oidcConf.RedirectURL,
		Scopes:       scopes,
	}
}

func GetProvider() *oidc.Provider {
	once.Do(initVariable)
	return provider
}

func GetVerifier() *oidc.IDTokenVerifier {
	return provider.Verifier(&oidc.Config{ClientID: config.GetConf().Oidc.ClientID})
}

func GetOauth2Config() *oauth2.Config {
	once.Do(initVariable)
	return conf
}
