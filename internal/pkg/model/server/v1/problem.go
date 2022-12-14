package v1

type Problem struct {
	BaseModel
	SchoolId uint64 `json:"schoolId,omitempty" gorm:"column:schoolId;not null"`
	// 所属course
	CourseName string `json:"courseName,omitempty" gorm:"column:courseName;size:64;not null"`
	// 题目类型选择，填空，大题
	Category   string `json:"category,omitempty" gorm:"column:category;size:256"`
	Title      string `json:"title" gorm:"column:title;size:256;not null;unique"`
	Question   string `json:"question" gorm:"column:question"`
	Ans        string `json:"ans" gorm:"column:ans"`
	Cnt        int    `gorm:"column:cnt"`
	Difficulty int8   `json:"difficulty" gorm:"column:difficulty"`
	Tags       []Tag  `json:"tags" gorm:"many2many:problem_tags;"`
	// 不存数据库，但是可以用于返回
	Pass   bool `json:"pass" gorm:"-"`
	Favour bool `json:"favour" gorm:"-"`

	Solutions []Solution `json:"solutions" gorm:"many2many:problem_solution"`
}

func (s *Problem) TableName() string {
	return "problem"
}

type Tag struct {
	Name string `json:"tagName,omitempty" gorm:"primary_key;AUTO_INCREMENT;column:name;size:64"`
}

func (t *Tag) TableName() string {
	return "tag"
}

type ProblemTag struct {
	ProblemId uint64 `gorm:"column:problem_id"`
	TagName   string `gorm:"column:tag_name"`
}

func (pt *ProblemTag) TableName() string {
	return "problemTag"
}
