package clone_test

import (
	"github.com/swarupdonepudi/gitr/v2/pkg/clone"
	"testing"
)

func TestGetClonePath(t *testing.T) {
	testClonePaths := []struct {
		scmHost      string
		repoPath     string
		repoName     string
		scmHome      string
		creDir       bool
		includeHost  bool
		expectedPath string
	}{
		{"github.com", "swarupdonepudi/gitr", "gitr", "/User/john", false, false, "/User/john/gitr"},
		{"github.com", "swarupdonepudi/gitr", "gitr", "/User/john", true, false, "/User/john/swarupdonepudi/gitr"},
		{"github.com", "swarupdonepudi/gitr", "gitr", "/User/john", true, true, "/User/john/github.com/swarupdonepudi/gitr"},
		{"github.com", "swarupdonepudi/gitr", "gitr", "", true, true, "github.com/swarupdonepudi/gitr"},
		{"github.com", "swarupdonepudi/gitr", "gitr", "", true, false, "swarupdonepudi/gitr"},
	}

	t.Run("test clone paths should be as per the clone config", func(t *testing.T) {
		for _, p := range testClonePaths {
			cp := clone.GetClonePath(p.scmHost, p.repoPath, p.repoName, p.scmHome, p.creDir, p.includeHost)
			if cp != p.expectedPath {
				t.Errorf("expected %s path got %s", p.expectedPath, cp)
			}
		}
	})
}
