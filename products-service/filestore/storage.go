package filestore

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

const (
	UNABLE_TO_CREATE_DIRECTORY_ERR_MSG = "unable to create directory with error %v"
	UNABLE_TO_GET_FILE_INFO            = "unable to get file info with error %v"
	UNABLE_TO_REMOVE_FILE              = "unable to remove file with error %v"
	UNABLE_TO_CREATE_FILE              = "unable to create file with error %v"
	UNABLE_TO_WRITE_FILE               = "unable to write file with error %v"
	UNABLE_TO_READ_FILE                = "unable to read file with error %v"
)

type Storage interface {
	Save(path string, file io.Reader) error
}

type LocalFileStorage struct {
	maxFileSize int64  // Max. file size in bytes
	basePath    string // Location for file storage
	log         *log.Logger
}

func NewLocalFileStorage(basePath string, fileSize int64, logr *log.Logger) (*LocalFileStorage, error) {
	path, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}
	return &LocalFileStorage{
		basePath:    path,     // Set absolute path as base
		maxFileSize: fileSize, // Set max file size
		log:         logr,
	}, nil
}

// Saves given file contents on a given path in the base path directory.
func (local *LocalFileStorage) Save(path string, contents io.Reader) error {
	// Create local path and make directory on local directory path
	localPath := local.fullPath(path)
	dir := filepath.Dir(localPath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return xerrors.Errorf(UNABLE_TO_CREATE_DIRECTORY_ERR_MSG, err)
	}

	// If file already exists try to delete that
	_, err = os.Stat(localPath)
	if err == nil {
		err = os.Remove(localPath)
		if err != nil {
			return xerrors.Errorf(UNABLE_TO_REMOVE_FILE, err)
		}
	} else if !os.IsNotExist(err) {
		// File path does not exists
		return xerrors.Errorf(UNABLE_TO_GET_FILE_INFO, err)
	}

	// Create file at the local path
	file, err := os.Create(localPath)
	if err != nil {
		return xerrors.Errorf(UNABLE_TO_CREATE_FILE, err)
	}

	defer file.Close()

	// Write contents to the new file
	_, err = io.Copy(file, contents)
	if err != nil {
		return xerrors.Errorf(UNABLE_TO_WRITE_FILE, err)
	}

	return nil
}

// Returns absolute path i.e. base path + file path.
func (local *LocalFileStorage) fullPath(path string) string {
	return filepath.Join(local.basePath, path)
}
