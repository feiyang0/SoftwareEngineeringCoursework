package v1

type ProblemListOption struct {
	Category   string  `json:"category"`
	CourseName string  `json:"courseName"`
	Orders     []Order `json:"orders"`

	Tag string `json:"tag"`

	SearchKeyWords string `json:"searcherKeyWords"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Order struct {
	OrderBy   string `json:"orderBy"`
	SortOrder string `json:"sortOrder"`
}
