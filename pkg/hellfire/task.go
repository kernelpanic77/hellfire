package hellfire

import "github.com/kernelpanic77/hellfire/common"

type Test interface {
	Check(val interface{}, checks CheckFuncMap, tag string) bool
	Fatal(msg string)
	Log(s string)
	ProgressBar()
}

type Task func(t Test, client common.Client) bool

// type iteration func(t *T, client *hellfire_http.Client) bool
type CheckFunc func(interface{}) bool

type CheckFuncMap map[string]CheckFunc