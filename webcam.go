package webcamchecker

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func getWebcamPid(ctx context.Context) (int, error) {
	cmdStr := " ps -u _cmiodalassistants | grep VDCAssistant | awk '{print $2}'"
	b, err := exec.CommandContext(ctx, "/bin/bash", "-c", fmt.Sprintf("echo $(%s)", cmdStr)).Output()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.Replace(string(b), "\n", "", -1))
}

func getCPUForPid(ctx context.Context, pid int) (float64, error) {
	cmdStr := fmt.Sprintf("ps -p %d -o %%cpu | grep -v %%CPU | awk '{print $1}'", pid)
	b, err := exec.CommandContext(ctx, "/bin/bash", "-c", fmt.Sprintf("echo $(%s)", cmdStr)).Output()
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(strings.Replace(string(b), "\n", "", -1), 64)

}

// IsWebcamOn returns if the device's webcam is on or off
func IsWebcamOn(ctx context.Context) (bool, error) {
	pid, err := getWebcamPid(ctx)
	if err != nil {
		return false, err
	}

	cpu, err := getCPUForPid(ctx, pid)
	if err != nil {
		return false, err
	}

	return cpu > 0.0, nil
}
