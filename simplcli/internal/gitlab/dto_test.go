package gitlab_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/gitlab"
)

//go:embed dto-mock/MergeRequestResponseDto.json
var mergeRequestResponseDtoJson []byte

//go:embed dto-mock/JobRetryResponseDto.json
var jobRetryResponseDto []byte

func TestMergeRequestResponseDto(t *testing.T) {
	res := &gitlab.MergeRequestResponseDto{}
	if err := json.Unmarshal([]byte(mergeRequestResponseDtoJson), res); err != nil {
		t.Errorf("unable to unmashal gitlab.MergeRequestResponseDto: %s", err)
	}
}

func TestJobRetryResponseDto(t *testing.T) {
	res := &gitlab.JobRetryResponseDto{}
	if err := json.Unmarshal([]byte(jobRetryResponseDto), res); err != nil {
		t.Errorf("unable to unmashal gitlab.JobRetryResponseDto: %s", err)
	}
}
