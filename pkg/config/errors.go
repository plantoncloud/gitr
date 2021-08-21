package config

import "fmt"

type UnknownScmHostErr struct {
	ScmHost string
}

func (r *UnknownScmHostErr) Error() string {
	return fmt.Sprintf("unknown scm host %s", r.ScmHost)
}
