package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

func PopulateQueryParams(_ context.Context, req *http.Request, queryParams interface{}, globals interface{}) error {
	values := url.Values{}

	globalsAlreadyPopulated, err := populateQueryParams(queryParams, globals, values, []string{})
	if err != nil {
		return err
	}

	if globals != nil {
		_, err = populateQueryParams(globals, nil, values, globalsAlreadyPopulated)
		if err != nil {
			return err
		}
	}

	req.URL.RawQuery = values.Encode()

	return nil
}

func populateQueryParams(queryParams interface{}, globals interface{}, values url.Values, skipFields []string) ([]string, error) {
	queryParamsStructType, queryParamsValType := dereferencePointers(reflect.TypeOf(queryParams), reflect.ValueOf(queryParams))

	globalsAlreadyPopulated := []string{}

	for i := 0; i < queryParamsStructType.NumField(); i++ {
		fieldType := queryParamsStructType.Field(i)
		valType := queryParamsValType.Field(i)

		if contains(skipFields, fieldType.Name) {
			continue
		}

		requestTag := getRequestTag(fieldType)
		if requestTag != nil {
			continue
		}

		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			continue
		}

		if globals != nil {
			var globalFound bool
			fieldType, valType, globalFound = populateFromGlobals(fieldType, valType, queryParamTagKey, globals)
			if globalFound {
				globalsAlreadyPopulated = append(globalsAlreadyPopulated, fieldType.Name)
			}
		}

		if qpTag.Serialization != "" {
			vals, err := populateSerializedParams(qpTag, fieldType.Type, valType)
			if err != nil {
				return nil, err
			}
			for k, v := range vals {
				values.Add(k, v)
			}
		} else {
			switch qpTag.Style {
			case "deepObject":
				vals := populateDeepObjectParams(qpTag, fieldType.Type, valType)
				for k, v := range vals {
					for _, vv := range v {
						values.Add(k, vv)
					}
				}
			case "form":
				vals := populateFormParams(qpTag, fieldType.Type, valType, ",")
				for k, v := range vals {
					for _, vv := range v {
						values.Add(k, vv)
					}
				}
			case "pipeDelimited":
				vals := populateFormParams(qpTag, fieldType.Type, valType, "|")
				for k, v := range vals {
					for _, vv := range v {
						values.Add(k, vv)
					}
				}
			default:
				return nil, fmt.Errorf("unsupported style: %s", qpTag.Style)
			}
		}
	}

	return globalsAlreadyPopulated, nil
}

func populateSerializedParams(tag *paramTag, objType reflect.Type, objValue reflect.Value) (map[string]string, error) {
	if isNil(objType, objValue) {
		return nil, nil
	}

	if objType.Kind() == reflect.Pointer {
		objValue = objValue.Elem()
	}

	values := map[string]string{}

	switch tag.Serialization {
	case "json":
		data, err := json.Marshal(objValue.Interface())
		if err != nil {
			return nil, fmt.Errorf("error marshaling json: %v", err)
		}
		values[tag.ParamName] = string(data)
	}

	return values, nil
}

func populateDeepObjectParams(tag *paramTag, objType reflect.Type, objValue reflect.Value) url.Values {
	values := url.Values{}

	if isNil(objType, objValue) {
		return values
	}

	if objType.Kind() == reflect.Pointer {
		objType = objType.Elem()
		objValue = objValue.Elem()
	}

	switch objType.Kind() {
	case reflect.Struct:
		for i := 0; i < objType.NumField(); i++ {
			fieldType := objType.Field(i)
			valType := objValue.Field(i)

			if isNil(fieldType.Type, valType) {
				continue
			}

			if fieldType.Type.Kind() == reflect.Pointer {
				valType = valType.Elem()
			}

			qpTag := parseQueryParamTag(fieldType)
			if qpTag == nil {
				continue
			}

			switch valType.Kind() {
			case reflect.Array, reflect.Slice:
				for i := 0; i < valType.Len(); i++ {
					values.Add(fmt.Sprintf("%s[%s]", tag.ParamName, qpTag.ParamName), valToString(valType.Index(i).Interface()))
				}
			default:
				values.Add(fmt.Sprintf("%s[%s]", tag.ParamName, qpTag.ParamName), valToString(valType.Interface()))
			}
		}
	case reflect.Map:
		iter := objValue.MapRange()
		for iter.Next() {
			switch iter.Value().Kind() {
			case reflect.Array, reflect.Slice:
				for i := 0; i < iter.Value().Len(); i++ {
					values.Add(fmt.Sprintf("%s[%s]", tag.ParamName, iter.Key().String()), valToString(iter.Value().Index(i).Interface()))
				}
			default:
				values.Add(fmt.Sprintf("%s[%s]", tag.ParamName, iter.Key().String()), valToString(iter.Value().Interface()))
			}
		}
	}

	return values
}

func populateFormParams(tag *paramTag, objType reflect.Type, objValue reflect.Value, delimiter string) url.Values {
	return populateForm(tag.ParamName, tag.Explode, objType, objValue, delimiter, func(fieldType reflect.StructField) string {
		qpTag := parseQueryParamTag(fieldType)
		if qpTag == nil {
			return ""
		}

		return qpTag.ParamName
	})
}

type paramTag struct {
	Style         string
	Explode       bool
	ParamName     string
	Serialization string
}

func parseQueryParamTag(field reflect.StructField) *paramTag {
	return parseParamTag(queryParamTagKey, field, "form", true)
}
