package utils

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func PopulateHeaders(_ context.Context, req *http.Request, headers interface{}, globals interface{}) {
	globalsAlreadyPopulated := populateHeaders(headers, globals, req.Header, []string{})
	if globals != nil {
		_ = populateHeaders(globals, nil, req.Header, globalsAlreadyPopulated)
	}
}

func populateHeaders(headers interface{}, globals interface{}, reqHeaders http.Header, skipFields []string) []string {
	headerParamsStructType, headerParamsValType := dereferencePointers(reflect.TypeOf(headers), reflect.ValueOf(headers))

	globalsAlreadyPopulated := []string{}

	for i := 0; i < headerParamsStructType.NumField(); i++ {
		fieldType := headerParamsStructType.Field(i)
		valType := headerParamsValType.Field(i)

		if contains(skipFields, fieldType.Name) {
			continue
		}

		if globals != nil {
			var globalFound bool
			fieldType, valType, globalFound = populateFromGlobals(fieldType, valType, headerParamTagKey, globals)
			if globalFound {
				globalsAlreadyPopulated = append(globalsAlreadyPopulated, fieldType.Name)
			}
		}

		tag := parseParamTag(headerParamTagKey, fieldType, "simple", false)
		if tag == nil {
			continue
		}

		value := serializeHeader(fieldType.Type, valType, tag.Explode)
		if value != "" {
			reqHeaders.Add(tag.ParamName, value)
		}
	}

	return globalsAlreadyPopulated
}

func serializeHeader(objType reflect.Type, objValue reflect.Value, explode bool) string {
	if isNil(objType, objValue) {
		return ""
	}

	if objType.Kind() == reflect.Pointer {
		objType = objType.Elem()
		objValue = objValue.Elem()
	}

	switch objType.Kind() {
	case reflect.Struct:
		items := []string{}

		for i := 0; i < objType.NumField(); i++ {
			fieldType := objType.Field(i)
			valType := objValue.Field(i)

			if isNil(fieldType.Type, valType) {
				continue
			}

			if fieldType.Type.Kind() == reflect.Pointer {
				valType = valType.Elem()
			}

			tag := parseParamTag(headerParamTagKey, fieldType, "simple", false)
			if tag == nil {
				continue
			}

			fieldName := tag.ParamName

			if fieldName == "" {
				continue
			}

			if explode {
				items = append(items, fmt.Sprintf("%s=%s", fieldName, valToString(valType.Interface())))
			} else {
				items = append(items, fieldName, valToString(valType.Interface()))
			}
		}

		return strings.Join(items, ",")
	case reflect.Map:
		items := []string{}

		iter := objValue.MapRange()
		for iter.Next() {
			if explode {
				items = append(items, fmt.Sprintf("%s=%s", iter.Key().String(), valToString(iter.Value().Interface())))
			} else {
				items = append(items, iter.Key().String(), valToString(iter.Value().Interface()))
			}
		}

		return strings.Join(items, ",")
	case reflect.Slice, reflect.Array:
		items := []string{}

		for i := 0; i < objValue.Len(); i++ {
			items = append(items, valToString(objValue.Index(i).Interface()))
		}

		return strings.Join(items, ",")
	default:
		return valToString(objValue.Interface())
	}
}
