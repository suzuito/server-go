package usecase

type UserAgentMatcher interface {
	IsMatched(userAgent string) bool
}
