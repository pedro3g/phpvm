package types

type Release struct {
	Tag    string `json:"tag"`
	Source source `json:"source"`
}

type source map[string]string
