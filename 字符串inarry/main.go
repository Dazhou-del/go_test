package main

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"reflect"
)

func main() {
	JsonData := "{\"switch_value\":\"2\",\"server_id\":[\"19698\",\"199\",\"15505\",\"19701\"]}"
	jsonQuery := gojsonq.New().FromString(JsonData)
	serverIdList := jsonQuery.Find("server_id")
	str := "19701"
	if stringSlice, ok := serverIdList.([]interface{}); ok {
		var stringSliceResult []string
		for _, item := range stringSlice {
			if str, ok := item.(string); ok {
				stringSliceResult = append(stringSliceResult, str)
			}
		}
		fmt.Println(stringSliceResult)
		if InArray(str, stringSliceResult) {
			fmt.Println("sww")
		}
	}

}

func InArray(d interface{}, arr interface{}) bool {
	return In(d, arr)
}

func In(d interface{}, arr interface{}) bool {
	dValue := reflect.ValueOf(d)
	arrValue := reflect.ValueOf(arr)

	return inByRValue(dValue, arrValue)
}

// inByRValue in data by reflect value
func inByRValue(dV reflect.Value, arrV reflect.Value) bool {
	dVKind, dvType := indirect(dV)
	arrType, arrElemType, _ := indirectForArr(arrV)
	if dvType != arrElemType {
		// d does not consist with arr elem
		return false
	}

	isExist := false
	switch dVKind {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.Struct:
		isExist = inDeepEqual(dV, arrV, arrType)
	default:
		isExist = inEqual(dV, arrV, arrType)
	}

	return isExist
}

func indirect(rv reflect.Value) (reflect.Kind, string) {
	// get reflect value and type
	rvValue, rvType := indirectSimple(rv)

	return rvValue.Kind(), rvType.String()
}

func indirectSimple(rv reflect.Value) (reflect.Value, reflect.Type) {
	// check valid
	if !rv.IsValid() {
		panic("indirectSimple: reflect.Value is nil")
	}

	// get reflect value and type
	var rvValue reflect.Value
	var rvType reflect.Type
	switch rv.Kind() {
	case reflect.Ptr:
		rvValue = rv.Elem()
		rvType = rv.Type().Elem()
	default:
		rvValue = rv
		rvType = rv.Type()
	}

	return rvValue, rvType
}

func indirectForArr(rv reflect.Value) (reflect.Kind, string, string) {
	// get reflect value and type
	rvValue, rvType := indirectSimple(rv)

	vType := rvValue.Kind()
	var vKeyType, vElemType string
	switch vType {
	case reflect.Slice, reflect.Array:
		vKeyType = reflect.Int.String()
		vElemType = rvType.Elem().String()
	case reflect.Map:
		vKeyType = rvType.Key().String()
		vElemType = rvType.Elem().String()
	default:
		panic("haystack: v type must be slice, array or map")
	}

	return vType, vElemType, vKeyType
}

func inDeepEqual(dV reflect.Value, arrV reflect.Value, arrT reflect.Kind) bool {
	isExist := false
	dV = ptrToElem(dV)     // check ptr
	arrV = ptrToElem(arrV) // check ptr
	switch arrT {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrV.Len(); i++ {
			if isExist = reflect.DeepEqual(dV.Interface(), arrV.Index(i).Interface()); isExist {
				break
			}
		}
	case reflect.Map:
		for _, k := range arrV.MapKeys() {
			if isExist = reflect.DeepEqual(dV.Interface(), arrV.MapIndex(k).Interface()); isExist {
				break
			}
		}
	default:
		panic("haystack: d type must be slice, array or map")
	}

	return isExist
}

func ptrToElem(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	return v
}

func inEqual(dV reflect.Value, arrV reflect.Value, arrT reflect.Kind) bool {
	isExist := false
	dV = ptrToElem(dV)     // check ptr
	arrV = ptrToElem(arrV) // check ptr
	switch arrT {
	case reflect.Slice, reflect.Array:
		for i := 0; i < arrV.Len(); i++ {
			if isExist = dV.Interface() == arrV.Index(i).Interface(); isExist {
				break
			}
		}
	case reflect.Map:
		for _, k := range arrV.MapKeys() {
			if isExist = dV.Interface() == arrV.MapIndex(k).Interface(); isExist {
				break
			}
		}
	default:
		panic("haystack: arrV type must be slice, array or map")
	}

	return isExist
}
