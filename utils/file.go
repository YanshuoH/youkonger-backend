package utils

import "path/filepath"

func GetAbsFilePath(file string) (string, error) {
	if !filepath.IsAbs(file) {
		f, e := filepath.Abs(file)
		return f, e
	}
	return file, nil
}
