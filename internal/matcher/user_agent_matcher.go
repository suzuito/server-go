package matcher

import (
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/xerrors"
)

type UserAgentMatcher struct {
	re *regexp.Regexp
}

func (m *UserAgentMatcher) IsMatched(userAgent string) bool {
	return m.re.MatchString(userAgent)
}

func NewUserAgentMatcher(
	expresses []string,
) (*UserAgentMatcher, error) {
	if len(expresses) <= 0 {
		return nil, fmt.Errorf("Empty expresses")
	}
	r, err := regexp.Compile(strings.Join(expresses, "|"))
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return &UserAgentMatcher{
		re: r,
	}, nil
}

func NewExternalBotMatcherDefault() (*UserAgentMatcher, error) {
	m, err := NewUserAgentMatcher([]string{
		"^.*googlebot.*$",
		"^.*twitterbot.*$",
		"^.*facebookexternalhit.*$",
	})
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return m, nil
}

func NewHealthCheckBotMatcherDefault() (*UserAgentMatcher, error) {
	m, err := NewUserAgentMatcher([]string{
		"^kube-probe.*$",
		"^GoogleHC.*$",
	})
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	return m, nil
}
