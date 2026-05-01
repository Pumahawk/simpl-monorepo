package gitlab

import (
	"net/url"
)

type Client struct {
	BaseUrl string
}

func (c *Client) Pipelines(id string, search *SearchPipeline) (*PipelinesResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "projects", url.PathEscape(id), "pipelines")
	if err != nil {
		panic(err)
	}

	items := make([]PipelineResponseItemDto, 0, 10)
	err = searchApi(c, rawUrl, search, &items)
	if err != nil {
		return nil, err
	}
	return &PipelinesResponseDto{
		Items: items,
	}, nil
}

func (c *Client) Pipeline(projectId, pipelineId string) (*PipelineResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "projects", url.PathEscape(projectId), "pipelines", url.PathEscape(pipelineId))
	if err != nil {
		panic(err)
	}

	res := &PipelineResponseDto{}
	_, err = request("GET", rawUrl, nil, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) PipelineJobs(projectId, pipelineId string, search *SearchPipelineJob) (*PipelineJobsResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "projects", url.PathEscape(projectId), "pipelines", url.PathEscape(pipelineId), "jobs")
	if err != nil {
		panic(err)
	}

	items := make([]PipelineJobsResponseItemDto, 0, 10)
	err = searchApi(c, rawUrl, search, &items)
	if err != nil {
		return nil, err
	}
	return &PipelineJobsResponseDto{
		Items: items,
	}, nil
}
