package types

type Commit struct {
	Files   []File `json:"Files"`
	Message string `json:"Message"`
	// Branch  string `json:"Branch"`
}

type File struct {
	Name string `json:"Name"`
	Path string `json:"Path"`
}
