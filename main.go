package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/lithdew/quickjs"
)

func Check(err error) {
	if err != nil {
		var evalErr *quickjs.Error
		if errors.As(err, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
		panic(err)
	}
}

func main() {

	runtime.LockOSThread()
	source := os.Args[1:][0]

	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	globals := context.Globals()

	globals.Set("__dispatch", context.Function(DoneNS))

	snap, _ := Asset("target/done.js")

	k, e := context.Eval(string(snap))
	Check(e)
	defer k.Free()

	bundle := BundleModule(source)
	a := func(val quickjs.Value) {
		if !val.IsUndefined() {
			return
		}
		fmt.Println(val)
	}
	dat, e := ioutil.ReadFile(source)
	if e != nil {
		panic(e)
	}
	go Compile(string(dat), a)
	result, e := context.EvalFile(bundle, source)

	defer result.Free()
	if e != nil {
		var evalErr *quickjs.Error
		if errors.As(e, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
		panic(e)
	}
}
