package render

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	assert := assert.New(t)

	dir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(dir)

	tmplPath := filepath.Join(dir, "test.conf.tmpl")
	confPath := strings.TrimSuffix(tmplPath, ".tmpl")
	tmpl := `London Bridge is {{ var "WHAT_HAPPEN" | default "broken down" }}`
	ioutil.WriteFile(tmplPath, []byte(tmpl), 0644)

	err := Render(confPath)

	assert.NoError(err)
	conf, err := ioutil.ReadFile(confPath)
	assert.NoError(err)
	assert.Equal("London Bridge is broken down", string(conf))
}

func TestRenderWithEnv(t *testing.T) {
	assert := assert.New(t)
	t.Setenv("WHAT_HAPPEN", "falling down")

	dir, _ := ioutil.TempDir("", "")
	defer os.RemoveAll(dir)

	tmplPath := filepath.Join(dir, "test.conf.tmpl")
	confPath := strings.TrimSuffix(tmplPath, ".tmpl")
	tmpl := `London Bridge is {{ var "WHAT_HAPPEN" | default "broken down" }}`
	ioutil.WriteFile(tmplPath, []byte(tmpl), 0644)

	err := Render(confPath)

	assert.NoError(err)
	conf, err := ioutil.ReadFile(confPath)
	assert.NoError(err)
	assert.Equal("London Bridge is falling down", string(conf))
}

func TestRenderWithoutTmpl(t *testing.T) {
	assert := assert.New(t)
	err := Render("/no/such/file")
	assert.EqualError(err, "open /no/such/file.tmpl: no such file or directory")
}
