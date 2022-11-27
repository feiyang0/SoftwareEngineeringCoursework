package v1

type ProblemListOption struct {
	Category   string  `json:"category,omitempty"`
	CourseName string  `json:"courseName,omitempty"`
	Orders     []Order `json:"orders,omitempty"`

	Tag string `json:"tag,omitempty"`

	SearchKeyWords string `json:"searcherKeyWords,omitempty"`

	Limit  int `json:"limit" binding:"required"`
	Offset int `json:"offset" binding:"required"`
}

type Order struct {
	OrderBy   string `json:"orderBy"`
	SortOrder string `json:"sortOrder"`
}
