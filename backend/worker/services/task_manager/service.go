package taskmanager

import "github.com/julianfbeck/universal/backend/services/stock/shared/repository/redis"

type TaskManager struct {
	redisRepository *redis.RedisRepository
}

func NewTaskManager(redisRepository *redis.RedisRepository) *TaskManager {
	return &TaskManager{
		redisRepository: redisRepository,
	}
}

func (tm *TaskManager) StartTaskQueue(queueName string, numberOfWorkers int) (string, error) {
	for i := 0; i < numberOfWorkers; i++ {

	}
	return "", nil
}

func (tm *TaskManager) StartRealtimeQueue(queueName string) (string, error) {
	return "", nil
}

func (tm *TaskManager) StartNotificationQueue(queueName string) (string, error) {
	return "", nil
}
