package example

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func createRequestString(csr interface{}) string {
	// Create string with each field name + value in request struct
	v := url.Values{}
	s := reflect.ValueOf(csr).Elem()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		name := strings.ToLower(s.Type().Field(i).Name)

		switch f.Kind() {
		case reflect.String, reflect.Int, reflect.Bool:
			v.Add(name, fmt.Sprint(f.Interface()))
		case reflect.Struct:
			for j := 0; j < f.NumField(); j++ {
				subName := fmt.Sprintf("%s.%d.%s", name, j+1, strings.ToLower(f.Type().Field(j).Name))
				v.Add(subName, fmt.Sprint(f.Field(j).Interface()))
			}
		}
	}

	return "?" + v.Encode()

}
