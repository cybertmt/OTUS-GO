package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	var limit, offset int64
	toPath := "out.txt"
	fromPath := "testdata/input.txt"

	t.Run("offset exceeds file size", func(t *testing.T) {
		limit = 1000
		offset = 7000

		err := Copy(fromPath, toPath, offset, limit)

		require.Truef(t, errors.Is(err, ErrOffsetExceedsFileSize), "actual err - %v", err)
	})
	t.Run("file does not exist", func(t *testing.T) {
		fromPath = "input.txt"
		limit = 1000
		offset = 100

		err := Copy(fromPath, toPath, offset, limit)

		require.Truef(t, errors.Is(err, ErrSrcFileNotFound), "actual err - %v", err)
	})
	t.Run("offset0_limit0", func(t *testing.T) {
		fromPath = "testdata/input.txt"
		limit = 0
		offset = 0

		Copy(fromPath, toPath, offset, limit)
		f1, err1 := ioutil.ReadFile(toPath)

		if err1 != nil {
			log.Fatal(err1)
		}

		f2, err2 := ioutil.ReadFile("testdata/out_offset0_limit0.txt")

		if err2 != nil {
			log.Fatal(err2)
		}
		defer os.Remove(toPath)
		require.Equal(t, f1, f2)
	})
	t.Run("offset0_limit10", func(t *testing.T) {
		limit = 10
		offset = 0

		Copy(fromPath, toPath, offset, limit)
		f1, err1 := ioutil.ReadFile(toPath)

		if err1 != nil {
			log.Fatal(err1)
		}

		f2, err2 := ioutil.ReadFile("testdata/out_offset0_limit10.txt")

		if err2 != nil {
			log.Fatal(err2)
		}
		defer os.Remove(toPath)
		require.Equal(t, f1, f2)
	})
	t.Run("offset0_limit1000", func(t *testing.T) {
		limit = 1000
		offset = 0

		Copy(fromPath, toPath, offset, limit)
		f1, err1 := ioutil.ReadFile(toPath)

		if err1 != nil {
			log.Fatal(err1)
		}

		f2, err2 := ioutil.ReadFile("testdata/out_offset0_limit1000.txt")

		if err2 != nil {
			log.Fatal(err2)
		}
		defer os.Remove(toPath)
		require.Equal(t, f1, f2)
	})
	t.Run("offset0_limit10000", func(t *testing.T) {
		limit = 10000
		offset = 0

		Copy(fromPath, toPath, offset, limit)
		f1, err1 := ioutil.ReadFile(toPath)

		if err1 != nil {
			log.Fatal(err1)
		}

		f2, err2 := ioutil.ReadFile("testdata/out_offset0_limit10000.txt")

		if err2 != nil {
			log.Fatal(err2)
		}
		defer os.Remove(toPath)
		require.Equal(t, f1, f2)
	})
	t.Run("offset100_limit1000", func(t *testing.T) {
		limit = 1000
		offset = 100

		Copy(fromPath, toPath, offset, limit)
		f1, err1 := ioutil.ReadFile(toPath)

		if err1 != nil {
			log.Fatal(err1)
		}

		f2, err2 := ioutil.ReadFile("testdata/out_offset100_limit1000.txt")

		if err2 != nil {
			log.Fatal(err2)
		}
		defer os.Remove(toPath)
		require.Equal(t, f1, f2)
	})
	t.Run("offset6000_limit1000", func(t *testing.T) {
		limit = 1000
		offset = 6000

		Copy(fromPath, toPath, offset, limit)
		f1, err1 := ioutil.ReadFile(toPath)

		if err1 != nil {
			log.Fatal(err1)
		}

		f2, err2 := ioutil.ReadFile("testdata/out_offset6000_limit1000.txt")

		if err2 != nil {
			log.Fatal(err2)
		}
		defer os.Remove(toPath)
		require.Equal(t, f1, f2)
	})
}
