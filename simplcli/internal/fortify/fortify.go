package fortify

type TokenFunc = func() (string, error)
type Client struct {
	BaseUrl   string
	TokenFunc TokenFunc
}
