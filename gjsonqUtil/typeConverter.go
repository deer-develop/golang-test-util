// gjsonqUtil 패키지는 gjsonq 라이브러리를 조금 더 편하게 사용할 수 있게 해주는 wrapper입니다.

package gjsonqUtil

import "github.com/thedevsaddam/gojsonq/v2"

// Bool 함수는 JSONQ와 findPath를 받아서 bool로 변환한 값을 에러를 반환합니다.
// 내부에서 jsonq.Reset() 메소드를 호출하므로 전달된 JSONQ를 재사용할 수 있습니다.
func Bool(jsonq *gojsonq.JSONQ, findR string) (bool, error) {
	result, err := jsonq.FindR(findR)
	jsonq.Reset()
	if err != nil {
		return false, err
	}
	b, err := result.Bool()
	if err != nil {
		return false, err
	}
	return b, nil
}

// Int 함수는 JSONQ와 findPath를 받아서 int 변환한 값을 에러를 반환합니다.
// 내부에서 jsonq.Reset() 메소드를 호출하므로 전달된 JSONQ를 재사용할 수 있습니다.
func Int(jsonq *gojsonq.JSONQ, findR string) (int, error) {
	result, err := jsonq.FindR(findR)
	jsonq.Reset()
	if err != nil {
		return 0, err
	}
	i, err := result.Int()
	if err != nil {
		return 0, err
	}
	return i, nil
}

// String 함수는 JSONQ와 findPath를 받아서 string 변환한 값을 에러를 반환합니다.
// 내부에서 jsonq.Reset() 메소드를 호출하므로 전달된 JSONQ를 재사용할 수 있습니다.
func String(jsonq *gojsonq.JSONQ, findR string) (string, error) {
	result, err := jsonq.FindR(findR)
	jsonq.Reset()
	if err != nil {
		return "", err
	}
	s, err := result.String()
	if err != nil {
		return "", err
	}
	return s, nil
}

// Float64 함수는 JSONQ와 findPath를 받아서 float64 변환한 값을 에러를 반환합니다.
// 내부에서 jsonq.Reset() 메소드를 호출하므로 전달된 JSONQ를 재사용할 수 있습니다.
func Float64(jsonq *gojsonq.JSONQ, findR string) (float64, error) {
	result, err := jsonq.FindR(findR)
	jsonq.Reset()
	if err != nil {
		return 0, err
	}
	f, err := result.Float64()
	if err != nil {
		return 0, err
	}
	return f, nil
}
