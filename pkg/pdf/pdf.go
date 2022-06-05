package pdf

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
)

func Merge(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	pdfFiles := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) == ".pdf" {
			pdfFiles = append(pdfFiles, file.Name())
			fmt.Printf("found %s\n", file.Name())
		}
	}
	if len(pdfFiles) == 0 {
		return errors.New("no pdf files in the current directory")
	}
	pdf.MergeCreateFile(pdfFiles, "merged.pdf", nil)
	fmt.Println("merged.pdf created!")
	return err
}
