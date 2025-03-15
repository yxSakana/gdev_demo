package utility

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yxSakana/gdev_demo/settings"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileType int

const (
	CoverFt FileType = iota
	ImageFt
)

var ErrFileHeaderIsNil = errors.New("file is nil")

func SaveFile(c *gin.Context, file *multipart.FileHeader, ft FileType) (filePath string, err error) {
	if file == nil {
		return "", ErrFileHeaderIsNil
	}
	filePath, err = GenerateFilePath(file)
	if err != nil {
		return "", err
	}

	switch ft {
	case CoverFt:
		if err := CheckCoverFile(file, filePath); err != nil {
			return "", err
		}
	case ImageFt:
		if err := CheckImageFile(file, filePath); err != nil {
			return "", err
		}
	}

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", err
	}
	return filePath, nil
}

func GenerateFilePath(file *multipart.FileHeader) (string, error) {
	hashName, err := hashFileHeader(file)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s%s",
		settings.Settings.Server.UploadDir,
		hashName, filepath.Ext(file.Filename)), nil
}

func CheckCoverFile(file *multipart.FileHeader, dst string) error {
	maxSize := settings.Settings.Server.CoverFileMaxSize
	return CheckFile(file, dst, maxSize)
}

func CheckImageFile(file *multipart.FileHeader, dst string) error {
	maxSize := settings.Settings.Server.ImageFileMaxSize
	return CheckFile(file, dst, maxSize)
}

func CheckFile(file *multipart.FileHeader, dst string, maxSize int64) error {
	if file == nil {
		return errors.New("file is nil")
	}

	if file.Size > maxSize {
		return errors.New("file too big")
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return errors.New("file type not supported")
	}

	err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func hashFileHeader(file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}
