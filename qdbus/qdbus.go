package qdbus

import (
	"fmt"
	"os/exec"
	"strings"
)

type QDBusClient struct {
	svc, path, iface string
}

func NewQDBusClient(svc, path, iface string) *QDBusClient {
	return &QDBusClient{svc, path, iface}
}

func (c *QDBusClient) cmd(method string, args ...any) (*exec.Cmd, error) {
	path, err := exec.LookPath("qdbus6")
	if err != nil {
		return nil, fmt.Errorf("qdbus6 not found: %w", err)
	}

	cmdArgs := []string{c.svc, c.path, c.iface + "." + method}
	for _, arg := range args {
		cmdArgs = append(cmdArgs, fmt.Sprintf("%v", arg))
	}

	return exec.Command(path, cmdArgs[1:]...), nil
}

func (c *QDBusClient) run(method string, args ...any) error {
	if cmd, err := c.cmd(method, args...); err != nil {
		return err
	} else if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to call method %s: %w", method, err)
	} else {
		return nil
	}
}

func (c *QDBusClient) output(method string, args ...any) (string, error) {
	if cmd, err := c.cmd(method, args...); err != nil {
		return "", err
	} else if output, err := cmd.Output(); err != nil {
		return "", fmt.Errorf("failed to call method %s: %w", method, err)
	} else {
		return string(output), nil
	}
}

func (c *QDBusClient) outputLines(method string, args ...any) ([]string, error) {
	if output, err := c.output(method, args...); err != nil {
		return nil, err
	} else {
		lines := strings.Split(output, "\n")
		for i := range lines {
			lines[i] = strings.TrimSpace(lines[i])
		}
		return lines, nil
	}
}
