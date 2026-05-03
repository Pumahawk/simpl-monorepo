package gitlab

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type TokenFunc = func() (string, error)
type Client struct {
	BaseUrl   string
	TokenFunc TokenFunc
}

func (c *Client) token() (string, error) {
	if c.TokenFunc != nil {
		return c.TokenFunc()
	} else {
		return "", nil
	}
}

func (c *Client) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return r, err
	}
	t, err := c.token()
	if err != nil {
		return nil, err
	} else if t != "" {
		r.Header.Add("PRIVATE-TOKEN", t)
	}
	return r, nil
}

func (c *Client) Pipelines(id string, search *SearchPipeline) (*PipelinesResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v4", "projects", url.PathEscape(id), "pipelines")
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
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v4", "projects", url.PathEscape(projectId), "pipelines", url.PathEscape(pipelineId))
	if err != nil {
		panic(err)
	}

	res := &PipelineResponseDto{}
	_, err = request(c, "GET", rawUrl, nil, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) PipelineJobs(projectId, pipelineId string, search *SearchPipelineJob) (*PipelineJobsResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v4", "projects", url.PathEscape(projectId), "pipelines", url.PathEscape(pipelineId), "jobs")
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

func (c *Client) JobLog(projectId, jobId string) ([]byte, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v4", "projects", url.PathEscape(projectId), "jobs", url.PathEscape(jobId), "trace")
	if err != nil {
		panic(err)
	}
	rq, err := c.NewRequest("GET", rawUrl, nil)
	if err != nil {
		panic(err)
	}
	cl := http.Client{}
	r, err := cl.Do(rq)
	if err != nil {
		return nil, fmt.Errorf("jobs projectId=%q jobsId=%q: make request: %w", projectId, jobId, err)
	}
	defer r.Body.Close()

	if r.StatusCode < 200 || r.StatusCode >= 300 {
		return nil, fmt.Errorf("jobs projectId=%q jobsId=%q: status code %d", projectId, jobId, r.StatusCode)
	}

	bf := &bytes.Buffer{}
	_, err = io.Copy(bf, r.Body)
	if err != nil {
		return nil, fmt.Errorf("jobs projectId=%q jobsId=%q: copy response: %w", projectId, jobId, err)
	}
	return bf.Bytes(), nil
}

func (c *Client) Registry(projectId string, search *SearchRegistry) (*RegistryResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v4", "projects", url.PathEscape(projectId), "packages")
	if err != nil {
		panic(err)
	}

	items := make([]RegistryResponseItemDto, 0, 10)
	err = searchApi(c, rawUrl, search, &items)
	if err != nil {
		return nil, err
	}

	return &RegistryResponseDto{
		Items: items,
	}, nil
}

func (c *Client) MergeRequests(projectId string, search *SearchMergeRequest) (*MergeRequestsResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v4", "projects", url.PathEscape(projectId), "merge_requests")
	if err != nil {
		panic(err)
	}

	items := make([]MergeRequestsResponseItemDto, 0, 10)
	err = searchApi(c, rawUrl, search, &items)
	if err != nil {
		return nil, err
	}

	return &MergeRequestsResponseDto{
		Items: items,
	}, nil
}

func (c *Client) MergeRequest(projectId, mergeRequestId string) (*MergeRequestResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v4", "projects", url.PathEscape(projectId), "merge_requests", url.PathEscape(mergeRequestId))
	if err != nil {
		panic(err)
	}

	res := &MergeRequestResponseDto{}
	_, err = request(c, "GET", rawUrl, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) JobRetry(projectId, jobId string) (*JobRetryResponseDto, error) {
	rawUrl, err := url.JoinPath(c.BaseUrl, "/api/v4", "projects", url.PathEscape(projectId), "jobs", url.PathEscape(jobId), "retry")
	if err != nil {
		panic(err)
	}

	res := &JobRetryResponseDto{}
	_, err = request(c, "POST", rawUrl, nil, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
