package vw

type View interface {
	Render(opt *RenderOpt, model any) error
}

type RenderOpt struct {
	Fields []string
}
