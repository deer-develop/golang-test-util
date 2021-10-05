package gomockUtil

import (
	"github.com/golang/mock/gomock"
)

// StringPointerMatcher 함수는 Mock Parameter로 전달된 string 포인터 변수와 전달된 값 변수의 값이 같은지 확인합니다.
// 만약 Mock의 파라미터가 string이 아닌 경우 false를 반환합니다
// 이는 Mock의 파라미터로 포인터가 사용되고 이 값만 검증하고싶은 경우 사용할 수 있습니다
func StringPointerMatcher(str string) gomock.Matcher {
	return NewCustomMatcher(func(arg interface{}) bool {
		strPointer, ok := arg.(*string)
		if !ok {
			return false
		}
		return str == *strPointer
	})
}

// IntPointerMatcher 함수는 Mock Parameter로 전달된 int 포인터 변수와 전달된 값 변수의 값이 같은지 확인합니다.
// 만약 Mock의 파라미터가 int가 아닌 경우 false를 반환합니다
// 이는 Mock의 파라미터로 포인터가 사용되고 이 값만 검증하고싶은 경우 사용할 수 있습니다
func IntPointerMatcher(i int) gomock.Matcher {
	return NewCustomMatcher(func(arg interface{}) bool {
		intPointer, ok := arg.(*int)
		if !ok {
			return false
		}
		return i == *intPointer
	})
}

// BoolPointerMatcher 함수는 Mock Parameter로 전달된 bool 포인터 변수와 전달된 값 변수의 값이 같은지 확인합니다.
// 만약 Mock의 파라미터가 bool이 아닌 경우 false를 반환합니다
// 이는 Mock의 파라미터로 포인터가 사용되고 이 값만 검증하고싶은 경우 사용할 수 있습니다
func BoolPointerMatcher(b bool) gomock.Matcher {
	return NewCustomMatcher(func(arg interface{}) bool {
		boolPointer, ok := arg.(*bool)
		if !ok {
			return false
		}
		return b == *boolPointer
	})
}

// Float32PointerMatcher 함수는 Mock Parameter로 전달된 float32 포인터 변수와 전달된 값 변수의 값이 같은지 확인합니다.
// 만약 Mock의 파라미터가 float32가 아닌 경우 false를 반환합니다
// 이는 Mock의 파라미터로 포인터가 사용되고 이 값만 검증하고싶은 경우 사용할 수 있습니다
func Float32PointerMatcher(f float32) gomock.Matcher {
	return NewCustomMatcher(func(arg interface{}) bool {
		float32Pointer, ok := arg.(*float32)
		if !ok {
			return false
		}
		return f == *float32Pointer
	})
}

// Float64PointerMatcher 함수는 Mock Parameter로 전달된 float64 포인터 변수와 전달된 값 변수의 값이 같은지 확인합니다.
// 만약 Mock의 파라미터가 float64가 아닌 경우 false를 반환합니다
// 이는 Mock의 파라미터로 포인터가 사용되고 이 값만 검증하고싶은 경우 사용할 수 있습니다
func Float64PointerMatcher(f float64) gomock.Matcher {
	return NewCustomMatcher(func(arg interface{}) bool {
		float64Pointer, ok := arg.(*float64)
		if !ok {
			return false
		}
		return f == *float64Pointer
	})
}
