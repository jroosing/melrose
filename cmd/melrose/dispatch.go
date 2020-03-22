package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/emicklei/melrose"
)

var assignmentRegex = regexp.MustCompile(`^([a-z]+\[[0-9]+\])=.*$`)

func dispatch(entry string) error {
	if len(entry) == 0 {
		fmt.Println()
		return nil
	}
	if value, ok := varStore.Get(entry); ok {
		printValue(value)
		return nil
	}
	// is assignment?
	// TODO not correct
	if strings.Contains(entry, "=") {
		parts := strings.Split(entry, "=")
		variable := strings.TrimSpace(parts[0])
		expression := parts[1]
		r, err := eval(expression)
		if err != nil {
			return err
		}
		// TODO check that we do not use a function name as variable
		varStore.Put(variable, r)
		printValue(r)
		return nil
	}
	// evaluate and print
	r, err := eval(entry)
	if err != nil {
		return err
	}
	printValue(r)
	return nil
}

func printValue(v interface{}) {
	if v == nil {
		fmt.Println()
		return
	}
	if s, ok := v.(melrose.Storable); ok {
		fmt.Printf("(%T) %s\n", v, s.Storex())
	} else {
		fmt.Printf("(%T) %v\n", v, v)
	}
}

func eval(entry string) (interface{}, error) {
	// flatten multiline ; expr does not support multiline strings
	entry = strings.Replace(entry, "\n", " ", -1)

	env := map[string]interface{}{}
	for k, f := range evalFunctions() {
		env[k] = f.Func
	}
	for k, v := range varStore.Variables() {
		env[k] = v
	}
	program, err := expr.Compile(entry, expr.Env(env))
	if err != nil {
		return nil, err
	}
	return expr.Run(program, env)
}
