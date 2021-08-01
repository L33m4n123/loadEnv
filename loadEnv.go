package loadEnv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type EnvError struct{
	filename string
	message string
}

func (m *EnvError) Error() string {
	return fmt.Sprintf("Unable to load %s File with error %s", m.filename, m.message)
}

func Load(fileName string) (err error) {
	err = loadFile(fileName)
	if err != nil {
		return err
	}

	return nil
}

func loadFile(fileName string) error {
	envFile, err := os.Open(fileName)
	if err != nil {
		return &EnvError{fileName, err.Error()}
	}

	defer func(envFile *os.File) {
		err := envFile.Close()
		if err != nil {
			return
		}
	}(envFile)

	return parseFile(envFile)
}

func parseFile(file *os.File) (err error) {
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return &EnvError{file.Name(), err.Error()}
	}

	for _, line := range lines {
		if isCommentOrEmpty(&line) {
			continue
		}
		lineErr := parseLine(&line, file)
		if lineErr != nil {
			return lineErr
		}
	}

	return nil
}

func parseLine(line *string, file *os.File) (err error) {
	splitLine := strings.SplitN(*line, "=", 2)
	if len(splitLine) != 2 {
		return &EnvError{file.Name(), "Could not split line. Correct format?"}
	}

	return os.Setenv(splitLine[0], strings.Trim(splitLine[1], "\""))
}

func isCommentOrEmpty(line *string) bool {
	*line = strings.TrimSpace(*line)
	return len(*line) == 0 || strings.HasPrefix(*line, "#")
}