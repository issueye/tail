//go:build windows
// +build windows

package tail

import (
	"github.com/issueye/tail/winfile"
	"os"
)

func OpenFile(name string) (file *os.File, err error) {
	return winfile.OpenFile(name, os.O_RDONLY, 0)
}
