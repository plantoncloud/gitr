package web

import (
	"fmt"
	"strings"

	"github.com/plantoncloud/gitr/pkg/config"
)

// GetFileURL returns the browser URL for a single file in the repo.
//
//	base   – repo web URL, e.g. https://github.com/org/repo
//	ref    – branch name or commit SHA
//	rel    – path inside repo, **forward-slash** format
//
// Provider-specific rules:
//
//	GitHub             : <base>/blob/<ref>/<rel>
//	GitLab             : <base>/-/blob/<ref>/<rel>
//	Bitbucket Cloud/DC : <base>/src/<ref>/<rel>
func GetFileURL(p config.ScmProvider, base, ref, rel string) string {
	rel = strings.TrimPrefix(rel, "/") // safety

	switch p {
	case config.GitLab:
		return fmt.Sprintf("%s/-/blob/%s/%s", base, ref, rel)
	case config.BitBucketCloud, config.BitBucketDatacenter:
		return fmt.Sprintf("%s/src/%s/%s", base, ref, rel)
	default: // GitHub and similar
		return fmt.Sprintf("%s/blob/%s/%s", base, ref, rel)
	}
}
