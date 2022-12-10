package v1

type ProblemListOption struct {
	Category   string  `json:"category,omitempty"`
	CourseName string  `json:"courseName,omitempty"`
	Orders     []Order `json:"orders,omitempty"`

	Difficulty int8 `json:"difficulty,omitempty"`

	Tag string `json:"tag,omitempty"`

	SearchKeyWords string `json:"searcherKeyWords,omitempty"`

	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type Order struct {
	OrderBy   string `json:"orderBy"`
	SortOrder string `json:"sortOrder"`
}

type SolutionListOption struct {
	Pid    uint64
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
