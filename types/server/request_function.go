package types

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// Request 구조체를 모두 하나의 String으로 변환해주는 함수
func CreateRequestString(csr *CreateServerRequest) string {
	v := url.Values{}
	s := reflect.ValueOf(csr).Elem()

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		jsonTag := strings.Split(s.Type().Field(i).Tag.Get("json"), ",")[0]

		// Skip if the field's value is the zero value for its type
		if field.Interface() == reflect.Zero(field.Type()).Interface() {
			continue
		}

		v.Add(jsonTag, fmt.Sprint(field.Interface()))
	}

	return "?" + v.Encode()
}
