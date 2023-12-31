package types

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strings"
)

// Request 구조체를 모두 하나의 String으로 변환해주는 함수
func CreateRequestString(csr *CreateServerRequest, networkInterfaceIndex int, acgNoList int) string {
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
			jsonTag = strings.Replace(jsonTag, "N", fmt.Sprint(networkInterfaceIndex), 1)
			jsonTag = strings.Replace(jsonTag, ".N", "."+fmt.Sprint(acgNoList), 1)
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

	responseBody = processTimestamp(responseBody)
	err := xml.Unmarshal(responseBody, v) // responseBody를 v로 매핑. 만약 CreateServerResponse 타입이면 CreateServerResponse로 매핑
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return v, nil
}

func RequestString(req interface{}) string {
	v := url.Values{}
	s := reflect.ValueOf(req).Elem()

	for i := 0; i < s.NumField(); i++ {
		field := s.Field(i)
		jsonTag := strings.Split(s.Type().Field(i).Tag.Get("json"), ",")[0]
		v.Add(jsonTag, fmt.Sprint(field.Interface()))
	}

	return "?" + v.Encode()
}

func processTimestamp(input []byte) (resultReponse []byte) {
	regex := regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2})([+-]\d{4})`)

	// Find all occurrences of timestamps in the input
	matches := regex.FindAllSubmatchIndex(input, -1)

	if len(matches) > 0 {
		// Create a copy of the original input
		modifiedInput := make([]byte, len(input)+len(matches))
		copy(modifiedInput, input)

		for i, match := range matches {
			start, end := match[2]-4*i, match[3]-4*i
			appendingTS := []byte{90} // represents 'Z' in ascii
			originalTS := make([]byte, 19)
			copy(originalTS, modifiedInput[start:end][:19])
			timestamp := append(originalTS, appendingTS...)
			modifiedInput = append(modifiedInput[:start], append([]byte(timestamp), modifiedInput[end+5:]...)...)
		}
		return modifiedInput
	} else {
		fmt.Println("Timestamps not found in the input")
	}
	return input
}
