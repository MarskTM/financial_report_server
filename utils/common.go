package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/golang/glog"
	"github.com/lithammer/shortuuid"
	"golang.org/x/crypto/bcrypt"
)

// GenerateKey using in set keys
func GenCode() string {
	id := shortuuid.New()
	return strings.ToUpper(id[0:10])
}

// PatternGet using in get keys
func PatternGet(id uint) string {
	return strconv.Itoa(int(id)) + "-:--*"
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// ---------------------------utils func suport for folder-------------------------------------
func GetRootPath() string {
	root, err := os.Getwd()
	if err != nil {
		glog.V(1).Info("Could not get root path: %v", err)
	}
	return root
}

func GetPublicPath() string {
	return GetRootPath() + "/public"
}

// ----------------------------- utils func suport for authen password -----------------------------------
func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		glog.V(1).Info("HashAndSalt - err: ", err)
	}
	return string(hash)
}

func ComparePassword(hashedPwd string, plainPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd)); err != nil {
		return false
	}
	return true
}

// GenerateKey random password
func GeneratePasswordKey(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer)[:length], nil
}
