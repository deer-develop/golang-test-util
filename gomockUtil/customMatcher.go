package gomockUtil

import (
	"github.com/golang/mock/gomock"
)

// NewCustomMatcher 함수는 전달된 함수로 Mock 파라미터를 검증하는 Matcher를 반환합니다
func NewCustomMatcher(customMatcher func(arg interface{}) bool) gomock.Matcher {
	return matcherCustomizer{customMatcher}
}

type matcherCustomizer struct {
	matcherFunction func(arg interface{}) bool
}

func (o matcherCustomizer) Matches(x interface{}) bool {
	return o.matcherFunction(x)
}

func (o matcherCustomizer) String() string {
	return "[call back function matcher has returned false]"
}
