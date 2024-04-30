package testx

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"time"
)

type Program struct {
	path     string
	compiled bool
}

func NewProgram(dirs ...string) Program {
	return Program{
		path:     path.Join(dirs...),
		compiled: false,
	}
}

func (p *Program) Compile(ctx context.Context) error {
	if p.compiled {
		return nil
	}

	if _, err := Exec(ctx, ExecInput{
		Command: "go",
		Args:    []string{"build", "-o", "main", "."},
	}); err != nil {
		return err
	}

	p.compiled = true
	return nil
}

func (p *Program) Run(ctx context.Context) (string, error) {
	if err := os.Chdir(p.path); err != nil {
		return "", err
	}

	if err := p.Compile(ctx); err != nil {
		return "", err
	}

	return Exec(context.Background(), ExecInput{Command: "./main"})
}

type ExecInput struct {
	Command string
	Args    []string
	Timeout time.Duration
}

func Exec(ctx context.Context, input ExecInput) (string, error) {
	if input.Timeout == 0 {
		input.Timeout = time.Second * 5
	}

	ctx, cancel := context.WithTimeout(ctx, input.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, input.Command, input.Args...)
	cmd.Cancel = func() error {
		err := cmd.Process.Kill()
		return err
	}
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	if stderr.Len() > 0 {
		return "", fmt.Errorf("exec failed: %s", stderr.String())
	}

	return stdout.String(), nil
}
