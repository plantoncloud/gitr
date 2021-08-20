package clone

import (
	"testing"
)

func TestScmHome(t *testing.T) {
	var tests = []struct {
		scmHostHomeDir string
		scmHomeDir     string
		expectedVal    string
	}{
		{scmHostHomeDir: "/home/john/scm/host", scmHomeDir: "", expectedVal: "/home/john/scm/host"},
		{scmHostHomeDir: "", scmHomeDir: "/home/john/scm", expectedVal: "/home/john/scm"},
		{scmHostHomeDir: "/home/john/scm/host/custom/path", scmHomeDir: "/home/john/scm", expectedVal: "/home/john/scm/host/custom/path"},
	}
	t.Run("validate scm home evaluation", func(t *testing.T) {
		for _, test := range tests {
			returnedVal := getScmHome(test.scmHostHomeDir, test.scmHomeDir)
			if returnedVal != test.expectedVal {
				t.Errorf("expecting %s but got %s", test.expectedVal, returnedVal)
			}
		}
	})
}
