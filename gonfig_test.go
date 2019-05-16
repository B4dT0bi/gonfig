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

	_, err = io.Copy(file, strings.NewReader(content))
	if err != nil {
		t.Error("Error writing test data", err)
	}

	return file.Name()
}

func createFileWithContent(filename string, content string, t *testing.T) string {

	err := ioutil.WriteFile(filename, []byte(content), 0644)

	if err != nil {
		t.Error("Error creating file with test data", err)
	}

	return filename
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

func Test_GetFromYAML_Filename_Can_Not_Read(t *testing.T) {

	filename := tmpFileWithContent("ID: abc", t)
	defer os.Remove(filename)

	type Conf struct {
		ID int
	}
	conf := Conf{}
	err := getFromYAML(filename, &conf)

	if err == nil {
		t.Error("getFromYAML should throw an error", err)
	}
	if conf.ID != 0 {
		t.Error("ID should be 0", conf.ID)
	}
}

func Test_GetFromYAML_Filename_Can_Read(t *testing.T) {

	filename := tmpFileWithContent("ID: 123\nTestString: hallo", t)
	defer os.Remove(filename)

	type Conf struct {
		ID         int
		TestString string
	}
	conf := Conf{}
	err := getFromYAML(filename, &conf)

	if err != nil {
		t.Error("getFromYAML file not found", err)
	}
	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.TestString != "hallo" {
		t.Error("TestString should be hallo", conf.ID)
	}
}

func Test_GetConfByFilename_Can_Read(t *testing.T) {

	filename := tmpFileWithContent("ID: 123\nTestString: hallo", t)
	defer os.Remove(filename)

	type Conf struct {
		ID         int
		TestString string
	}
	conf := Conf{}
	err := GetConfByFilename(filename, &conf)

	if err != nil {
		t.Error("GetConfByFilename unexpected error occured", err)
	}
	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.TestString != "hallo" {
		t.Error("TestString should be hallo", conf.ID)
	}
}

func Test_GetConf_Can_Read(t *testing.T) {

	filename := createFileWithContent("gonfigtest.yaml", "ID: 123\nTestString: hallo\nNotInStruct: this should not cause problems", t)
	oldArgs := os.Args
	os.Args = []string{"gonfigtest", "--MYFLOAT32=123.123", "--MYFLOAT64", "456.456"}
	defer func() {
		os.Args = oldArgs
		os.Remove(filename)
	}()

	type Conf struct {
		ID         int
		TestString string
		MyFloat    float32 `arg:"MYFLOAT32"`
		MYFLOAT64  float64
		NotInYaml  int
	}
	conf := Conf{}
	err := GetConf(&conf)

	if err != nil {
		t.Error("GetConf unexpected error occured", err)
	}
	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.NotInYaml != 0 {
		t.Error("NotInYaml should be 0", conf.NotInYaml)
	}
	if conf.TestString != "hallo" {
		t.Error("TestString should be hallo", conf.TestString)
	}
	if conf.MyFloat != 123.123 {
		t.Error("MyFloat should be 123.123", conf.MyFloat)
	}
	if conf.MYFLOAT64 != 456.456 {
		t.Error("MYFLOAT64 should be 456.456", conf.MYFLOAT64)
	}
}

func Test_getFromEnvVariables_should_not_panic_if_wrong_data(t *testing.T) {
	type Conf struct {
		ID int
	}
	os.Setenv("ID", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 0 {
		t.Error("ID should be 0", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int(t *testing.T) {
	type Conf struct {
		ID int
	}
	os.Setenv("ID", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int16(t *testing.T) {
	type Conf struct {
		ID int16
	}
	os.Setenv("ID", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int32(t *testing.T) {
	type Conf struct {
		ID int32
	}
	os.Setenv("ID", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_int64(t *testing.T) {
	type Conf struct {
		ID int64
	}
	os.Setenv("ID", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint(t *testing.T) {
	type Conf struct {
		ID uint
	}
	os.Setenv("ID", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint16(t *testing.T) {
	type Conf struct {
		ID uint16
	}
	os.Setenv("ID", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint32(t *testing.T) {
	type Conf struct {
		ID uint32
	}
	os.Setenv("ID", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_uint64(t *testing.T) {
	type Conf struct {
		ID uint64
	}
	os.Setenv("ID", "123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_bool(t *testing.T) {
	type Conf struct {
		ID bool
	}
	os.Setenv("ID", "true")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != true {
		t.Error("ID should be true", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_float32(t *testing.T) {
	type Conf struct {
		ID float32
	}
	os.Setenv("ID", "123.123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123.123 {
		t.Error("ID should be 123.123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_float64(t *testing.T) {
	type Conf struct {
		ID float64
	}
	os.Setenv("ID", "123.123")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != 123.123 {
		t.Error("ID should be 123.123", conf.ID)
	}
}

func Test_getFromEnvVariables_should_find_and_parse_string(t *testing.T) {
	type Conf struct {
		ID string
	}
	os.Setenv("ID", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != "abc" {
		t.Error("ID should be abc", conf.ID)
	}
}

func Test_getFromCustomEnvVariables_should_find_and_parse_string(t *testing.T) {
	type Conf struct {
		ID string `env:"CONF_ID"`
	}
	os.Setenv("CONF_ID", "abc")
	conf := Conf{}
	getFromEnvVariables(&conf)

	if conf.ID != "abc" {
		t.Error("ID should be abc", conf.ID)
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
		ID  string `arg:"confID"`
		ID2 string `arg:"confID2"`
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--confID=abc", "--confID2", "def"}
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
		ID  string `arg:"confID"`
		ID2 string `arg:"confID2"`
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--confID=abc"}
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
		ID  int
		ID2 int
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=abc", "--ID2", "def"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 0 {
		t.Error("ID should be 0", conf.ID)
	}
}

func Test_getFromArguments_should_find_and_parse_int(t *testing.T) {
	type Conf struct {
		ID  int
		ID2 int
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123", "--ID2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.ID2 != 456 {
		t.Error("ID2 should be 456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_int16(t *testing.T) {
	type Conf struct {
		ID  int16
		ID2 int16
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123", "--ID2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.ID2 != 456 {
		t.Error("ID2 should be 456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_int32(t *testing.T) {
	type Conf struct {
		ID  int32
		ID2 int32
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123", "--ID2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.ID2 != 456 {
		t.Error("ID2 should be 456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_int64(t *testing.T) {
	type Conf struct {
		ID  int64
		ID2 int64
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123", "--ID2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.ID2 != 456 {
		t.Error("ID2 should be 456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_uint(t *testing.T) {
	type Conf struct {
		ID  uint
		ID2 uint
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123", "--ID2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.ID2 != 456 {
		t.Error("ID2 should be 456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_uint16(t *testing.T) {
	type Conf struct {
		ID  uint16
		ID2 uint16
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123", "--ID2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.ID2 != 456 {
		t.Error("ID2 should be 456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_uint32(t *testing.T) {
	type Conf struct {
		ID  uint32
		ID2 uint32
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123", "--ID2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.ID2 != 456 {
		t.Error("ID2 should be 456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_uint64(t *testing.T) {
	type Conf struct {
		ID  uint64
		ID2 uint64
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123", "--ID2", "456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123 {
		t.Error("ID should be 123", conf.ID)
	}
	if conf.ID2 != 456 {
		t.Error("ID2 should be 456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_bool(t *testing.T) {
	type Conf struct {
		ID  bool
		ID2 bool
		ID3 bool
		ID4 bool
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=true", "--ID2", "true", "--ID3", "--ID4"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != true {
		t.Error("ID should be true", conf.ID)
	}
	if conf.ID2 != true {
		t.Error("ID2 should be true", conf.ID2)
	}
	if conf.ID3 != true {
		t.Error("ID3 should be true", conf.ID3)
	}
	if conf.ID4 != true {
		t.Error("ID4 should be true", conf.ID4)
	}
}

func Test_getFromArguments_should_find_and_parse_float32(t *testing.T) {
	type Conf struct {
		ID  float32
		ID2 float32
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123.123", "--ID2", "456.456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123.123 {
		t.Error("ID should be 123.123", conf.ID)
	}
	if conf.ID2 != 456.456 {
		t.Error("ID should be 456.456", conf.ID2)
	}
}

func Test_getFromArguments_should_find_and_parse_float64(t *testing.T) {
	type Conf struct {
		ID  float64
		ID2 float64
	}
	oldArgs := os.Args
	os.Args = []string{"cmd", "--ID=123.123", "--ID2", "456.456"}
	defer func() { os.Args = oldArgs }()
	conf := Conf{}
	getFromArguments(&conf)

	if conf.ID != 123.123 {
		t.Error("ID should be 123.123", conf.ID)
	}
	if conf.ID2 != 456.456 {
		t.Error("ID should be 456.456", conf.ID2)
	}
}

func Test_getFromDefaults_should_find_and_parse_float32(t *testing.T) {
	type Conf struct {
		ID float32 `default:"123.123"`
	}
	conf := Conf{}
	setDefaults(&conf)

	if conf.ID != 123.123 {
		t.Error("ID should be 123.123", conf.ID)
	}
}
