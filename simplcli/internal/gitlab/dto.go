package gitlab

type SearchPipeline struct {
	Page    string `search:"page"`
	PerPage string `search:"per_page"`
	Ref     string `search:"ref"`
	Status  string `search:"status"`
}

type PipelinesResponseDto struct {
	Items []PipelineResponseItemDto
}

type PipelineResponseItemDto struct {
	Id        int    `json:"id"`
	Iid       int    `json:"iid"`
	ProjectId int    `json:"project_id"`
	Sha       string `json:"sha"`
	Ref       string `json:"ref"`
	Status    string `json:"status"`
	Source    string `json:"source"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	WebUrl    string `json:"web_url"`
	Name      string `json:"name"`
}

type PipelineResponseDto struct {
	Id             int    `json:"id"`
	Iid            int    `json:"iid"`
	ProjectId      int    `json:"project_id"`
	Name           string `json:"name"`
	Sha            string `json:"sha"`
	Ref            string `json:"ref"`
	Status         string `json:"status"`
	Source         string `json:"source"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	WebUrl         string `json:"web_url"`
	BeforeSha      string `json:"before_sha"`
	Tag            bool   `json:"tag"`
	StartedAt      string `json:"started_at"`
	FinishedAt     string `json:"finished_at"`
	CommittedAt    string `json:"committed_at"`
	Duration       int    `json:"duration"`
	QueuedDuration int    `json:"queued_duration"`
	Archived       bool   `json:"archived"`
}
