package vw

import (
	"encoding/json"
	"io"
)

type JsonView struct {
	w io.Writer
}

func NewJsonView(w io.Writer) *JsonView {
	return &JsonView{w}
}

func (j *JsonView) Render(opt *RenderOpt, model any) error {
	w := json.NewEncoder(j.w)
	w.SetIndent(" ", " ")
	return w.Encode(model)
}
