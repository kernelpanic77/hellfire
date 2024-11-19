package hellfire

type Params struct {
	timeout   int32
	cookies   map[string]string
	redirects int32
	headers   map[string]string
}
