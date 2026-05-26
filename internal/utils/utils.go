package utils

func LoadFile(path string) ([]byte, error) {
	return readFile(path)
}
