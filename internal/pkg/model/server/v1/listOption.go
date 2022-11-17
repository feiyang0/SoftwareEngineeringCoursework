package v1

type ListOption struct {
	Category  string `json:"category"`
	OrderBy   string `json:"orderBy"`
	SortOrder string `json:"sortOrder"`
	Tags      []Tag  `json:"tags"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
}
