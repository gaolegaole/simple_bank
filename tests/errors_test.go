package tests

import (
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestErrors(t *testing.T) {
	/*_, err := readConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}*/
}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "open failed")
		//return nil, errors.WithMessage(err, "open failed")
		//return nil, errors.WithStack(err)
	}

	defer f.Close()

	buf, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	return buf, nil
}

func readConfig() ([]byte, error) {
	home := os.Getenv("HOME")
	config, err := readFile(filepath.Join(home, ".settings.xml"))
	return config, errors.Wrap(err, "could not read config")
}
