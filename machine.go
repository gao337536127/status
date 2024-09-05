package status

import (
	"fmt"
	"sync"
)

type ProgramStatus struct {
	StatusPool map[string]func() (string, error)
	lock       sync.Mutex
}

func NewProgramStatus() *ProgramStatus {
	return &ProgramStatus{
		StatusPool: make(map[string]func() (string, error)),
	}
}

func (m *ProgramStatus) SetStatus(name string, statusFunc func() (string, error)) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.StatusPool[name] = statusFunc
}

func (m *ProgramStatus) GetStatus(name string) (string, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if statusMethod, ok := m.StatusPool[name]; ok {
		return statusMethod()
	} else {
		return "", fmt.Errorf("unregister status named %s", name)
	}
}
