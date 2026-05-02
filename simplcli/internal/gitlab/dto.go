package gitlab

type SearchPipeline struct {
	Page    string `search:"page"`
	PerPage string `search:"per_page"`

	Id            string `search:"id"`
	Name          string `search:"name"`
	OrderBy       string `search:"order_by"`
	Ref           string `search:"ref"`
	Scope         string `search:"scope"`
	Sha           string `search:"sha"`
	Sort          string `search:"sort"`
	Source        string `search:"source"`
	Status        string `search:"status"`
	UpdatedAfter  string `search:"updated_after"`
	UpdatedBefore string `search:"updated_before"`
	CreatedAfter  string `search:"created_after"`
	CreatedBefore string `search:"created_before"`
	Username      string `search:"username"`
	YamlErrors    bool   `search:"yaml_errors"`
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

	PipelineId string `search:"pipeline_id"`
	Scope      string `search:"scope"`
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

type SearchRegistry struct {
	Page    string `search:"page"`
	PerPage string `search:"per_page"`

	Id                 string `search:"id"`
	OrderBy            string `search:"order_by"`
	Sort               string `search:"sort,def:desc"`
	PackageType        string `search:"package_type"`
	PackageName        string `search:"package_name"`
	PackageVersion     string `search:"package_version"`
	IncludeVersionless bool   `search:"include_versionless"`
	Status             string `search:"status"`
}

type RegistryResponseDto struct {
	Items []RegistryResponseItemDto
}

type RegistryResponseItemDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	PackageType string `json:"package_type"`
	CreatedAt   string `json:"created_at"`
	CreatorId   int    `json:"creator_id"`
}
