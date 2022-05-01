package shareddomain

type ScraperTask struct {
	TaskId        string `json:"task_id"`
	Url           string `json:"url"`
	ScheduledTime string `json:"scheduled_time"`
	Status        string `json:"status"`
	DatabaseId    string `json:"database_id"`
}
