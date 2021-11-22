package store

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

type Store struct {
	path string
	mx   sync.RWMutex
}

func (s *Store) Save(data interface{}) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	byteData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return ErrSave
	}
	if err := os.WriteFile(s.path, byteData, 0600); err != nil {
		return ErrSave
	}
	return nil
}

func (s *Store) Load(dest interface{}) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	jsonFile, err := os.Open(s.path)
	if err != nil {
		return err
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			log.Printf("error closing file [%s]: %s ", s.path, err)
		}
	}()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &dest); err != nil {
		return err
	}
	return nil
}

func (s *Store) Exists() bool {
	info, err := os.Stat(s.path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
