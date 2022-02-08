package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUserAgentMatcherDefault(t *testing.T) {
	matcher, err := NewExternalBotMatcherDefault()
	assert.Nil(t, err)
	assert.True(t, matcher.IsMatched("Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html)"))
	assert.True(t, matcher.IsMatched("Twitterbot/1.0"))
	matcher, err = NewHealthCheckBotMatcherDefault()
	assert.Nil(t, err)
	assert.True(t, matcher.IsMatched("kube-probe"))
	assert.True(t, matcher.IsMatched("GoogleHC"))
}

func TestNewUserAgentMatcher(t *testing.T) {
	testCases := []struct {
		desc           string
		inputExpresses []string
		expectedErr    string
	}{
		{
			desc:           "失敗 入力された正規表現が空",
			inputExpresses: []string{},
			expectedErr:    "Empty expresses",
		},
		{
			desc: "成功",
			inputExpresses: []string{
				"a",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			_, realErr := NewUserAgentMatcher(tC.inputExpresses)
			if realErr != nil {
				assert.NotEmpty(t, tC.expectedErr)
				assert.Regexp(t, tC.expectedErr, realErr.Error())
				return
			}
		})
	}
}

func TestIsBot(t *testing.T) {
	testCases := []struct {
		desc           string
		inputExpresses []string
		inputUserAgent string
		expected       bool
	}{
		{
			desc: "UserAgentがボット判定される",
			inputExpresses: []string{
				"^.*googlebot.*$",
			},
			inputUserAgent: "hogegooglebotfuga",
			expected:       true,
		},
		{
			desc: "UserAgentがボット判定されない",
			inputExpresses: []string{
				"^.*googlebot.*$",
			},
			inputUserAgent: "hogetwitterbotfuga",
			expected:       false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			m, _ := NewUserAgentMatcher(tC.inputExpresses)
			assert.Equal(t, tC.expected, m.IsMatched(tC.inputUserAgent))
		})
	}
}
