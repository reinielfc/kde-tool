package kscreen

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

func EnableOnly(outputs ...*Output) error {
	toEnable := make(map[string]*Output, len(outputs))
	for _, output := range outputs {
		toEnable[output.Name()] = output
		output.SetEnabled()
	}

	current, err := outputNames()
	if err != nil {
		return fmt.Errorf("failed to get current outputs: %w", err)
	}

	toDisable := make([]*Output, 0, len(current)-len(toEnable))
	for _, name := range current {
		if _, ok := toEnable[name]; !ok {
			output := NewOutput(name)
			output.SetDisabled()

			toDisable = append(toDisable, output)
		}
	}

	return applyCfgs(append(outputs, toDisable...))
}

func applyCfgs(outputs []*Output) error {
	args := make([]string, 0)
	for _, config := range outputs {
		args = append(args, config.BuildArgs()...)
	}

	cmd, err := kscreenCmd(args...)
	if err != nil {
		return fmt.Errorf("failed to create kscreen-doctor command: %w", err)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		if exitErr, ok := errors.AsType[*exec.ExitError](err); ok {
			return fmt.Errorf("kscreen-doctor exited with code %d: %s", exitErr.ExitCode(), string(output))
		}

		return fmt.Errorf("failed to execute kscreen-doctor %q, stderr: %s", err, string(output))
	}

	pattern, err := regexp.Compile(`Output with name or uuid .*? not found.`)
	if err != nil {
		return fmt.Errorf("failed to compile regex pattern: %w", err)
	}

	if len(output) > 0 && pattern.MatchString(string(output)) {
		return fmt.Errorf("output not found: %s", string(output))
	}

	return nil
}

func outputNames() ([]string, error) {
	cmd, err := kscreenCmd("--json")
	if err != nil {
		return nil, fmt.Errorf("failed to create kscreen-doctor command: %w", err)
	}

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute kscreen-doctor: %w", err)
	}

	var data struct {
		Outputs []struct {
			Name string `json:"name"`
		} `json:"outputs"`
	}

	if err := json.Unmarshal(output, &data); err != nil {
		return nil, fmt.Errorf("failed to parse kscreen-doctor output: %w", err)
	}

	names := make([]string, len(data.Outputs))
	for i, output := range data.Outputs {
		names[i] = output.Name
	}

	return names, nil
}

func kscreenCmd(args ...string) (*exec.Cmd, error) {
	path, err := exec.LookPath("kscreen-doctor")
	if err != nil {
		return nil, fmt.Errorf("kscreen-doctor not found: %w", err)
	}

	cmd := exec.Command(path, args...)
	return cmd, nil
}
