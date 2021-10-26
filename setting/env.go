package setting

import (
	"net/url"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

// Environment ...
type Environment struct {
	Env                string `envconfig:"ENV"`
	PrerenderURLString string `envconfig:"PRERENDER_URL"`
	PrerenderURL       *url.URL
	FrontURLString     string `envconfig:"FRONT_URL"`
	FrontURL           *url.URL
}

func NewEnvironment() (*Environment, error) {
	r := Environment{}
	if err := envconfig.Process("", &r); err != nil {
		return nil, xerrors.Errorf("Cannot envconfig.Process : %w", err)
	}
	var err error
	r.PrerenderURL, err = url.Parse(r.PrerenderURLString)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	r.FrontURL, err = url.Parse(r.FrontURLString)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return &r, nil
}
