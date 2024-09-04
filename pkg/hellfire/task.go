package hellfire

import "github.com/kernelpanic77/hellfire/pkg/hellfire/http"

type task func(t *T, client *http.Client) bool
