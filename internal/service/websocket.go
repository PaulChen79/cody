package service

import (
	"bufio"
	"cody/domain"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type RunCodeMsg struct {
	CodeContent string                  `json:"code_content" binding:"required"`
	Language    domain.CodeLanguageType `json:"language" binding:"required"`
}

func (svc *service) HandleRunCodeWsMessage(ws *websocket.Conn) {
	for {
		var msg RunCodeMsg
		err := ws.ReadJSON(&msg)
		if err != nil {
			zap.S().Warnf("HandleRunCodeWsMessage - Error occurred while reading msg: %v", err)
			return
		}

		err = ws.WriteJSON("Running code...")
		if err != nil {
			zap.S().Warnf("HandleRunCodeWsMessage - Error occurred while sending message: %v", err)
			return
		}
		// codeContent := "package main\n\nimport (\n\t\"fmt\"\n\t\"time\"\n)\n\nfunc main() {\n\ttime.Sleep(20 * time.Second)\n\tfmt.Println(\"Hello, playground\")\n}"
		tempPath := "/Users/paul/Paul/Programimg/cody/temp_code_files/go"
		filePath := tempPath + "/main.go"

		err = os.WriteFile(filePath, []byte(msg.CodeContent), 0644)
		if err != nil {
			zap.S().Warnf("HandleRunCodeWsMessage - err: %v", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		cmd := exec.CommandContext(ctx, "docker", "run", "--rm", "--memory=200m", "--cpus=1.0",
			"-v", fmt.Sprintf("%s:/app", tempPath), "go-runner")

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			zap.S().Warnf("HandleRunCodeWsMessage - err: %v", err)
			return
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			zap.S().Warnf("HandleRunCodeWsMessage - err: %v", err)
			return
		}

		if err := cmd.Start(); err != nil {
			zap.S().Warnf("HandleRunCodeWsMessage - err: %v", err)
			return
		}

		scanner := bufio.NewScanner(stdout)
		go func() {
			for scanner.Scan() {
				err := ws.WriteJSON(scanner.Text())
				if err != nil {
					zap.S().Warnf("HandleRunCodeWsMessage - Error occurred while sending message: %v", err)
					return
				}
			}
		}()

		errScanner := bufio.NewScanner(stderr)
		go func() {
			for errScanner.Scan() {
				err := ws.WriteJSON(errScanner.Text())
				if err != nil {
					zap.S().Warnf("HandleRunCodeWsMessage - Error occurred while sending error message: %v", err)
					return
				}
			}
		}()

		if err := cmd.Wait(); err != nil {
			if ctx.Err() == context.DeadlineExceeded {
				err := ws.WriteJSON("Time exceed limit")
				if err != nil {
					zap.S().Warnf("HandleRunCodeWsMessage - Error occurred while sending time exceed message: %v", err)
					return
				}
			}
		}

		err = ws.WriteJSON("Program execution finished")
		if err != nil {
			zap.S().Warnf("HandleRunCodeWsMessage - Error occurred while sending program finish message: %v", err)
			return
		}
	}
}
