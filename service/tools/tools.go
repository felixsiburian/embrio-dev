// all function that help code, register here

package tools

import (
	"crypto/md5"
	"crypto/sha1"
	"embrio-dev/service"
	"encoding/hex"
	"errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

type toolRepository struct {
}

func NewToolRepository() service.IToolsRepository {
	return toolRepository{}
}

func (t toolRepository) SaltAndHash(pwd string) (s string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New("[repository][tool][SaltAndHash] while Generate From Password")
		log.Print(err)
		return s, err
	}

	return string(bytes), err
}

func (t toolRepository) VerifyPassword(hashedPwd, pwd string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		log.Print(err)
		err = errors.New("[repository][tool][VerifyPassword] while CompareHashAndPassword")
		return err
	}

	return err
}

func (t toolRepository) CheckEmailValidation(email string) (res bool, err error) {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(email) < 3 && len(email) > 254 {
		err = errors.New("email length is invalid")
		log.Print(err)
		return false, err
	}

	if !emailRegex.MatchString(email) {
		err = errors.New("email char is invalid")
		log.Print(err)
		return false, err
	}

	parts := strings.Split(email, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		err = errors.New("email parts len is invalid")
		log.Print(err)
		return false, err
	}

	return true, err
}

func (t toolRepository) GUID() string {
	guid := uuid.NewV1()
	guid = uuid.NewV4()

	return guid.String()
}

func CheckArr(a []string, s string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}

	return false
}

func Chains(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func Env(key string, def ...string) string {
	return Chains(append([]string{os.Getenv(key)}, def...)...)
}

func MD5(val string) string {
	mdcp := md5.New()
	_, err := mdcp.Write([]byte(val))
	if err != nil {
		return ""
	}

	return hex.EncodeToString(mdcp.Sum(nil))
}

func SHA1(val string) string {
	sha := sha1.New()
	_, err := sha.Write([]byte(val))
	if err != nil {
		return ""
	}

	return hex.EncodeToString(sha.Sum(nil))
}
