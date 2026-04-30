package gitlab

type SearchPipeline struct {
	Page    string `search:"page"`
	PerPage string `search:"per_page"`
	Ref     string `search:"ref"`
	Status  string `search:"status"`
}

type PipelineResponseDto struct {
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
