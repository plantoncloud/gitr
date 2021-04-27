package lib_test

import (
	"github.com/swarupdonepudi/gitr/lib"
	"testing"
)

func TestGetRemUrl(t *testing.T) {
	var getRemUrlTests = []struct {
		remote      string
		expectedUrl string
	}{
		{"git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tree/master"},
		{"https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tree/master"},
	}
	r := lib.RemoteRepo{
		Scheme:   "https",
		Provider: "github",
		Branch:   "master",
	}
	for _, u := range getRemUrlTests {
		r.Url = u.remote
		if r.GetRemUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetRemUrl())
		}
	}
}
