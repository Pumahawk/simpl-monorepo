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
	Name      string `json:"name"`
	ProjectId int    `json:"project_id"`
	Iid       int    `json:"iid"`
	Sha       string `json:"sha"`
	Ref       string `json:"ref"`
	Status    string `json:"status"`
	Source    string `json:"source"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	WebUrl    string `json:"web_url"`
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

type SearchPipelineJob struct {
	Page    string `search:"page"`
	PerPage string `search:"per_page"`
	Scope   string `search:"scope"`
}

type PipelineJobsResponseDto struct {
	Items []PipelineJobsResponseItemDto
}

type PipelineJobsResponseItemDto struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Status            string  `json:"status"`
	Ref               string  `json:"ref"`
	Archived          bool    `json:"archived"`
	Source            string  `json:"source"`
	AllowFailure      bool    `json:"allow_failure"`
	CreatedAt         string  `json:"created_at"`
	StartedAt         string  `json:"started_at"`
	FinishedAt        string  `json:"finished_at"`
	ErasedAt          string  `json:"erased_at"`
	Duration          float32 `json:"duration"`
	QueuedDuration    float32 `json:"queued_duration"`
	ArtifactsExpireAt string  `json:"artifacts_expire_at"`
	Stage             string  `json:"stage"`
	FailureReason     string  `json:"failure_reason"`
	Tag               bool    `json:"tag"`
	WebUrl            string  `json:"web_url"`
}
