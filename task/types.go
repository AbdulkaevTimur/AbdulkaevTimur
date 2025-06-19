package task

import "time"

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Task struct {
	ID         string    `json:"id"`
	Status     Status    `json:"status"`
	Result     string    `json:"result"`
	CreatedAt  time.Time `json:"created_at"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
}

func (t *Task) Duration() time.Duration {
	if !t.FinishedAt.IsZero() {
		return t.FinishedAt.Sub(t.CreatedAt)
	}
	return time.Since(t.CreatedAt)
}
