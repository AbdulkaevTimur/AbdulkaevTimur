package task

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type Manager struct {
	tasks map[string]*Task
	mu    sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		tasks: make(map[string]*Task),
	}
}

func (m *Manager) runTask(t *Task) {
	defer func() {
		if r := recover(); r != nil {
			m.mu.Lock()
			t.Status = StatusFailed
			t.Result = "panic occurred"
			m.mu.Unlock()
		}
	}()
	m.mu.Lock()
	t.Status = StatusRunning
	t.StartedAt = time.Now()
	m.mu.Unlock()

	time.Sleep(3 * time.Minute)

	m.mu.Lock()
	t.Status = StatusCompleted
	t.FinishedAt = time.Now()
	t.Result = "Success!"
	m.mu.Unlock()
}

func (m *Manager) CreateTask() string {
	id := uuid.NewString()
	task := &Task{
		ID:        id,
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
	m.mu.Lock()
	m.tasks[id] = task
	m.mu.Unlock()

	go m.runTask(task)

	return id
}

func (m *Manager) GetTask(id string) (*Task, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	task, ok := m.tasks[id]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (m *Manager) DeleteTask(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.tasks[id]; !ok {
		return errors.New("task not found")
	}
	delete(m.tasks, id)
	return nil
}
