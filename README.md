# render

Rendering files using [Sigil](https://github.com/gliderlabs/sigil) templating.

It works similarly to [Entrykit](https://github.com/progrium/entrykit)'s render command.

## Usage

```
$ echo 'London Bridge is {{ var "WHAT_HAPPEN" | default "broken down" }}' > test.conf.tmpl
$ render test.conf
$ cat test.conf
London Bridge is broken down
```

## Related Links

* [gliderlabs/sigil: Standalone string interpolator and template processor](https://github.com/gliderlabs/sigil)
* [template package - text/template - Go Packages](https://pkg.go.dev/text/template)
* https://github.com/progrium/entrykit#render---template-rendering

