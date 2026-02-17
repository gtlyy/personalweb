package utils

import (
	"archive/zip"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ErrUnsafePath = errors.New("unsafe file path detected")

func Unzip(zipPath, dest string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)

		// 安全检查：防止路径穿越攻击 (Zip Slip)
		// 确保解压后的路径在目标目录内
		absDest, err := filepath.Abs(dest)
		if err != nil {
			return err
		}
		absFpath, err := filepath.Abs(fpath)
		if err != nil {
			return err
		}
		if !strings.HasPrefix(absFpath, absDest+string(filepath.Separator)) {
			return ErrUnsafePath
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, 0755)
			continue
		}

		// 确保父目录存在
		parentDir := filepath.Dir(fpath)
		if err := os.MkdirAll(parentDir, 0755); err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		w, err := os.Create(fpath)
		if err != nil {
			return err
		}
		defer w.Close()

		if _, err := io.Copy(w, rc); err != nil {
			return err
		}
	}
	return nil
}
