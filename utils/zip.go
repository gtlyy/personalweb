package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(zipPath, dest string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0755)
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		w, _ := os.Create(fpath)
		io.Copy(w, rc)
		w.Close()
		rc.Close()
	}
	return nil
}
