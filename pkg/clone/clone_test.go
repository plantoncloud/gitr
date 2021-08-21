package clone

import (
	"testing"
)

func TestScmHome(t *testing.T) {
	var tests = []struct {
		scmHostHomeDir string
		scmHomeDir     string
		expectedVal    string
		expectErr      error
	}{
		{scmHostHomeDir: "/home/john/scm/host", scmHomeDir: "", expectedVal: "/home/john/scm/host", expectErr: nil},
		{scmHostHomeDir: "", scmHomeDir: "/home/john/scm", expectedVal: "/home/john/scm", expectErr: nil},
		{scmHostHomeDir: "/home/john/scm/host/custom/path", scmHomeDir: "/home/john/scm", expectedVal: "/home/john/scm/host/custom/path", expectErr: nil},
	}
	t.Run("validate scm home evaluation", func(t *testing.T) {
		for _, test := range tests {
			returnedVal, err := getScmHome(test.scmHostHomeDir, test.scmHomeDir)
			if err != test.expectErr {
				t.Errorf("expecting %v err but got %v err", test.expectErr, err)
			}
			if returnedVal != test.expectedVal {
				t.Errorf("expecting %s but got %s", test.expectedVal, returnedVal)
			}
		}
	})
}
