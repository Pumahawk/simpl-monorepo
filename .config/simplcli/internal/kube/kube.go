package kube

import (
	"bytes"
	"encoding/json"
	"fmt"
	"simplcli/internal/ex"
)

func GetByName(name string) (*GetResponse, error) {
	out := &bytes.Buffer{}
	cmd := ex.New("kubectl", "get", "pod", "-ojson", "--sort-by={.metadata.creationTimestamp}", "-l", fmt.Sprintf("app.kubernetes.io/name=%s", name))
	cmd.Stdout = out
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("get by %q", name)
	}
	r := &GetResponse{}
	if err := json.Unmarshal(out.Bytes(), r); err != nil {
		return nil, fmt.Errorf("response unmarshal %s: %w", out.Bytes(), err)
	}
	return r, nil
}

type GetResponse struct {
	Items []Item
}

type Item struct {
	Metadata Metadata
	Status   Status
}

type Metadata struct {
	Name string
}
type Status struct {
	ContainerStatuses []ContainerStatuses
}

type ContainerStatuses struct {
	Name  string
	Ready bool
}
