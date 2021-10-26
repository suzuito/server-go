package matcher

type UserAgentMatcher struct{}

func (m *UserAgentMatcher) IsBot(userAgent string) bool {
	return false
}

func NewUserAgentMatcher() *UserAgentMatcher {
	return &UserAgentMatcher{}
}
