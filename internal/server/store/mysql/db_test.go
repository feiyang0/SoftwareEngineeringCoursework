package mysql

import (
	v1 "SoftwareEngine/internal/pkg/model/server/v1"
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func initDb() {
	viper.SetConfigFile("../../../../config/server.yaml")
	viper.ReadInConfig()
}
func TestStudent(t *testing.T) {
	initDb()
	dbIns, err := GetMySQLFactoryOr()
	if err != nil {
		fmt.Println("dbIns err:", err)
	}
	stu := dbIns.Students()

	stu.Commit(114, 11)

}

func TestProblems(t *testing.T) {
	initDb()
	dbIns, err := GetMySQLFactoryOr()
	if err != nil {
		fmt.Println("dbIns err:", err)
	}
	p := dbIns.Problems()
	problem := &v1.Problem{
		SchoolId:   123,
		CourseName: "软工",
		Category:   "学科名",
		Title:      "abc",
		Question:   "题目内容",
		Difficulty: 1,
		Tags: []v1.Tag{
			{Name: "t1"},
			{Name: "t2"},
		},
	}
	//p.Create(problem)
	//pb, _ := p.GetProblem(36)
	problem.ID = 36
	p.Update(problem)
	pb, _ := p.GetProblem(36)
	fmt.Println(pb.Tags)
	//p.Delete(12)
	//opts := &v1.ProblemListOption{
	//	//Category:   "选择",
	//	//CourseName: "学科分类",
	//	Orders: []v1.Order{
	//		{OrderBy: "cnt", SortOrder: "asc"},
	//	},
	//	Tag: "t2",
	//	//SearchKeyWords: "题目",
	//	Limit:  10,
	//	Offset: 0,
	//}
	//ps, err := p.GetAllWithTag(11, opts)
	//for _, pm := range ps {
	//	fmt.Println(pm)
	//}
}
