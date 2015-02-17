package mbapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

//RenderQuery creates a URL Path for a Query object,
//If you need a Query that is not Provied by this Package you can Create your own, look at query.go for Examles.
func RenderQuery(query Query) (queryUrl string, err error) {
	v := reflect.Indirect(reflect.ValueOf(query))
	t := v.Type()

	queryUrl = query.Path()
	fields := make([]string, 0)

	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := v.Field(i)

		if ft.Type.Name() != "QueryField" {
			continue
		}

		if fv.IsNil() {
			if ft.Tag.Get("optional") != "true" {
				return "", fmt.Errorf("Requred field not set: %s", ft.Name)
			}
			continue
		}

		if fv.Elem().Type().String() != ft.Tag.Get("type") {
			return "", fmt.Errorf("Field %s has wrong Type:\n\tGot: %s Expected: %s\n", ft.Name, fv.Elem().Type().String(), ft.Tag.Get("type"))
		}

		switch ft.Tag.Get("encoding") {

		case "path": //Simple String replacemen uses {{Key}} syntax
			if err := isValueAllowed(fv.Elem(), ft.Tag); err != nil {
				return "", err
			}
			queryUrl = strings.Replace(queryUrl, fmt.Sprintf("{{%s}}", ft.Name), fmt.Sprintf("%s", fv.Elem()), -1)

		case "para": //Field will be encoded into Query parametes(?Key=Value&Key=Value)
			var fs string
			switch fv.Elem().Type().Kind() {
			case reflect.Slice:
				if ft.Tag.Get("sep") == "" {
					return "", fmt.Errorf("Field %s is a Slice(%s) but has no Seperator", ft.Name, fv.Elem().Type())
				}
				for j := 0; j < fv.Elem().Len(); j++ {
					if err := isValueAllowed(fv.Elem().Index(j), ft.Tag); err != nil {
						return "", err
					}
				}

				fs = fmt.Sprintf("%s=%s", ft.Name, fv.Elem().Index(0))
				for j := 1; j < fv.Elem().Len(); j++ {
					fs = fmt.Sprintf("%s%s%s", fs, ft.Tag.Get("sep"), fv.Elem().Index(j).String())
				}
			case reflect.Bool:
				if fv.Elem().Bool() {
					fs = fmt.Sprintf("%s=true", ft.Name)
				} else {
					fs = fmt.Sprintf("%s=false", ft.Name)
				}

			default:
				if err := isValueAllowed(fv.Elem(), ft.Tag); err != nil {
					return "", err
				}
				fs = fmt.Sprintf("%s=%s", ft.Name, fv.Elem())
			}

			fields = append(fields, fs)

		default:
			return "", fmt.Errorf("Invalid encoding for Field %s, valid encodings are: para, path", ft.Name)
		}

	}
	if len(fields) > 0 {
		queryUrl = fmt.Sprintf("%s?%s", queryUrl, fields[0])
	}
	for i := 1; i < len(fields); i++ {
		queryUrl = fmt.Sprintf("%s&%s", queryUrl, fields[i])
	}

	return queryUrl, nil
}

func isValueAllowed(val reflect.Value, tag reflect.StructTag) error {
	// fmt.Printf("Checking Value: %v\n", val)
	if tag.Get("options") != "" {
		allowed := strings.Split(tag.Get("options"), ",")
		ok := false
		for _, a := range allowed {
			if a == val.String() {
				ok = true
			}
		}
		if !ok {
			return fmt.Errorf("Field has not allowed Value: \n\tGot: %s Expected: %s\n", val.String(), allowed)
		}

	}
	return nil
}

//RenderResponse is basicly just a Wrapper around json.Unmarshal
func RenderResponse(resp *http.Response, out Response) error {
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Http call failed: %s", resp.Status)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// log.Printf("Http Body: %s\n", body)
	if err != nil {
		return fmt.Errorf("Failed to Render Body: %v", err)
	}

	if err := json.Unmarshal(body, out); err != nil {
		return fmt.Errorf("Failed to Parse JSON Response: %v", err)
	}
	return nil
}
