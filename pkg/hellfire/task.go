package hellfire

import (
	http "github.com/kernelpanic77/hellfire/pkg/hellfire/http"
)

type task func(t *T, client *http.Client) bool

// type iteration func(t *T, client *hellfire_http.Client) bool
