package pkg

import (
	"github.com/pkg/errors"
	"os"
)

func Remove(path string) error {
	if err := os.Remove(path); err != nil {
		return errors.Wrapf(err, "failed to delete the last folder in dir %s", path)
	}
	return nil
}
