package utils

import (
	"errors"
	"fmt"
	"reflect"
)

func GetStructStringField(input interface{}, key string) (value string, err error) {
	v, err := getStructField(input, key)
	if err != nil {
		return
	}
	value, ok := v.(string)
	if !ok {
		return value, errors.New("can't convert key'v to string")
	}
	return
}

func GetStructFloat64Field(input interface{}, key string) (value float64, err error) {
	v, err := getStructField(input, key)
	if err != nil {
		return
	}
	value, ok := v.(float64)
	if !ok {
		return value, errors.New("can't convert key'v to float64")
	}
	return
}

func GetStrucIntField(input interface{}, key string) (value int, err error) {
	v, err := getStructField(input, key)
	if err != nil {
		return
	}
	value, ok := v.(int)
	if !ok {
		return value, errors.New("can't convert key'v to int")
	}
	return
}

func getStructField(input interface{}, key string) (value interface{}, err error) {
	rv := reflect.ValueOf(input)
	rt := reflect.TypeOf(input)
	if rt.Kind() != reflect.Struct {
		return value, errors.New("input must be struct")
	}
	keyExist := false
	for i := 0; i < rt.NumField(); i++ {
		curField := rv.Field(i)
		if rt.Field(i).Name == key {
			switch curField.Kind() {
			case reflect.String, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Float64, reflect.Float32:
				keyExist = true
				value = curField.Interface()
			default:
				return value, errors.New("key must be int float or string")
			}
		}
	}
	if !keyExist {
		return value, errors.New(fmt.Sprintf("key %s not found in %s's field", key, rt))
	}
	return
}
