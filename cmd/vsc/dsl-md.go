package main

import (
	"bytes"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/emicklei/melrose"
	"github.com/emicklei/melrose/dsl"
)

var tmpl = template.Must(template.New("dsl").Parse(
	`---
title: Melrose Language
---

[Home](index.html)
[Usage](cli.html)
[Language](dsl.html)
[DAW](daw.html)
[Install](install.html)

### variables

Variable names must start with a non-digit character and can zero or more characters in [a-z A-Z _ 0-9].
An assigment "=" is used to create a Variable.
To delete a variable, assign it to the special value "nil".

### comment

Use "//" to add comment.

## Creation functions
{{range .Core}}
### {{.Title}}<a name="{{.Anchor}}"></a>
{{.Description}}
{{ range .Examples }}
	{{ . }}
{{ end }}{{end}}

## Composition functions
{{range .Composer}}
### {{.Title}}<a name="{{.Anchor}}"></a>
{{.Description}}
{{ range .Examples }}
	{{ . }}
{{ end }}{{end}}

## Audio control functions
{{range .Audio}}
### {{.Title}}<a name="{{.Anchor}}"></a>
{{.Description}}
{{ range .Examples }}
	{{ . }}
{{ end }}{{end}}

##### generated by dsl-md.go
`))

type DocumentedFunction struct {
	Title            string
	ShortDescription string
	Description      string
	Examples         []string
	Anchor           string
}

type GroupedFunctions struct {
	Core     []DocumentedFunction
	Composer []DocumentedFunction
	Audio    []DocumentedFunction
}

func dslmarkdown() {
	varstore := dsl.NewVariableStore()
	gf := GroupedFunctions{}
	for k, each := range dsl.EvalFunctions(varstore, melrose.NoLooper) {
		df := DocumentedFunction{
			Title:            k,
			ShortDescription: each.Title,
			Description:      firstUpcaseAndDot(each.Description),
			Examples:         strings.Split(each.Samples, "\n"),
			Anchor:           k,
		}
		if each.ControlsAudio {
			gf.Audio = append(gf.Audio, df)
		}
		if each.IsCore {
			gf.Core = append(gf.Core, df)
		}
		if each.IsComposer {
			gf.Composer = append(gf.Composer, df)
		}
	}
	sort.Slice(gf.Core, func(i, j int) bool { return gf.Core[i].Title < gf.Core[j].Title })
	sort.Slice(gf.Composer, func(i, j int) bool { return gf.Composer[i].Title < gf.Composer[j].Title })
	sort.Slice(gf.Audio, func(i, j int) bool { return gf.Audio[i].Title < gf.Audio[j].Title })

	out, err := os.Create("../../docs/dsl.md")
	checkErr(err)
	defer out.Close()
	err = tmpl.Execute(out, gf)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func firstUpcaseAndDot(s string) string {
	b := bytes.Buffer{}
	last := ' '
	for i, each := range []rune(s) {
		if i == 0 {
			b.WriteRune(unicode.ToUpper(each))
		} else {
			b.WriteRune(each)
		}
		last = each
	}
	if last != '.' {
		b.WriteRune('.')
	}
	return b.String()
}
