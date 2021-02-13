package obj

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/remrain/weakjson"
	"os"
	"path/filepath"
	"pure-init/lib/log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func PrettyPrint(data interface{}) {
	fmt.Println(PrettyEncode(data))
}

func PrettyEncode(data interface{}) string {
	encoded, _ := weakjson.MarshalIndent(data, "", "  ")
	return string(encoded)
}

func Encode(data interface{}) string {
	encoded, _ := weakjson.Marshal(data)
	return string(encoded)
}

func Decode(data interface{}, encoded string) error {
	return weakjson.Unmarshal([]byte(encoded), data)
}

func Copy(to interface{}, from interface{}) error {
	encoded, err := weakjson.Marshal(from)
	if err != nil {
		return err
	}
	err = weakjson.Unmarshal(encoded, to)
	return err
}
func ReflectCopy(dst interface{}, src interface{}) (err error) {
	dstValue := reflect.ValueOf(dst)
	if dstValue.Kind() != reflect.Ptr {
		err = errors.New("dst isn't a pointer to struct")
		return
	}
	dstElem := dstValue.Elem()
	if dstElem.Kind() != reflect.Struct {
		err = errors.New("pointer doesn't point to struct")
		return
	}

	srcValue := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)
	if srcType.Kind() != reflect.Struct {
		err = errors.New("src isn't struct")
		return
	}

	for i := 0; i < srcType.NumField(); i++ {
		sf := srcType.Field(i)
		sv := srcValue.FieldByName(sf.Name)
		// make sure the value which in dst is valid and can set
		if dv := dstElem.FieldByName(sf.Name); dv.IsValid() && dv.CanSet() {
			dv.Set(sv)
		}
	}
	return
}
func Contains(s interface{}, elem interface{}) bool {
	arrV := reflect.ValueOf(s)

	if arrV.Kind() == reflect.Slice {
		for i := 0; i < arrV.Len(); i++ {

			// XXX - panics if slice element points to an unexported struct field
			// see https://golang.org/pkg/reflect/#Value.Interface
			if arrV.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

func ToString(v int) string {
	return strconv.Itoa(v)
}

func ToInt(v string) int {
	value, _ := strconv.Atoi(v)
	return value
}

func ImplodeRepeatString(str string, delimiter string, count int) string {
	repeatArray := ArrayRepeatString(str, count)
	return strings.Join(repeatArray, delimiter)
}

func ArrayRepeatString(str string, count int) []string {
	ret := []string{}
	for i := 0; i < count; i++ {
		ret = append(ret, str)
	}
	return ret
}

func ArrayRepeatInt(v int, count int) []int {
	ret := []int{}
	for i := 0; i < count; i++ {
		ret = append(ret, v)
	}
	return ret
}

var camel = regexp.MustCompile("(^[^A-Z]*|[A-Z]*)([A-Z][^A-Z]+|$)")

// Convert "HelloWorld" to "hello_world"
func CamelToUnderScore(s string) string {
	var a []string
	for _, sub := range camel.FindAllStringSubmatch(s, -1) {
		if sub[1] != "" {
			a = append(a, sub[1])
		}
		if sub[2] != "" {
			a = append(a, sub[2])
		}
	}
	return strings.ToLower(strings.Join(a, "_"))
}

func ChangeToWorkingDir() error {
	selfProgram, err := filepath.Abs(os.Args[0])
	if err != nil {
		return err
	}
	workingDir := filepath.Dir(filepath.Dir(selfProgram))
	if err := os.Chdir(workingDir); err != nil {
		return err
	}
	return nil
}

func Retry(desc string, retryTimes int, method func() error) (err error) {
	for i := 1; i < retryTimes+1; i++ {
		err = method()
		if err == nil {
			return
		} else if err == sql.ErrNoRows {
			log.Warnf("%s no rows found: %s", desc, err)
			return
		}
		log.Warnf("%s failed: %s, retry %d times", desc, err, i)
		time.Sleep(time.Duration(i*i) * time.Second)
	}
	log.Warnf("%s failed after retry %d times, last error: %s", desc, retryTimes, err)
	return err
}

func IsValidPtr(data interface{}) bool {
	v := reflect.ValueOf(data)
	return (v.Kind() == reflect.Ptr && !v.IsNil())
}
