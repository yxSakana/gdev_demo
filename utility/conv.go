package utility

import (
	"reflect"
	"strconv"
	"time"
)

func MustInt(str string) int {
	v, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return v
}

func MustUint(str string) uint {
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(v)
}

func MustUint8(str string) uint8 {
	v, err := strconv.ParseUint(str, 10, 8)
	if err != nil {
		panic(err)
	}
	return uint8(v)
}

func MustInt64(str string) int64 {
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func MustUint64(str string) uint64 {
	v, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func MustFloat64(str string) float64 {
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func MustBool(str string) bool {
	v, err := strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	return v
}

func MustTime(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		panic(err)
	}
	return t
}

func MapToStruct(m map[string]interface{}, outputPtr interface{}) error {
	outVal := reflect.ValueOf(outputPtr).Elem()
	for i := 0; i < outVal.NumField(); i++ {
		filed := outVal.Type().Field(i)
		redisTag := filed.Tag.Get("redis")

		if v, ok := m[redisTag]; ok {
			fileVal := outVal.Field(i)
			switch fileVal.Kind() {
			case reflect.String:
				fileVal.SetString(v.(string))
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				fileVal.SetInt(int64(MustInt(v.(string))))
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				fileVal.SetUint(uint64(MustUint(v.(string))))
			case reflect.Float32, reflect.Float64:
				fileVal.SetFloat(MustFloat64(v.(string)))
			case reflect.Bool:
				fileVal.SetBool(MustBool(v.(string)))
			default:
				panic("unhandled default case")
			}
		}
	}
	return nil
}
