package lib_test

import (
	gitr "github.com/swarupdonepudi/gitr/lib"
	"testing"
)

func TestGitrConfig(t *testing.T) {
	gc1 := gitr.GitrConfig{
		ScmSystems: []gitr.ScmSystem{{
			Hostname: "gitlab.mycompany.com",
			Provider: gitr.GitLab,
			DefaultBranch: "main",
		}},
	}

	t.Run("scm provider lookup using gitlab.com", func(t *testing.T) {
		scm, _ := gc1.GetScmSystem("gitlab.com")
		if scm.Provider != gitr.GitLab {
			t.Errorf("expecting %s provider and got %s", gitr.GitLab, scm.Provider)
		}
	})
	t.Run("scm provider lookup using github.com", func(t *testing.T) {
		scm, _ := gc1.GetScmSystem("github.com")
		if scm.Provider != gitr.GitHub {
			t.Errorf("expecting %s provider and got %s", gitr.GitHub, scm.Provider)
		}
	})
	t.Run("scm provider lookup using bitbucket.org", func(t *testing.T) {
		scm, _ := gc1.GetScmSystem("bitbucket.org")
		if scm.Provider != gitr.BitBucketCloud {
			t.Errorf("expecting %s provider and got %s", gitr.BitBucketCloud, scm.Provider)
		}
	})
	t.Run("scm provider lookup using for custom internal host", func(t *testing.T) {
		scm, _ := gc1.GetScmSystem("gitlab.mycompany.com")
		if scm.Provider != gitr.GitLab {
			t.Errorf("expecting %s provider and got %s", gitr.GitLab, scm.Provider)
		}
	})
	t.Run("default branch for gitlab.com should be main", func(t *testing.T) {
		scm, _ := gc1.GetScmSystem("gitlab.com")
		if scm.DefaultBranch != "main" {
			t.Errorf("expecting default branch to be %s and got %s", "main", scm.DefaultBranch)
		}
	})
}
