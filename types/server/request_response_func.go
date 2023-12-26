package types

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// Request 구조체를 모두 하나의 String으로 변환해주는 함수
func CreateRequestString(csr *CreateServerRequest, networkInterfaceIndex int) string {
	v := url.Values{}
	s := reflect.ValueOf(csr).Elem()

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		fieldName := s.Type().Field(i).Name // Get the field name
		jsonTag := strings.Split(s.Type().Field(i).Tag.Get("json"), ",")[0]

		// If the field is NetworkInterfaceOrder, change jsonTag's 'N' value to networkInterfaceIndex
		// Before: networkInterfaceList.N.networkInterfaceOrder
		// After: networkInterfaceList.0.networkInterfaceOrder
		if fieldName == "NetworkInterfaceOrder" {
			jsonTag = strings.Replace(jsonTag, "N", fmt.Sprint(networkInterfaceIndex), 1)
		}

		if fieldName == "AccessControlGroupNoListN" {

		}

		v.Add(jsonTag, fmt.Sprint(field.Interface()))
	}

	return "?" + v.Encode()
}

// Response XML을 struct로 변환하는 함수
// responseBody: API 서버로부터 받은 XML response body
// v: interface{} 타입의 struct 포인터
// 예시: mapResponse(responseBody, &CreateServerResponse{})
func MapResponse(responseBody []byte, v interface{}) (interface{}, error) {
	rv := reflect.ValueOf(v)                    // reflect를 이용해 v가 어떤 구조체 타입인지 알아냄
	if rv.Kind() != reflect.Ptr || rv.IsNil() { // 만약 v가 포인터가 아니거나 nil이면 에러 반환
		return nil, fmt.Errorf("non-nil pointer expected")
	}

	err := xml.Unmarshal(responseBody, v) // responseBody를 v로 매핑. 만약 CreateServerResponse 타입이면 CreateServerResponse로 매핑
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return v, nil
}
