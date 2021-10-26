package usecase

type UserAgentMatcher interface {
	IsBot(userAgent string) bool
}
