package utils

import (
	"fmt"
	"os/exec"

	"github.com/mukezhz/geng/pkg"
)

// ModTidy runs go mod tidy command in the specified directory
func ModTidy(dir string) error {
	logger := pkg.GetLogger()

	logger.Info("tidying up mod file")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir

	op, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run go mod tidy. err: %w, %s \n", err, string(op))
	}

	logger.Infof("TIDY OUTPUT:\n%s", string(op))
	logger.Info("go mod tidy completed.")

	return nil
}
