package gitlab

import (
	"net/url"
)

type Client struct {
	BaseUrl string
}

func (c *Client) Project(id string, search *SearchPipeline) (*PipelineResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "projects", url.PathEscape(id), "pipelines")
	if err != nil {
		panic(err)
	}

	items := make([]PipelineResponseItemDto, 0, 10)
	err = searchApi(c, rawUrl, search, &items)
	if err != nil {
		return nil, err
	}
	return &PipelineResponseDto{
		Items: items,
	}, nil
}
