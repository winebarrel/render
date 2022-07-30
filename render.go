package render

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gliderlabs/sigil"
	_ "github.com/gliderlabs/sigil/builtin"
)

const (
	TemplateSuffix = ".tmpl"
)

func Render(target string) error {
	tmpl := target + TemplateSuffix
	input, err := ioutil.ReadFile(tmpl)

	if err != nil {
		return err
	}

	buf, err := sigil.Execute(input, nil, filepath.Base(target))

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(target, buf.Bytes(), 0644)

	return err
}
