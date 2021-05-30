package file

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func IsFileExists(f string) bool {
	info, err := os.Stat(f)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GetAbsPath(f string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if strings.HasPrefix(f, "~/") {
		f = filepath.Join(dir, f[2:])
	}
	return f
}

func GetPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}
