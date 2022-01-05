package ssh

import (
	"bytes"
	"github.com/kevinburke/ssh_config"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ErrHostCfgNotFound = errors.New("host config not found for host")

func GetKeyPath(hostname string) (string, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get ssh config path")
	}
	if !file.IsFileExists(configPath) {
		return "", errors.Errorf("ssh config file %s not found", configPath)
	}
	configReader, err := getConfigReader(configPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get config reader for %s file", configPath)
	}
	keyPath, err := getKeyPathFromConfig(configReader, hostname)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get key path for %s host from config", hostname)
	}
	if keyPath == "" {
		keyPath, err = getDefaultSshKeyPath()
		if err != nil {
			return "", errors.Wrapf(err, "failed to get default ssh path")
		}
	}
	keyFileAbsPath, err := file.GetAbsPath(keyPath)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get abs path of %s", keyPath)
	}
	return keyFileAbsPath, nil
}

func getDefaultSshKeyPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get home dir")
	}
	return filepath.Join(homeDir, ".ssh", "id_rsa"), nil
}

func getConfigReader(configFilePath string) (io.Reader, error) {
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file %s", configFilePath)
	}
	return bytes.NewReader(file), nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get home dir")
	}
	return filepath.Join(homeDir, ".ssh", "config"), nil
}

func getKeyPathFromConfig(sshConfigReader io.Reader, hostname string) (string, error) {
	cfg, err := ssh_config.Decode(sshConfigReader)
	if err != nil {
		return "", errors.Wrapf(err, "failed to decode ssh config")
	}
	hostCfg := getHostConfigByHostname(cfg, hostname)
	if hostCfg == nil {
		return "", ErrHostCfgNotFound
	}
	keyPath, err := getKeyPathFromHostConfig(hostCfg)
	if err != nil {
		return "", errors.Wrapf(err, "failed to get key path from host config")
	}
	return keyPath, nil
}

func getHostConfigByHostname(cfg *ssh_config.Config, hostname string) *ssh_config.Host {
	for _, host := range cfg.Hosts {
		if len(host.Patterns) == 0 {
			continue
		}
		for _, n := range host.Nodes {
			if strings.Contains(n.String(), "HostName") {
				hn := strings.Split(strings.TrimSpace(n.String()), " ")[1]
				if hn == hostname {
					return host
				}
			}
		}
	}
	return nil
}

func getKeyPathFromHostConfig(hostCfg *ssh_config.Host) (string, error) {
	for _, n := range hostCfg.Nodes {
		if strings.Contains(n.String(), "IdentityFile") {
			return strings.Split(strings.TrimSpace(n.String()), " ")[1], nil
		}
	}
	return "", errors.Errorf("identity file not found")
}
