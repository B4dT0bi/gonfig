package gonfig

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func Test_GetFromYAML_Filename_Empty_Should_Not_Panic(t *testing.T) {

	type Conf struct {
	}
	conf := Conf{}
	err := getFromYAML("", &conf)

	if err != nil {
		t.Error("getFromYAML should not panic", err)
	}
}

func tmpFileWithContent(content string, t *testing.T) string {

	file, err := ioutil.TempFile("", "gonfig_test_data_")
	if err != nil {
		t.Error("Error creating file with test data", err)
	}

	_, err = io.Copy(file, strings.NewReader("{}"))
	if err != nil {
		t.Error("Error writing test data", err)
	}

	return file.Name()
}

func Test_GetFromYAML_Filename_Should_Not_be_Panic(t *testing.T) {

	filename := tmpFileWithContent("{}", t)
	defer os.Remove(filename)

	type Conf struct {
	}
	conf := Conf{}
	err := getFromYAML(filename, &conf)

	if err != nil {
		t.Error("getFromYAML file not found", err)
	}
}

func Test_getFromEnvVariables_should_not_panic_if_wrong_data(t *testing.T) {
	type Conf struct {
		Id int
	}
	os.Setenv("Id", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 0 {
		t.Error("Id should be 0", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int(t *testing.T) {
	type Conf struct {
		Id int
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int16(t *testing.T) {
	type Conf struct {
		Id int16
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int32(t *testing.T) {
	type Conf struct {
		Id int32
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int64(t *testing.T) {
	type Conf struct {
		Id int64
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint(t *testing.T) {
	type Conf struct {
		Id uint
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint16(t *testing.T) {
	type Conf struct {
		Id uint16
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint32(t *testing.T) {
	type Conf struct {
		Id uint32
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint64(t *testing.T) {
	type Conf struct {
		Id uint64
	}
	os.Setenv("Id", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_bool(t *testing.T) {
	type Conf struct {
		Id bool
	}
	os.Setenv("Id", "true")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != true {
		t.Error("Id should be true", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_float32(t *testing.T) {
	type Conf struct {
		Id float32
	}
	os.Setenv("Id", "123.123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123.123 {
		t.Error("Id should be 123.123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_float64(t *testing.T) {
	type Conf struct {
		Id float64
	}
	os.Setenv("Id", "123.123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != 123.123 {
		t.Error("Id should be 123.123", conf.Id)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_string(t *testing.T) {
	type Conf struct {
		Id string
	}
	os.Setenv("Id", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != "abc" {
		t.Error("Id should be abc", conf.Id)
	}
}

func Test_getFromCustomEnvVariables_should_find_and_parse_string(t *testing.T) {
	type Conf struct {
		Id string `env:"CONF_ID"`
	}
	os.Setenv("CONF_ID", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.Id != "abc" {
		t.Error("Id should be abc", conf.Id)
	}
}

func Test_getFromArguments_should_find_and_parse_string(t *testing.T) {
	type Conf struct {
		ID  string
		ID2 string
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=abc", "--ID2", "def"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != "abc" {
		t.Error("ID should be abc but is : ", conf.ID)
	}
	if conf.ID2 != "def" {
		t.Error("ID2 should be def but is : ", conf.ID2)
	}
}

func Test_getFromCustomArguments_should_find_and_parse_string(t *testing.T) {
	type Conf struct {
		ID  string `arg:"confId"`
		ID2 string `arg:"confId2"`
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--confId=abc", "--confId2", "def"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != "abc" {
		t.Error("ID should be abc but is : ", conf.ID)
	}
	if conf.ID2 != "def" {
		t.Error("ID2 should be def but is : ", conf.ID2)
	}
}

func Test_getFromCustomArguments_should_not_find_string(t *testing.T) {
	type Conf struct {
		ID  string `arg:"confId"`
		ID2 string `arg:"confId2"`
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--confId=abc"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	conf.ID2 = "def"
	getFromArguments(&conf)

	if conf.ID != "abc" {
		t.Error("ID should be abc but is : ", conf.ID)
	}
	if conf.ID2 != "def" {
		t.Error("ID2 should be def but is : ", conf.ID2)
	}
}

func Test_getFromArguments_should_not_panic_if_wrong_data(t *testing.T) {
	type Conf struct {
		Id  int
		Id2 int
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=abc", "--Id2", "def"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 0 {
		t.Error("Id should be 0", conf.Id)
	}
}

func Test_getFromArguments_should_find_and_parse_int(t *testing.T) {
	type Conf struct {
		Id  int
		Id2 int
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123", "--Id2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
	if conf.Id2 != 456 {
		t.Error("Id2 should be 456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_int16(t *testing.T) {
	type Conf struct {
		Id  int16
		Id2 int16
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123", "--Id2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
	if conf.Id2 != 456 {
		t.Error("Id2 should be 456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_int32(t *testing.T) {
	type Conf struct {
		Id  int32
		Id2 int32
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123", "--Id2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
	if conf.Id2 != 456 {
		t.Error("Id2 should be 456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_int64(t *testing.T) {
	type Conf struct {
		Id  int64
		Id2 int64
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123", "--Id2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
	if conf.Id2 != 456 {
		t.Error("Id2 should be 456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_uint(t *testing.T) {
	type Conf struct {
		Id  uint
		Id2 uint
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123", "--Id2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
	if conf.Id2 != 456 {
		t.Error("Id2 should be 456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_uint16(t *testing.T) {
	type Conf struct {
		Id  uint16
		Id2 uint16
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123", "--Id2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
	if conf.Id2 != 456 {
		t.Error("Id2 should be 456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_uint32(t *testing.T) {
	type Conf struct {
		Id  uint32
		Id2 uint32
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123", "--Id2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
	if conf.Id2 != 456 {
		t.Error("Id2 should be 456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_uint64(t *testing.T) {
	type Conf struct {
		Id  uint64
		Id2 uint64
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123", "--Id2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123 {
		t.Error("Id should be 123", conf.Id)
	}
	if conf.Id2 != 456 {
		t.Error("Id2 should be 456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_bool(t *testing.T) {
	type Conf struct {
		Id  bool
		Id2 bool
		Id3 bool
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=true", "--Id2", "true", "--Id3"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != true {
		t.Error("Id should be true", conf.Id)
	}
	if conf.Id2 != true {
		t.Error("Id2 should be true", conf.Id2)
	}
	if conf.Id3 != true {
		t.Error("Id3 should be true", conf.Id3)
	}
}

func Test_getFromArguments_should_find_and_parse_float32(t *testing.T) {
	type Conf struct {
		Id  float32
		Id2 float32
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123.123", "--Id2", "456.456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123.123 {
		t.Error("Id should be 123.123", conf.Id)
	}
	if conf.Id2 != 456.456 {
		t.Error("Id should be 456.456", conf.Id2)
	}
}

func Test_getFromArguments_should_find_and_parse_float64(t *testing.T) {
	type Conf struct {
		Id  float64
		Id2 float64
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--Id=123.123", "--Id2", "456.456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.Id != 123.123 {
		t.Error("Id should be 123.123", conf.Id)
	}
	if conf.Id2 != 456.456 {
		t.Error("Id should be 456.456", conf.Id2)
	}
}
