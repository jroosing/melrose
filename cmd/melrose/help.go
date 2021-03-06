package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/emicklei/melrose"
	"github.com/emicklei/melrose/dsl"
	"github.com/emicklei/melrose/notify"
)

func showHelp(args []string) notify.Message {
	var b bytes.Buffer

	fmt.Fprintf(&b, "\nversion %s, syntax: %s\n", version, dsl.Syntax)
	fmt.Fprintf(&b, "https://emicklei.github.io/melrose \n")

	// detect help for a command or function
	if len(args) > 0 {
		cmdfunc := strings.TrimSpace(args[0])
		if cmd, ok := cmdFunctions()[cmdfunc]; ok {
			fmt.Fprintf(&b, "%s\n", cmdfunc)
			fmt.Fprintf(&b, "%s\n", cmd.Description)
			fmt.Fprintf(&b, "%s\n", cmd.Sample)
			return notify.Infof("%s", b.String())
		}
		if fun, ok := dsl.EvalFunctions(globalStore, melrose.NoLooper)[cmdfunc]; ok {
			fmt.Fprintf(&b, "%s\n", cmdfunc)
			fmt.Fprintf(&b, "%s\n", fun.Description)
			fmt.Fprintf(&b, "%s\n", fun.Template)
			return notify.Infof("%s", b.String())
		}
	}
	io.WriteString(&b, "\n")
	{
		funcs := dsl.EvalFunctions(globalStore, melrose.NoLooper)
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
		funcs := dsl.EvalFunctions(globalStore, melrose.NoLooper)
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
