package tools

import (
	"os/exec"
	"runtime"
)

func OpenLocalHost(URL string) error {

	var cmd *exec.Cmd //-> Declare External Command outside of switch to be able to use it later

	//-> Prepare External Command depending on OS
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", URL)
	case "windows":
		cmd = exec.Command("start", URL)
	case "darwin":
		cmd = exec.Command("open", URL)
	default:
		return nil //-> In case of Unsupported OS
	}

	//-> Run External Command
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
