package domain

import "time"

type Statistics struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64       // nil, если задач нет
	TasksAverageCompletionTime *time.Duration // nil, если нет выполненных
}

func CreateStatistics(tasks []Task) Statistics {
	tasksCreated := len(tasks)

	tasksCompleted := 0
	var totalDuration time.Duration
	for _, t := range tasks {
		if !t.Completed || t.CompletedAt == nil {
			continue
		}
		tasksCompleted++
		totalDuration += t.CompletedAt.Sub(t.CreatedAt)
	}

	var completedRate *float64
	if tasksCreated > 0 {
		rate := float64(tasksCompleted) / float64(tasksCreated) * 100
		completedRate = &rate
	}

	var avgTime *time.Duration
	if tasksCompleted > 0 && totalDuration != 0 {
		avg := totalDuration / time.Duration(tasksCompleted)
		avgTime = &avg
	}

	return Statistics{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         completedRate,
		TasksAverageCompletionTime: avgTime,
	}
}
