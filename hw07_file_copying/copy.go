package main

import (
	"errors"
	"io"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrSrcFileNotFound       = errors.New("source file not found")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Открываем файл на чтение.
	src, err := os.OpenFile(fromPath, os.O_RDONLY, 0)
	if err != nil {
		// Существует ли файл источника?
		if os.IsNotExist(err) {
			return ErrSrcFileNotFound
		}
		return err
	}
	defer src.Close()

	// Проверяем, не превысил ли offset размер файла в байтах.
	stat, err := src.Stat()
	if err != nil {
		return err
	}
	size := stat.Size()
	if size < offset {
		return ErrOffsetExceedsFileSize
	}

	// Проверяем, поддерживается ли файл.
	if !stat.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	// Создаем результирующий файл.
	out, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Если лимит на чтение не выставлен, то читаем до EOF.
	// Настраиваем progressbar.
	barSize := limit
	if limit == 0 {
		limit = size
		barSize = size - offset
	}
	if limit+offset > size {
		barSize = size - offset
	}
	// Запускаем progressbar. Проксируем в него src io.Reader.
	bar := pb.Full.Start64(barSize)
	barReader := bar.NewProxyReader(src)
	// Выставляем offset и копируем src в out через progressbar.
	src.Seek(offset, 0)
	io.CopyN(out, barReader, limit)
	bar.Finish()

	return nil
}
