package fortify

type MarkFalsePositive struct {
	DeveloperStatus           string `json:"developerStatus"`
	AuditorStatus             string `json:"auditorStatus"`
	IncludeAllVulnerabilities bool   `json:"includeAllVulnerabilities"`
}
