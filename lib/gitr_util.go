package lib

import (
	"regexp"
	"strings"
)

func isGitSshUrl(repoUrl string) bool {
	return strings.HasPrefix(repoUrl, "ssh://") || strings.HasPrefix(repoUrl, "git@")
}

func isGitHttpUrlHasUsername(repoUrl string) bool {
	matched, err := regexp.MatchString("https*:\\/\\/.*@+.*", repoUrl)
	if err != nil {
		println(err.Error())
		return false
	} else {
		return matched
	}
}
