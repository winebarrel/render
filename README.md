# render

[![test](https://github.com/winebarrel/render/actions/workflows/test.yml/badge.svg)](https://github.com/winebarrel/render/actions/workflows/test.yml)

Rendering files using [Sigil](https://github.com/gliderlabs/sigil) templating.

It works similarly to [Entrykit](https://github.com/progrium/entrykit)'s render command.

## Usage

```
$ render -h
Usage: render [OPTION] FILE...
  -version
    	print version and exit
```

```
$ echo 'London Bridge is {{ var "WHAT_HAPPEN" | default "broken down" }}' > test.conf.tmpl
$ render test.conf
$ cat test.conf
London Bridge is broken down

$ export WHAT_HAPPEN='falling down'
$ render test.conf
$ cat test.conf
London Bridge is falling down
```

## Related Links

* [gliderlabs/sigil: Standalone string interpolator and template processor](https://github.com/gliderlabs/sigil)
* [template package - text/template - Go Packages](https://pkg.go.dev/text/template)
* https://github.com/progrium/entrykit#render---template-rendering

## Installation

See https://github.com/winebarrel/render/releases/latest

### Homebrew

```
brew install winebarrel/render/render
```

