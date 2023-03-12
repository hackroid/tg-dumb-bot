package utils

import "os"

func CloseFile(jsonFile *os.File) {
	_ = jsonFile.Close()
}
