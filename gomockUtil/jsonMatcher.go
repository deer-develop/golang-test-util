package gomockUtil

import (
	"fmt"
	"iot-translator/util/test/matcher"
)

func NewJsonMatcher() *GomockJsonMatcher {
	return new(GomockJsonMatcher)
}

type GomockJsonMatcher struct {
	target []string
	value  []interface{}
	reason string
}

func (m *GomockJsonMatcher) Matches(x interface{}) bool {
	if len(m.target) != len(m.value) {
		m.reason = fmt.Sprintf("Internal error: target[%v] length is not equal value[%v] length", m.target, m.value)
		return false
	}

	jsonString, ok := x.(string)
	if !ok {
		m.reason = fmt.Sprintf("Value is not String[%v]", x)
		return false
	}

	jsonMatcherImpl := matcher.Json(jsonString)
	for i, target := range m.target {
		value := m.value[i]
		jsonMatcherImpl = jsonMatcherImpl.Eq(target, value, matcher.DefaultTolerance)
		if !jsonMatcherImpl.Result {
			m.reason = jsonMatcherImpl.ResultString
			return false
		}
	}
	m.reason = jsonMatcherImpl.ResultString
	return jsonMatcherImpl.Result
}

func (m *GomockJsonMatcher) String() string {
	return m.reason
}

func (m *GomockJsonMatcher) Eq(target string, value interface{}) *GomockJsonMatcher {
	m.target = append(m.target, target)
	m.value = append(m.value, value)
	return m
}
