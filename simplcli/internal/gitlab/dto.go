package gitlab

import "fmt"

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

type PipelineAttributesResponseDto struct {
	Items []PipelineAttributesItemResponseDto
}

type PipelineAttributesItemResponseDto struct {
	Key          string `json:"key"`
	VariableType string `json:"variable_type"`
	Value        string `json:"value"`
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

type SearchMergeRequest struct {
	Page    string `search:"page"`
	PerPage string `search:"per_page"`

	Id                     string `search:"id"`
	AssigneeId             int    `search:"assignee_id"`
	AuthorId               int    `search:"author_id"`
	AuthorUsername         string `search:"author_username"`
	CreatedAfter           string `search:"created_after"`
	CreatedBefore          string `search:"created_before"`
	DeployedAfter          string `search:"deployed_after"`
	DeployedBefore         string `search:"deployed_before"`
	Environment            string `search:"environment"`
	In                     string `search:"in"`
	Labels                 string `search:"labels"`
	MergeUserId            int    `search:"merge_user_id"`
	MergeUserUsername      string `search:"merge_user_username"`
	Milestone              string `search:"milestone"`
	MyReactionEmoji        string `search:"my_reaction_emoji"`
	OrderBy                string `search:"order_by"`
	ReviewerId             int    `search:"reviewer_id"`
	ReviewerUsername       string `search:"reviewer_username"`
	Scope                  string `search:"scope"`
	Search                 string `search:"search"`
	Sort                   string `search:"sort"`
	SourceBranch           string `search:"source_branch"`
	State                  string `search:"state"`
	TargetBranch           string `search:"target_branch"`
	UpdatedAfter           string `search:"updated_after"`
	UpdatedBefore          string `search:"updated_before"`
	View                   string `search:"view"`
	Wip                    string `search:"wip"`
	WithLabelsDetails      bool   `search:"with_labels_details"`
	WithMergeStatusRecheck bool   `search:"with_merge_status_recheck"`
}

type MergeRequestsResponseDto struct {
	Items []MergeRequestsResponseItemDto
}

type MergeRequestsResponseItemDto struct {
	Id                          int                     `json:"id"`
	Iid                         int                     `json:"iid"`
	Assignee                    UserInfoDto             `json:"assignee"`
	Assignees                   []UserInfoDto           `json:"assignees"`
	Author                      UserInfoDto             `json:"author"`
	BlockingDiscussionsResolved bool                    `json:"blocking_discussions_resolved"`
	ClosedAt                    string                  `json:"closed_at"`
	ClosedBy                    UserInfoDto             `json:"closed_by"`
	CreatedAt                   string                  `json:"created_at"`
	Description                 string                  `json:"description"`
	DetailedMergeStatus         string                  `json:"detailed_merge_status"`
	DiscussionLocked            bool                    `json:"discussion_locked"`
	Downvotes                   int                     `json:"downvotes"`
	Draft                       bool                    `json:"draft"`
	ForceRemoveSourceBranch     bool                    `json:"force_remove_source_branch"`
	HasConflicts                bool                    `json:"has_conflicts"`
	Labels                      []string                `json:"labels"`
	MergeCommitSha              string                  `json:"merge_commit_sha"`
	MergeStatus                 string                  `json:"merge_status"`
	MergeUser                   UserInfoDto             `json:"merge_user"`
	MergeWhenPipelineSucceeds   bool                    `json:"merge_when_pipeline_succeeds"`
	MergedAt                    string                  `json:"merged_at"`
	MergedBy                    UserInfoDto             `json:"merged_by"`
	Milestone                   MilestoneInfoDto        `json:"milestone"`
	PreparedAt                  string                  `json:"prepared_at"`
	ProjectId                   int                     `json:"project_id"`
	Reference                   string                  `json:"reference"`
	References                  ReferencesDto           `json:"references"`
	Reviewers                   []UserInfoDto           `json:"reviewers"`
	Sha                         string                  `json:"sha"`
	ShouldRemoveSourceBranch    bool                    `json:"should_remove_source_branch"`
	SourceBranch                string                  `json:"source_branch"`
	SourceProjectId             int                     `json:"source_project_id"`
	Squash                      bool                    `json:"squash"`
	SquashCommitSha             string                  `json:"squash_commit_sha"`
	SquashOnMerge               bool                    `json:"squash_on_merge"`
	State                       string                  `json:"state"`
	TargetBranch                string                  `json:"target_branch"`
	TargetProjectId             int                     `json:"target_project_id"`
	TaskCompletionStatus        TaskCompletionStatusDto `json:"task_completion_status"`
	TimeStats                   TimeStatsDto            `json:"time_stats"`
	Title                       string                  `json:"title"`
	UpdatedAt                   string                  `json:"updated_at"`
	Upvotes                     int                     `json:"upvotes"`
	UserNotesCount              int                     `json:"user_notes_count"`
	WebUrl                      string                  `json:"web_url"`
	WorkInProgress              bool                    `json:"work_in_progress"`
}

type UserInfoDto struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Id        int    `json:"id"`
	State     string `json:"state"`
	Locked    bool   `json:"locked"`
	AvatarUrl string `json:"avatar_url"`
	WebUrl    string `json:"web_url"`
}

func (u UserInfoDto) String() string {
	return fmt.Sprintf("%s [%s]", u.Name, u.Username)
}

type MilestoneInfoDto struct {
	Id          int    `json:"id"`
	Iid         int    `json:"iid"`
	ProjectId   int    `json:"project_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	State       string `json:"state"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DueDate     string `json:"due_date"`
	StartDate   string `json:"start_date"`
	WebUrl      string `json:"web_url"`
}

func (m MilestoneInfoDto) String() string {
	return fmt.Sprintf("%d", m.Id)
}

type TaskCompletionStatusDto struct {
	Count          int `json:"count"`
	CompletedCount int `json:"completed_count"`
}

func (t TaskCompletionStatusDto) String() string {
	return fmt.Sprintf("(%d/%d)", t.CompletedCount, t.Count)
}

type TimeStatsDto struct {
	TimeEstimate        int    `json:"time_estimate"`
	TotalTimeSpent      int    `json:"total_time_spent"`
	HumanTimeEstimate   string `json:"human_time_estimate"`
	HumanTotalTimeSpent string `json:"human_total_time_spent"`
}

func (t TimeStatsDto) String() string {
	return fmt.Sprintf("%s/%s", t.HumanTimeEstimate, t.HumanTotalTimeSpent)
}

type ReferencesDto struct {
	Short    string `json:"short"`
	Relative string `json:"relative"`
	Full     string `json:"full"`
}

func (r ReferencesDto) String() string {
	return fmt.Sprintf("%s", r.Short)
}

type MergeRequestResponseDto struct {
	AllowCollaboration          bool                    `json:"allow_collaboration"`
	AllowMaintainerToPush       bool                    `json:"allow_maintainer_to_push"`
	Assignee                    UserInfoDto             `json:"assignee"`
	Assignees                   []UserInfoDto           `json:"assignees"`
	Author                      UserInfoDto             `json:"author"`
	BlockingDiscussionsResolved bool                    `json:"blocking_discussions_resolved"`
	ChangesCount                string                  `json:"changes_count"`
	ClosedAt                    string                  `json:"closed_at"`
	ClosedBy                    UserInfoDto             `json:"closed_by"`
	CreatedAt                   string                  `json:"created_at"`
	Description                 string                  `json:"description"`
	DetailedMergeStatus         string                  `json:"detailed_merge_status"`
	DiffRefs                    DiffRefDto              `json:"diff_refs"`
	DiscussionLocked            bool                    `json:"discussion_locked"`
	DivergedCommitsCount        int                     `json:"diverged_commits_count"`
	Downvotes                   int                     `json:"downvotes"`
	Draft                       bool                    `json:"draft"`
	FirstContribution           bool                    `json:"first_contribution"`
	FirstDeployedToProductionAt string                  `json:"first_deployed_to_production_at"`
	ForceRemoveSourceBranch     bool                    `json:"force_remove_source_branch"`
	HasConflicts                bool                    `json:"has_conflicts"`
	Id                          int                     `json:"id"`
	Iid                         int                     `json:"iid"`
	Imported                    bool                    `json:"imported"`
	ImportedFrom                string                  `json:"imported_from"`
	Labels                      []LabelsDto             `json:"labels"`
	LatestBuildFinishedAt       string                  `json:"latest_build_finished_at"`
	LatestBuildStartedAt        string                  `json:"latest_build_started_at"`
	MergeAfter                  string                  `json:"merge_after"`
	MergeCommitSha              string                  `json:"merge_commit_sha"`
	MergeError                  string                  `json:"merge_error"`
	MergeStatus                 string                  `json:"merge_status"`
	MergeUser                   UserInfoDto             `json:"merge_user"`
	MergeWhenPipelineSucceeds   bool                    `json:"merge_when_pipeline_succeeds"`
	MergedAt                    string                  `json:"merged_at"`
	MergedBy                    UserInfoDto             `json:"merged_by"`
	Milestone                   []MilestoneInfoDto      `json:"milestone"`
	Pipeline                    PipelineInfoDto         `json:"pipeline"`
	PreparedAt                  string                  `json:"prepared_at"`
	ProjectId                   int                     `json:"project_id"`
	RebaseInProgress            bool                    `json:"rebase_in_progress"`
	HeadPipeline                PipelineInfoDto         `json:"head_pipeline"`
	Reference                   string                  `json:"reference"`
	References                  ReferencesDto           `json:"references"`
	Reviewers                   []UserInfoDto           `json:"reviewers"`
	Sha                         string                  `json:"sha"`
	ShouldRemoveSourceBranch    bool                    `json:"should_remove_source_branch"`
	SourceBranch                string                  `json:"source_branch"`
	SourceProjectId             int                     `json:"source_project_id"`
	Squash                      bool                    `json:"squash"`
	SquashCommitSha             string                  `json:"squash_commit_sha"`
	SquashOnMerge               bool                    `json:"squash_on_merge"`
	State                       string                  `json:"state"`
	Subscribed                  bool                    `json:"subscribed"`
	TargetBranch                string                  `json:"target_branch"`
	TargetProjectId             int                     `json:"target_project_id"`
	TaskCompletionStatus        TaskCompletionStatusDto `json:"task_completion_status"`
	TimeStats                   TimeStatsDto            `json:"time_stats"`
	Title                       string                  `json:"title"`
	UpdatedAt                   string                  `json:"updated_at"`
	Upvotes                     int                     `json:"upvotes"`
	User                        UserInfoDto             `json:"user"`
	UserNotesCount              int                     `json:"user_notes_count"`
	WebUrl                      string                  `json:"web_url"`
	WorkInProgress              bool                    `json:"work_in_progress"`
}

type PipelineInfoDto struct {
	YamlErrors     string                        `json:"yaml_errors"`
	WebUrl         string                        `json:"web_url"`
	User           UserInfoDto                   `json:"user"`
	UpdatedAt      string                        `json:"updated_at"`
	Tag            bool                          `json:"tag"`
	Status         string                        `json:"status"`
	StartedAt      string                        `json:"started_at"`
	Source         string                        `json:"source"`
	Sha            string                        `json:"sha"`
	Ref            string                        `json:"ref"`
	QueuedDuration int                           `json:"queued_duration"`
	ProjectId      int                           `json:"project_id"`
	Iid            int                           `json:"iid"`
	Id             int                           `json:"id"`
	FinishedAt     string                        `json:"finished_at"`
	Duration       int                           `json:"duration"`
	DetailedStatus MergeRequestDetailedStatusDto `json:"detailed_status"`
	CreatedAt      string                        `json:"created_at"`
	Coverage       string                        `json:"coverage"`
	CommittedAt    string                        `json:"committed_at"`
	BeforeSha      string                        `json:"before_sha"`
}

type DiffRefDto struct {
	BaseSha  string `json:"base_sha"`
	StartSha string `json:"start_sha"`
	HeadSha  string `json:"head_sha"`
}

type MergeRequestDetailedStatusDto struct {
	Tooltip     string      `json:"tooltip"`
	Text        string      `json:"text"`
	Label       string      `json:"label"`
	Icon        string      `json:"icon"`
	HasDetails  bool        `json:"has_details"`
	Group       string      `json:"group"`
	Favicon     string      `json:"favicon"`
	DetailsPath string      `json:"details_path"`
	Action      []ActionDto `json:"action"`
}

type ActionDto struct {
	Title               string `json:"title"`
	Path                string `json:"path"`
	Method              string `json:"method"`
	Icon                string `json:"icon"`
	ConfirmationMessage string `json:"confirmation_message"`
	ButtonTitle         string `json:"button_title"`
}

type LabelsDto struct {
	Archived        bool   `json:"archived"`
	Color           string `json:"color"`
	Description     string `json:"description"`
	DescriptionHtml string `json:"description_html"`
	Id              int    `json:"id"`
	Name            string `json:"name"`
	TextColor       string `json:"text_color"`
}

type JobRetryResponseDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
