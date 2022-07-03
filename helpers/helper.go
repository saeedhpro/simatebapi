package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func TimeDiff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}

func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	if arr.Kind() == reflect.Slice || arr.Kind() == reflect.Array {
		for i := 0; i < arr.Len(); i++ {
			if arr.Index(i).Interface() == item {
				return true
			}
		}
	}

	return false
}

func SaveImageToDisk(location string, names []string, data string) (string, string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", "", fmt.Errorf("invalid image")
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		log.Println("errpeed")
		return "", "", err
	}
	//imgCfg, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	//if err != nil {
	//	log.Println("errpeed 2")
	//	return "", err
	//}
	//
	//if imgCfg.Width != 750 || imgCfg.Height != 685 {
	//	return "", fmt.Errorf("invalid size")
	//}
	name := ""
	for {
		name = RandomString(8)
		if !ItemExists(names, fmt.Sprintf("%s.jpg", name)) {
			break
		}
	}
	fileName := fmt.Sprintf("%s/%s.jpg", location, name)
	err = ioutil.WriteFile(fileName, buff.Bytes(), 0644)
	if err != nil {
		fmt.Println(err.Error(), "cf")
		return "", "", fmt.Errorf("cant save file")
	}
	return fileName, name, err
}

func NormalizePhoneNumber(number string) string {
	match, _ := regexp.MatchString("(\\+98)9\\d{9}", number)
	if match {
		return number
	} else {
		if len(number) == 11 {
			n := fmt.Sprintf("+98%s", number[1:])
			return n
		} else {
			return ""
		}
	}
}
