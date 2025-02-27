// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package utils

import (
	"errors"
	"reflect"
	"strings"
)

var (
	DefaultValidateFunction = "Validate"
)

func getTypeName(value string) string {
	values := strings.Split(value, ".")
	if len(values) > 1 {
		values := strings.Split(values[1], " ")
		return values[0]
	} else {
		return values[0]
	}
}

func validateCallbackByValue(data reflect.Value) error {
	method := data.MethodByName(DefaultValidateFunction)
	if method.IsValid() {
		response := method.Call(nil)
		if len(response) > 0 {
			err := response[0]
			if !err.IsNil() {
				typeName := getTypeName(data.String())
				if len(typeName) > 0 {
					errStr := err.Interface().(error).Error()
					if !strings.Contains(errStr, ")") {
						errStr = errStr + " (" + typeName + ")"
					} else {
						errStr = errStr[:len(errStr)-1] + ", " + typeName + ")"
					}
					return errors.New(errStr)
				}
				return err.Interface().(error)
			}
		}
	}
	return nil
}

// to validate interface
func Validate(r interface{}) error {
	var err error
	fields := reflect.ValueOf(r).Elem()
	for i := 0; i < fields.NumField(); i++ {
		fieldData := fields.Field(i)
		kind := fieldData.Kind()
		if kind == reflect.Slice {
			for i := 0; i < fieldData.Len(); i++ {
				err = validateCallbackByValue(fieldData.Index(i))
				if err != nil {
					return err
				}
			}
		} else if kind == reflect.Map {
			for _, key := range fieldData.MapKeys() {
				err = validateCallbackByValue(fieldData.MapIndex(key))
				if err != nil {
					return err
				}
			}
		} else if kind == reflect.Ptr {
			if fieldData.Pointer() != 0 {
				err = validateCallbackByValue(fieldData)
				if err != nil {
					return err
				}
			}
		} else {
			err = validateCallbackByValue(fieldData)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
