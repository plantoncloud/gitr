package ssh

import (
	"strings"
	"testing"
)

//Host scm-bitbucket-org
//HostName bitbucket.org
//User git
//IdentityFile /Users/swarup/.ssh/scm/bitbucket.org

func TestGetKeyPath(t *testing.T) {
	keyPathTests := []struct {
		name      string
		sshConfig string
		hostname  string
		expected  string
		err       error
	}{
		{
			name: "alias is same as hostname with single host",
			sshConfig: `
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/scm/github.com
`,
			hostname: "github.com",
			expected: "~/.ssh/scm/github.com",
		}, {
			name: "alias is same as hostname with multiple hosts",
			sshConfig: `
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/scm/github.com

Host gitlab.com
  HostName gitlab.com
  User git
  IdentityFile ~/.ssh/scm/gitlab.com

Host bitbucket.org
  HostName bitbucket.org
  User git
  IdentityFile ~/.ssh/scm/bitbucket.com
`,
			hostname: "github.com",
			expected: "~/.ssh/scm/github.com",
		}, {
			name: "alias is not same as hostname",
			sshConfig: `
Host scm-github-com
  HostName github.com
  User git
  IdentityFile ~/.ssh/scm/github.com
`,
			hostname: "github.com",
			expected: "~/.ssh/scm/github.com",
		}, {
			name: "hostname does not exist",
			sshConfig: `
Host not-github.com
  HostName not-github.com
  User git
  IdentityFile ~/.ssh/scm/github.com
`,
			hostname: "github.com",
			expected: "",
			err:      ErrHostCfgNotFound,
		},
	}
	t.Run("validate key paths", func(t *testing.T) {
		for _, u := range keyPathTests {
			t.Run(u.name, func(t *testing.T) {
				result, err := getKeyPathFromConfig(strings.NewReader(u.sshConfig), u.hostname)
				if err != nil {
					if err != u.err {
						t.Errorf("unexpected error: %v", err)
					}
				} else {
					if result != u.expected {
						t.Errorf("expecting %s but got %s", u.expected, result)
					}
				}
			})
		}
	})
}
