package gitlab_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/Pumahawk/simpl-monorepo/internal/gitlab"
)

//go:embed dto-mock/MergeRequestResponseDto.json
var mergeRequestResponseDtoJson []byte

func TestMergeRequestResponseDto(t *testing.T) {
	res := &gitlab.MergeRequestResponseDto{}
	if err := json.Unmarshal([]byte(mergeRequestResponseDtoJson), res); err != nil {
		t.Errorf("unable to unmashal gitlab.MergeRequestResponseDto: %s", err)
	}
}
