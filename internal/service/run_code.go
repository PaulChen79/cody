package service

import (
	"cody/domain"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func (svc *service) RunCode(codeContent string, language domain.CodeLanguageType) (string, error) {
	tempPath := svc.config.TempCodePath + "/go"
	filePath := tempPath + "/main.go"

	err := os.WriteFile(filePath, []byte(codeContent), 0644)
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "docker", "run", "--rm", "--memory=10m", "--cpus=1.0",
		"-v", fmt.Sprintf("%s:/app", tempPath), "go-runner")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
