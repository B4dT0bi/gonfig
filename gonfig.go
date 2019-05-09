// Package gonfig implements simple configuration reading
// from both YAML files and environment variables.
package gonfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"log"

	"github.com/ghodss/yaml"
)

// tag name to override the field name of an environment variable
const envTagName = "env"
const argTagName = "arg"
const defaultTagName = "default"

// GetConf aggregates all the YAML and environment variable values
// and puts them into the passed interface.
func GetConf(configuration interface{}) (err error) {
	GetConfByFilename(getProgramName()+".yaml", configuration)
	return
}

// GetConfByFilename aggregates all the YAML and environment variable values
// and puts them into the passed interface.
func GetConfByFilename(filename string, configuration interface{}) (err error) {

	configValue := reflect.ValueOf(configuration)
	if typ := configValue.Type(); typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("configuration should be a pointer to a struct type")
	}

	setDefaults(configuration)
	getFromYAML(filename, configuration)
	getFromArguments(configuration)
	getFromEnvVariables(configuration)

	return
}

func getProgramName() string {

	splitArg := strings.Split(os.Args[0], "\\")
	if len(splitArg) <= 1 {
		splitArg = strings.Split(os.Args[0], "/")
	}
	return strings.TrimSuffix(splitArg[len(splitArg)-1], ".exe")
}

func setDefaults(configuration interface{}) (err error) {
	getFromEnvVariablesOrArguments(defaultTagName, getFromDefault, configuration)
	return
}

func getFromYAML(filename string, configuration interface{}) (err error) {

	if len(filename) == 0 {
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Println("Could not open file : " + filename + " skipping reading config from YAML.")
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Could not read from file : " + filename + " skipping reading config from YAML.")
		return
	}
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		log.Println("Could not unmarschal from file : " + filename + " skipping extracting config from YAML.")
		return
	}

	return
}

func getFromArguments(configuration interface{}) {
	getFromEnvVariablesOrArguments(argTagName, getFromArg, configuration)
}

func getFromEnvVariables(configuration interface{}) {
	getFromEnvVariablesOrArguments(envTagName, getFromEnv, configuration)
}

type getData func(reflect.StructField, string) string

func getFromDefault(p reflect.StructField, key string) string {
	tagContent := p.Tag.Get(defaultTagName)
	if len(tagContent) > 0 {
		return tagContent
	}
	return ""
}

func getFromEnv(p reflect.StructField, key string) string {
	return os.Getenv(key)
}

func getFromArg(p reflect.StructField, key string) string {
	for i := range os.Args {
		if strings.HasPrefix(os.Args[i], ("--" + key + "=")) {
			return os.Args[i][len(key)+3:]
		} else if os.Args[i] == ("--" + key) {
			if len(os.Args) > i+1 {
				if strings.HasPrefix(os.Args[i+1], "-") {
					return "true"
				}
				return os.Args[i+1]
			}
			return "true"
		}
	}
	return ""
}

func getFromEnvVariablesOrArguments(tagName string, fnGetData getData, configuration interface{}) {
	typ := reflect.TypeOf(configuration)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)

		// check if we've got a field name override for the environment
		tagContent := p.Tag.Get(tagName)
		value := ""
		if len(tagContent) > 0 {
			value = fnGetData(p, tagContent)
		} else {
			value = fnGetData(p, p.Name)
		}

		if !p.Anonymous && len(value) > 0 {
			// struct
			s := reflect.ValueOf(configuration).Elem()

			if s.Kind() == reflect.Struct {
				// exported field
				f := s.FieldByName(p.Name)
				if f.IsValid() && f.CanSet() {
					// A Value can be changed only if it is
					// addressable and was not obtained by
					// the use of unexported struct fields.

					// change value
					kind := f.Kind()
					if kind == reflect.Int || kind == reflect.Int64 {
						setStringToInt(f, value, 64)
					} else if kind == reflect.Int32 {
						setStringToInt(f, value, 32)
					} else if kind == reflect.Int16 {
						setStringToInt(f, value, 16)
					} else if kind == reflect.Uint || kind == reflect.Uint64 {
						setStringToUInt(f, value, 64)
					} else if kind == reflect.Uint32 {
						setStringToUInt(f, value, 32)
					} else if kind == reflect.Uint16 {
						setStringToUInt(f, value, 16)
					} else if kind == reflect.Bool {
						setStringToBool(f, value)
					} else if kind == reflect.Float64 {
						setStringToFloat(f, value, 64)
					} else if kind == reflect.Float32 {
						setStringToFloat(f, value, 32)
					} else if kind == reflect.String {
						f.SetString(value)
					}

				}
			}
		}
	}
}

func setStringToInt(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseInt(value, 10, bitSize)

	if err == nil {
		if !f.OverflowInt(convertedValue) {
			f.SetInt(convertedValue)
		}
	}
}

func setStringToUInt(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseUint(value, 10, bitSize)

	if err == nil {
		if !f.OverflowUint(convertedValue) {
			f.SetUint(convertedValue)
		}
	}
}

func setStringToBool(f reflect.Value, value string) {
	convertedValue, err := strconv.ParseBool(value)

	if err == nil {
		f.SetBool(convertedValue)
	}
}

func setStringToFloat(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseFloat(value, bitSize)

	if err == nil {
		if !f.OverflowFloat(convertedValue) {
			f.SetFloat(convertedValue)
		}
	}
}
