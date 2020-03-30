package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/emicklei/melrose/dsl"
	"github.com/emicklei/melrose/notify"
)

func showHelp(entry string) notify.Message {
	var b bytes.Buffer
	io.WriteString(&b, "\n")
	{
		funcs := dsl.EvalFunctions(varStore)
		keys := []string{}
		width := 0
		for k, f := range funcs {
			if f.ControlsAudio {
				continue
			}
			if len(k) > width {
				width = len(k)
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			f := funcs[k]
			fmt.Fprintf(&b, "%s --- %s\n", strings.Repeat(" ", width-len(k))+k, f.Description)
		}
	}
	io.WriteString(&b, "\n")
	{
		funcs := dsl.EvalFunctions(varStore)
		keys := []string{}
		width := 0
		for k, f := range funcs {
			if !f.ControlsAudio {
				continue
			}
			if len(k) > width {
				width = len(k)
			}
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			f := funcs[k]
			fmt.Fprintf(&b, "%s --- %s\n", strings.Repeat(" ", width-len(k))+k, f.Description)
		}
	}
	io.WriteString(&b, "\n")
	{
		cmds := cmdFunctions()
		keys := []string{}
		for k, _ := range cmds {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			c := cmds[k]
			fmt.Fprintf(&b, "%s --- %s\n", k, c.Description)
		}
	}
	return notify.Infof("%s", b.String())
}