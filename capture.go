package testx

import (
	"os"
)

func CaptureStdout(fn func()) (string, error) {
	old := os.Stdout
	defer func() {
		os.Stdout = old
	}()

	fakeOut, err := os.CreateTemp(os.TempDir(), "stdout-*")
	if err != nil {
		return "", err
	}

	defer os.Remove(fakeOut.Name())
	defer fakeOut.Close()

	os.Stdout = fakeOut
	fn()

	stdoutBytes, err := os.ReadFile(fakeOut.Name())
	if err != nil {
		return "", err
	}

	return string(stdoutBytes), nil
}

func InputStdin(stdin string, fn func()) error {
	fakeIn, err := os.CreateTemp(os.TempDir(), "stdin-*")
	if err != nil {
		return err
	}
	if err := os.WriteFile(fakeIn.Name(), []byte(stdin), 0o600); err != nil {
		return err
	}
	os.Stdin = fakeIn
	fn()
	if err := os.Remove(fakeIn.Name()); err != nil {
		return err
	}
	return nil
}
