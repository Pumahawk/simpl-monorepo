package fortify_test

import (
	_ "embed"
	"encoding/json"
	"testing"

	"github.com/Pumahawk/simpl-monorepo/simplcli/internal/fortify"
)

//go:embed dto-mock/MarkFalsePositive.json
var markFalsePositive []byte

func TestMarkFalsePositive(t *testing.T) {
	res := &fortify.MarkFalsePositive{}
	if err := json.Unmarshal([]byte(markFalsePositive), res); err != nil {
		t.Errorf("unable to unmashal fortify.MarkFalsePositive: %s", err)
	}
}
