package simplcmd

import (
	"fmt"
	"os"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/kc"
	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/simpl"
)

type simplACFT struct {
	User      string
	Pass      string
	Realm     string
	BaseUrl   string
	KubeProxy bool
}

type Client struct {
	*simpl.Client
	*KeyCloakTokenizator
}

func (s *simplACFT) NewClient(defUsr string) *Client {
	e := s.newSimplEndpoint()
	t := sacf.newKeyCloakTokenizator(e, defUsr)
	return &Client{
		Client: &simpl.Client{
			Endpoints: e,
			TokenFunc: func() (string, error) {
				if s.KubeProxy {
					return GetKubernetesToken()
				} else {
					return t.ApiToken()
				}
			},
		},
		KeyCloakTokenizator: t,
	}
}

func (s *simplACFT) newSimplEndpoint() *simpl.SingleAccessEndpoints {
	if s.KubeProxy {
		return &simpl.SingleAccessEndpoints{
			BaseUrl:                   s.BaseUrl,
			KeycloakUrl:               "/services/keycloak/proxy",
			AuthenticationProviderUrl: "/services/authentication-provider:8080/proxy",
		}
	} else {
		return &simpl.SingleAccessEndpoints{
			BaseUrl: s.BaseUrl,
		}
	}
}

func (s *simplACFT) newKeyCloakTokenizator(e simpl.Endpoints, defUsr string) *KeyCloakTokenizator {
	kcclient := &kc.Client{
		BaseUrl: e.Keycloak(),
		AuthFunc: func() (*kc.AuthInfo, error) {
			user := s.User
			if user == "" {
				user = defUsr
			}
			return &kc.AuthInfo{
				Username:  user,
				Passaword: s.Pass,
				Realm:     s.Realm,
				ClientId:  "frontend-cli",
			}, nil
		},
	}
	return &KeyCloakTokenizator{
		Endpoints: e,
		keycl:     kcclient,
	}
}

type KeyCloakTokenizator struct {
	Endpoints simpl.Endpoints
	keycl     *kc.Client
}

func (k *KeyCloakTokenizator) ApiToken() (string, error) {
	kt, err := k.keycl.Tokenize()
	if err != nil {
		return "", err
	}

	return kt.AccessToken, nil
}

func GetKubernetesToken() (string, error) {
	if t, ok := os.LookupEnv("KTOKEN"); ok {
		return t, nil
	} else {
		return "", fmt.Errorf("env KTOKEN not defined")
	}
}
