package matcher

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"math"
	"reflect"
)

type JsonMatcher struct {
	json         *gojsonq.JSONQ
	ResultString string
	Result       bool
}

// DefaultTolerance 상수는 float 값을 비교할 때 사용하는 기본 tolerance 값입니다
const DefaultTolerance = 0.0001

// Json 함수는 json 문자열로 Matcher를 생성해서 반환합니다.
func Json(jsonString string) JsonMatcher {
	return JsonMatcher{
		json:   gojsonq.New().FromString(jsonString),
		Result: true,
	}
}

func (m JsonMatcher) findValue(path string) (*gojsonq.Result, error) {
	findValue, err := m.json.FindR(path)
	m.json.Reset()
	return findValue, err
}

// Eq 함수는 앞서 Json(jsonString)으로 생성한 Json string에 대해서
// 전달된 target으로 프로퍼티를 찾아 전달된 value 값과 비교합니다.
// 이때 부동소수점 비교할 때 두 값의 차이가 tolerance 값보다 작은 경우 같다고 판단합니다.
// 이 값을 고르기 힘들다면 DefaultTolerance 상수를 사용해도 됩니다
func (m JsonMatcher) Eq(target string, value interface{}, tolerance float64) JsonMatcher {
	if m.Result == false {
		return m
	}

	findValue, err := m.findValue(target)
	if err != nil {
		m.ResultString = fmt.Sprintf("cannot find %s in json", target)
		m.Result = false
		return m
	}

	if value == nil {
		isNil := findValue.Nil()
		if !isNil {
			m.ResultString = fmt.Sprintf("%v from target[%s] is not nil", findValue, target)
			m.Result = false
		}
		return m
	}

	rValue := reflect.ValueOf(value)
	switch value.(type) {
	case int:
		i := int(rValue.Int())
		i2, err := findValue.Int()
		if err != nil {
			m.ResultString = fmt.Sprintf("cannot convert %s in json to int err[%v]", target, err)
			m.Result = false
			return m
		}
		if i != i2 {
			m.ResultString = fmt.Sprintf("%v from target[%s] is not equal to %v", i2, target, i)
			m.Result = false
		}
		return m
	case string:
		s := rValue.String()
		s2, err := findValue.String()
		if err != nil {
			m.ResultString = fmt.Sprintf("cannot convert %s in json to string err[%v]", target, err)
			m.Result = false
			return m
		}
		if s != s2 {
			m.ResultString = fmt.Sprintf("%v from target[%s] is not equal to %v", s2, target, s)
			m.Result = false
		}
		return m
	case float64:
		f := rValue.Float()
		f2, err := findValue.Float64()
		if err != nil {
			m.ResultString = fmt.Sprintf("cannot convert %s in json to float64 err[%v]", target, err)
			m.Result = false
			return m
		}
		if math.Abs(f-f2) > tolerance {
			m.ResultString = fmt.Sprintf("%v from target[%s] is not equal to %v", f2, target, f)
			m.Result = false
		}
		return m
	case bool:
		b := rValue.Bool()
		b2, err := findValue.Bool()
		if err != nil {
			m.ResultString = fmt.Sprintf("cannot convert %s in json to bool err[%v]", target, err)
			m.Result = false
			return m
		}
		if b != b2 {
			m.ResultString = fmt.Sprintf("%v from target[%s] is not equal to %v", b2, target, b)
			m.Result = false
		}
		return m
	default:
		m.ResultString = fmt.Sprintf("not supported type[%v]", reflect.TypeOf(value))
		m.Result = false
		return m
	}
}
