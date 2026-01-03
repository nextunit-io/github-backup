package backup

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func (b *backup) zip(outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	w := zip.NewWriter(file)
	defer w.Close()

	walker := func(path string, info os.FileInfo, err error) error {
		outputPath := strings.TrimPrefix(path, b.workingDir+string(os.PathSeparator))
		if b.verbose {
			log.Printf("Crawling: %#v for output %v\n", path, outputPath)
		}
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Ensure that `path` is not absolute; it should not start with "/".
		// This snippet happens to work because I don't use
		// absolute paths, but ensure your real-world code
		// transforms path into a zip-root relative path.
		f, err := w.Create(outputPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, file)
		if err != nil {
			return err
		}

		return nil
	}
	return filepath.Walk(b.workingDir, walker)
}
