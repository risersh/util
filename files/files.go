package files

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

// FileExists checks if the file exists at the given file path.
// It returns true if the file exists, otherwise false.
// This function does not check if the file is a directory.
func GetFileSize(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		return -1
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return -1
	}

	return stat.Size()
}

// FileExists checks if the file exists at the given file path.
// It returns true if the file exists, otherwise false.
func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}

// WaitForFileExists waits for the file to exist at the given file path.
// It returns true if the file exists within the specified timeout, otherwise false.
// This function periodically checks for the file existence.
func WaitForFileExists(filePath string, timeout time.Duration) bool {
	// Create a timer for the timeout
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	// Create a ticker for periodically checking the file existence
	checkInterval := 100 * time.Millisecond
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-timer.C:
			// Timeout occurred
			return false
		case <-ticker.C:
			if _, err := os.Stat(filePath); err == nil {
				// File exists
				return true
			} else if !os.IsNotExist(err) {
				// An error other than "not exist", stop waiting
				return false
			}
			// If file does not exist, continue checking
		}
	}
}

// WaitForNoFileHandlers waits for all file handlers to be closed for the given file path.
// It returns true if all file handlers are closed within the specified timeout, otherwise false.
// This function uses the `lsof` command to check for open file handlers.
func WaitForNoFileHandlers(filePath string, timeout time.Duration, local bool) bool {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		var cmd *exec.Cmd
		if local {
			cmd = exec.Command("lsof", filePath)
		} else {
			cmd = exec.Command("sh", "-c", fmt.Sprintf("lsof | grep %s", filePath))
		}

		var out bytes.Buffer
		cmd.Stdout = &out

		err := cmd.Run()
		if err != nil {
			// lsof returns an error if no file handlers are found.
			return true
		}

		if out.Len() == 0 {
			// Also check if the output is empty, indicating no open file handlers.
			return true
		}

		time.Sleep(100 * time.Millisecond) // Wait before trying again.
	}

	return false // Timeout reached.
}

func MoveFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = sourceFile.Close()
	if err != nil {
		return err
	}

	err = os.Remove(src)
	if err != nil {
		return err
	}

	return nil
}
