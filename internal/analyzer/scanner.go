package analyzer

import (
	"os"
	"path/filepath"
)

// ScanResult holds the files and directories discovered during the FS scan phase.
type ScanResult struct {
	Root        string
	Files       []string
	Directories []string
}

// walking the file system from root.
func Scan(root string) (*ScanResult, error) {
	result := &ScanResult{
		Root:        root,
		Files:       []string{},
		Directories: []string{},
	}

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// skip unnecessary dirs
			if info.Name() == ".git" || info.Name() == "node_modules" {
				return filepath.SkipDir
			}
			result.Directories = append(result.Directories, path)
			return nil
		}

		// collect all files
		// later phases I'll decide which ones to process based on language
		result.Files = append(result.Files, path)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
