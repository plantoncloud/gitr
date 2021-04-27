package lib_test

import (
	"github.com/swarupdonepudi/gitr/lib"
	"testing"
)

func TestGetRemUrl(t *testing.T) {
	r := lib.RemoteRepo{
		Url:      "git@github.com:swarupdonepudi/gitr.git",
		Scheme:   "https",
		Provider: "github",
		Branch:   "master",
	}
	expectedRemUrl := "https://github.com/swarupdonepudi/gitr/tree/master"
	if r.GetRemUrl() != expectedRemUrl {
		t.Errorf("expecting %s but got %s", expectedRemUrl, r.GetRemUrl())
	}
	r.Url = "https://github.com/swarupdonepudi/gitr.git"
	if r.GetRemUrl() != expectedRemUrl {
		t.Errorf("expecting %s but got %s", expectedRemUrl, r.GetRemUrl())
	}
}
