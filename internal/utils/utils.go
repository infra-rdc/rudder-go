package utils

import (
	"context"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ericlagergren/decimal"
)

const (
	queryParamTagKey  = "queryParam"
	headerParamTagKey = "header"
	pathParamTagKey   = "pathParam"
)

var (
	paramRegex                       = regexp.MustCompile(`({.*?})`)
	SerializationMethodToContentType = map[string]string{
		"json":      "application/json",
		"form":      "application/x-www-form-urlencoded",
		"multipart": "multipart/form-data",
		"raw":       "application/octet-stream",
		"string":    "text/plain",
	}
)

func UnmarshalJsonFromResponseBody(body io.Reader, out interface{}, tag string) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	if err := UnmarshalJSON(data, out, reflect.StructTag(tag), true, false); err != nil {
		return fmt.Errorf("error unmarshalling json response body: %w", err)
	}

	return nil
}

func ReplaceParameters(stringWithParams string, params map[string]string) string {
	if len(params) == 0 {
		return stringWithParams
	}

	return paramRegex.ReplaceAllStringFunc(stringWithParams, func(match string) string {
		match = match[1 : len(match)-1]
		return params[match]
	})
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func MatchStatusCodes(expectedCodes []string, statusCode int) bool {
	for _, codeStr := range expectedCodes {
		code, err := strconv.Atoi(codeStr)
		if err == nil {
			if code == statusCode {
				return true
			}
			continue
		}

		codeRange, err := strconv.Atoi(string(codeStr[0]))
		if err != nil {
			continue
		}

		if statusCode >= (codeRange*100) && statusCode < ((codeRange+1)*100) {
			return true
		}
	}

	return false
}

func AsSecuritySource(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

func parseStructTag(tagKey string, field reflect.StructField) map[string]string {
	tag := field.Tag.Get(tagKey)
	if tag == "" {
		return nil
	}

	values := map[string]string{}

	options := strings.Split(tag, ",")
	for _, optionConf := range options {
		parts := strings.Split(optionConf, "=")

		switch len(parts) {
		case 1:
			// flag option
			parts = append(parts, "true")
		case 2:
			// key=value option
		default:
			// invalid option
			continue
		}

		values[parts[0]] = parts[1]
	}

	return values
}

func parseParamTag(tagKey string, field reflect.StructField, defaultStyle string, defaultExplode bool) *paramTag {
	// example `{tagKey}:"style=simple,explode=false,name=apiID"`
	values := parseStructTag(tagKey, field)
	if values == nil {
		return nil
	}

	tag := &paramTag{
		Style:     defaultStyle,
		Explode:   defaultExplode,
		ParamName: strings.ToLower(field.Name),
	}

	for k, v := range values {
		switch k {
		case "style":
			tag.Style = v
		case "explode":
			tag.Explode = v == "true"
		case "name":
			tag.ParamName = v
		case "serialization":
			tag.Serialization = v
		}
	}

	return tag
}

func valToString(val interface{}) string {
	switch v := val.(type) {
	case time.Time:
		return v.Format(time.RFC3339Nano)
	case big.Int:
		return v.String()
	case decimal.Big:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func populateFromGlobals(fieldType reflect.StructField, valType reflect.Value, paramType string, globals interface{}) (reflect.StructField, reflect.Value, bool) {
	if globals == nil {
		return fieldType, valType, false
	}

	globalsStruct := reflect.TypeOf(globals)
	globalsStructVal := reflect.ValueOf(globals)

	globalsField, found := globalsStruct.FieldByName(fieldType.Name)
	if !found {
		return fieldType, valType, false
	}

	if fieldType.Type.Kind() != reflect.Ptr || !valType.IsNil() {
		return fieldType, valType, true
	}

	globalsVal := globalsStructVal.FieldByName(fieldType.Name)

	if !globalsVal.IsValid() {
		return fieldType, valType, false
	}

	switch paramType {
	case queryParamTagKey:
		qpTag := parseQueryParamTag(globalsField)
		if qpTag == nil {
			return fieldType, valType, false
		}
	default:
		tag := parseParamTag(paramType, fieldType, "simple", false)
		if tag == nil {
			return fieldType, valType, false
		}
	}

	return globalsField, globalsVal, true
}

func isNil(typ reflect.Type, val reflect.Value) bool {
	// `reflect.TypeOf(nil) == nil` so calling typ.Kind() will cause a nil pointer
	// dereference panic. Catch it and return early.
	// https://github.com/golang/go/issues/51649
	// https://github.com/golang/go/issues/54208
	if typ == nil {
		return true
	}

	if typ.Kind() == reflect.Ptr || typ.Kind() == reflect.Map || typ.Kind() == reflect.Slice || typ.Kind() == reflect.Interface {
		return val.IsNil()
	}

	return false
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}