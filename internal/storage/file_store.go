package storage

import (
	"bufio"
	"os"
)

type FileStore struct {
	path string
}

func NewFileStore(path string) *FileStore {
	return &FileStore{path: path}
}

func (f *FileStore) Add(email string) error {
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(email + "\n")
	return err
}

func (f *FileStore) Exists(email string) (bool, error) {
	subscribers, err := f.List()
	if err != nil {
		return false, err
	}

	for _, s := range subscribers {
		if s == email {
			return true, nil
		}
	}
	return false, nil
}

func (f *FileStore) List() ([]string, error) {
	file, err := os.Open(f.path)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var subscribers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		subscribers = append(subscribers, scanner.Text())
	}

	return subscribers, scanner.Err()
}
