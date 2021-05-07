package lib_test

import (
	gitr "github.com/swarupdonepudi/gitr/lib"
	"testing"
)

func TestGetClonePath(t *testing.T) {
	c := gitr.GitrClone{}
	gc := &gitr.GitrConfig{
		Clone: gitr.GitrCloneConfig{
			ScmHome:              "/User/john",
			AlwaysCreDir:         true,
			IncludeHostForCreDir: false,
		},
	}

	testClonePaths := []struct {
		url          string
		scmHome      string
		creDir       bool
		includeHost  bool
		expectedPath string
	}{
		{"git@github.com:swarupdonepudi/gitr.git", "/User/john", false, false, "/User/john/gitr"},
		{"git@github.com:swarupdonepudi/gitr.git", "/User/john", true, false, "/User/john/swarupdonepudi/gitr"},
		{"git@github.com:swarupdonepudi/gitr.git", "/User/john", true, true, "/User/john/github.com/swarupdonepudi/gitr"},
		{"git@github.com:swarupdonepudi/gitr.git", "", true, true, "github.com/swarupdonepudi/gitr"},
		{"git@github.com:swarupdonepudi/gitr.git", "", true, false, "swarupdonepudi/gitr"},
	}

	t.Run("test clone paths should be as per the clone config", func(t *testing.T) {
		for _, p := range testClonePaths {
			gc.Clone.ScmHome = p.scmHome
			gc.Clone.AlwaysCreDir = p.creDir
			gc.Clone.IncludeHostForCreDir = p.includeHost
			c.Url = p.url
			c.Gc = gc
			if c.GetClonePath() != p.expectedPath {
				t.Errorf("expected %s path got %s", p.expectedPath, c.GetClonePath())
			}
		}
	})
}
