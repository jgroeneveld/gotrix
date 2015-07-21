package views

import "io"

type Renderer interface {
	Render(io.Writer) error
}

func RenderWithLayout(w io.Writer, r Renderer) error {
	err := Header(w)
	if err != nil {
		return err
	}
	err = r.Render(w)
	if err != nil {
		return err
	}
	err = Footer(w)
	if err != nil {
		return err
	}
	return nil
}
