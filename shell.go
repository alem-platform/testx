package testx

import (
	"context"
	"os"
	"time"
)

func (p *Program) RunScript(ctx context.Context, args ...string) (string, error) {
	if err := os.Chdir(p.path); err != nil {
		return "", err
	}
	stdout, err := Exec(ctx, ExecInput{
		Command: "bash",
		Args:    args,
		Timeout: time.Second * 15,
	})
	if err != nil {
		return "", err
	}
	return stdout, nil
}

func (p *Program) CPTestFile(ctx context.Context, file, newPath string) error {
	if err := os.Chdir(p.path); err != nil {
		return err
	}
	if _, err := Exec(ctx, ExecInput{
		Command: "cp",
		Args:    []string{"-r", file, newPath},
	}); err != nil {
		return err
	}
	return nil
}
