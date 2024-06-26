package testx

import (
	"context"
	"path"
	"time"
)

type Workdir struct {
	path string
}

func NewWorkdir(dirs ...string) Workdir {
	return Workdir{
		path: path.Join(dirs...),
	}
}

func (p *Workdir) Bash(ctx context.Context, scriptPath ...string) (string, error) {
	stdout, err := Exec(ctx, ExecInput{
		Command: "bash",
		Args:    []string{path.Join(scriptPath...)},
		Timeout: time.Second * 10,
		Dir:     p.path,
	})
	if err != nil {
		return "", err
	}
	return stdout, nil
}
