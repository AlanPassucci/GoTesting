package storage

import (
	"encoding/json"
	"os"

	"github.com/stretchr/testify/mock"
)

type Storage interface {
	GetValue(key string) interface{}
}

type storage struct {
	file string
}

func NewStorage() Storage {
	file := "config.json"
	return &storage{file: file}
}

func (s *storage) GetValue(key string) interface{} {
	file, err := os.ReadFile(s.file)
	if err != nil {
		return nil
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil
	}

	if v, ok := data[key]; ok {
		return v
	}

	return nil
}

type MockStorage struct {
	mock.Mock
}

func (m *MockStorage) GetValue(key string) interface{} {
	args := m.Called(key)
	return args.Get(0)
}
