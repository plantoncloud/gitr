package clone_test

import (
	"github.com/swarupdonepudi/gitr/v2/pkg/clone"
	"testing"
)

func TestGetClonePath(t *testing.T) {
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
			cp := clone.GetClonePath(p.url, p.scmHome, p.creDir, p.includeHost)
			if cp != p.expectedPath {
				t.Errorf("expected %s path got %s", p.expectedPath, cp)
			}
		}
	})
}
