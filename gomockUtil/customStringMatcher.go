package gomockUtil

import (
	"fmt"
	"github.com/golang/mock/gomock"
)

// NewCustomStringMatcher 함수는 전달된 함수로 Mock 파라미터를 검증하는 Matcher를 반환합니다
// 이때 Mock parameter가 string임을 보장해줍니다
func NewCustomStringMatcher(customMatcher func(arg string) (string, bool)) gomock.Matcher {
	return stringMatcherCustomizer{matcherFunction: customMatcher}
}

type stringMatcherCustomizer struct {
	matcherFunction func(arg string) (string, bool)
	val             interface{}
	reason          string
}

func (o stringMatcherCustomizer) Matches(x interface{}) bool {
	o.val = x
	value, ok := x.(string)
	if !ok {
		o.reason = fmt.Sprintf("value[%v] is not string", x)
		return false
	}
	reason, result := o.matcherFunction(value)
	o.reason = reason
	return result
}

func (o stringMatcherCustomizer) String() string {
	if len(o.reason) != 0 {
		return o.reason
	}
	switch o.val.(type) {
	case string:
		return fmt.Sprintf("value %s is not matched in NewCustomStringMatcher", o.val)
	default:
		return fmt.Sprintf("value[%v] is not string in NewCustomStringMatcher", o.val)
	}
}
