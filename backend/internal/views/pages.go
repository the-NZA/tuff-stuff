package views

import (
	"io"
)

func Homepage(w io.Writer, p *HomepageViewData) error {
	return homepage.Execute(w, p)
}
