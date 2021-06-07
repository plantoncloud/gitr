package config

type GitrConfig struct {
	CopyRepoPathCdCmdToClipboard bool `yaml:"copyRepoPathCdCmdToClipboard"`
	Scm                          Scm  `yaml:"scm"`
}

type Scm struct {
	Hosts   []ScmHost `yaml:"hosts"`
	HomeDir string    `yaml:"homeDir"`
}

type ScmHost struct {
	Hostname      string      `yaml:"hostname"`
	Provider      ScmProvider `yaml:"provider"`
	DefaultBranch string      `yaml:"defaultBranch"`
	Clone         CloneConfig `yaml:"clone"`
	Scheme        HttpScheme  `yaml:"scheme"`
}

type CloneConfig struct {
	HomeDir              string `yaml:"homeDir"`
	AlwaysCreDir         bool   `yaml:"alwaysCreDir"`
	IncludeHostForCreDir bool   `yaml:"includeHostForCreDir"`
}
