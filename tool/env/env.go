package env

import (
	"github.com/yydsorg/goyyds/tool/nfs/homedir"
	"log"
	"os"
	"path/filepath"
)

func GetENV(k, defaultV string) string {
	env := os.Getenv(k)
	if env == "" {
		env = defaultV
	}
	return env
}

var (
	basePath = ""
)

func GetBasePath() string {
	if basePath == "" {
		path := GetENV("BASE_PATH", homedir.BasePath)
		abspath, err := filepath.Abs(path)
		if err != nil {
			log.Println("BASE_PATH has wrong")
		}
		basePath = abspath
	}
	return basePath
}

func GetPublicPath() string {
	return GetBasePath() + "/pub"
}
func GetPrivatePath() string {
	return GetBasePath() + "/pri"
}
func GetTempPath() string {
	return GetBasePath() + "/temp"
}

func GetPublicUrlPath() string {
	return GetENV("Protocol", "http") + "://" + GetENV("Host", "localhost") + ":" + GetENV("Port", "8080") + "/static"
}
