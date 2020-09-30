// +build dev

package bundle

import (
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/aclisp/godashboard/backend"
)

var Assets = WasmServeDir{http.Dir("frontend/html/")}

var logger = backend.Logger

type WasmServeDir struct {
	http.Dir
}

func (d WasmServeDir) Open(name string) (http.File, error) {
	if name == "/main.wasm" {
		if err := d.compile(name); err != nil {
			logger.Errorf("can not compile %v: %v", name, err)
		}
	}
	return d.Dir.Open(name)
}

func (d WasmServeDir) compile(name string) error {
	fpath := filepath.Join(string(d.Dir), name)
	if _, err := os.Stat(fpath); err != nil && !os.IsNotExist(err) {
		return err
	}
	// go build
	args := []string{"build", "-o", fpath, "frontend/frontend.go"}
	logger.Print("go ", strings.Join(args, " "))
	cmdBuild := exec.Command("go", args...)
	cmdBuild.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	// If GO111MODULE is not specified explicilty, enable Go modules.
	// Enabling this is for backward compatibility of wasmserve.
	if !hasGo111Module(cmdBuild.Env) {
		cmdBuild.Env = append(cmdBuild.Env, "GO111MODULE=on")
	}
	cmdBuild.Dir = "."
	out, err := cmdBuild.CombinedOutput()
	if err != nil {
		logger.Print(string(out))
		return err
	}
	if len(out) > 0 {
		logger.Print(string(out))
	}
	return nil
}

func hasGo111Module(env []string) bool {
	for _, e := range env {
		if strings.HasPrefix(e, "GO111MODULE=") {
			return true
		}
	}
	return false
}
