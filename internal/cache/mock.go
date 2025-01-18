package cache

import (
	"errors"
	"time"
)

type MockCache struct {
	data map[string]string
}

func NewMockCache() *MockCache {
	return &MockCache{data: make(map[string]string)}
}

func (m *MockCache) Get(key string) (string, error) {
	if val, ok := m.data[key]; ok {
		return val, nil
	}
	return "", errors.New("cache miss")
}

func (m *MockCache) Set(key string, value string, _ time.Duration) error {
	m.data[key] = value
	return nil
}
