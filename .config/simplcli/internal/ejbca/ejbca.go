package ejbca

import (
	"fmt"
	"simplcli/internal/kube"
)

var NotFound = fmt.Errorf("Not found")

func Get() (*kube.Item, error) {
	r, err := kube.GetByName("ejbca")
	if err != nil {
		return nil, err
	}
	if len(r.Items) > 0 {
		return &r.Items[len(r.Items)-1], nil
	} else {
		return nil, NotFound
	}
}
