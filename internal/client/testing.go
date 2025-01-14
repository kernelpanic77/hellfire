package client

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kernelpanic77/hellfire/common"
	"github.com/schollz/progressbar/v3"
)

type T struct {
	
}

func (t *T) ProgressBar() {
	bar := progressbar.Default(100)
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(40 * time.Millisecond)
	}
}

func (t *T) Log(s string) {
	log.Println(s)
}

func (t *T) Check(val interface{}, checks common.CheckFuncMap, tag string) bool {
	check_val := true

	for check_name, check_func := range checks {
		curr_check := check_func(val)
		if !curr_check {
			// t.("%s has failed!", check_name)
			t.Log(fmt.Sprintf("%s check has failed!", check_name))
		}
		check_val = check_val && curr_check
	}

	return check_val
}

func (t *T) Fatal(msg string) {
	t.Log(msg)
	panic(errors.New(msg))
}