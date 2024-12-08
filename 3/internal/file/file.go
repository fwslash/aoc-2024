package file

import (
	"os"
)


func ReadFile(filePath string) string {
    f, err := os.ReadFile(filePath)

    if err != nil {
        return ""
    }

    return string(f)
}
