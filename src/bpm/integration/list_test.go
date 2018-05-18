package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestList(t *testing.T) {
	s, err := NewSandbox()
	if err != nil {
		t.Fatalf("sandbox setup failed: %v", err)
	}
	defer s.Cleanup()

	cmd := s.BPMCmd("start", "blah")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run bpm: %s", output)
	}
	fmt.Println(string(output))

	// cmd := s.BPMCmd("list")
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	t.Fatalf("failed to run bpm: %s", output)
	// }
	// fmt.Println(string(output))
}

type Sandbox struct {
	bpmExe  string
	runcExe string

	root string
}

var bpmExe = flag.String("bpmexe", "", "path to the bpm executable to test")
var runcExe = flag.String("runcexe", "", "path to the runc executable to test")

func NewSandbox() (*Sandbox, error) {
	if *bpmExe == "" {
		return nil, errors.New("-bpmexe must specified")
	}

	if *runcExe == "" {
		return nil, errors.New("-runcexe must specified")
	}

	root, err := ioutil.TempDir("", "bpm_sandbox")
	if err != nil {
		return nil, fmt.Errorf("could not create sandbox root directory: %v", err)
	}

	if err := os.MkdirAll(filepath.Join(root, "packages", "bpm", "bin"), 0700); err != nil {
		return nil, fmt.Errorf("could not create sandbox directory structure: %v", err)
	}

	runcSandboxPath := filepath.Join(root, "packages", "bpm", "bin", "runc")
	if err := os.Symlink(*runcExe, runcSandboxPath); err != nil {
		return nil, fmt.Errorf("could not link runc executable into sandbox: %v", err)
	}
	if err := os.Chmod(*runcExe, 0755); err != nil {
		return nil, fmt.Errorf("could not make runc executable +x: %v", err)
	}

	return &Sandbox{
		bpmExe:  *bpmExe,
		runcExe: *runcExe,
		root:    root,
	}, nil
}

func (s *Sandbox) BPMCmd(args ...string) *exec.Cmd {
	cmd := exec.Command(*bpmExe, args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("BPM_BOSH_ROOT=%s", s.root))
	return cmd
}

func (s *Sandbox) Cleanup() {
	_ = os.RemoveAll(s.root)
}
