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

func TestSolution(t *testing.T) {
	initDb()
	dbIns, err := GetMySQLFactoryOr()
	if err != nil {
		fmt.Println("dbIns err:", err)
	}
	s := dbIns.Solutions()

	// create solution
	//s.Create(12, &v1.Solution{
	//	SchoolId: 123,
	//	Content:  "题解内容",
	//})

	// solution list
	//opts := &v1.SolutionListOption{
	//	Pid:    12,
	//	Offset: 0,
	//	Limit:  2,
	//}
	//ss, number, _ := s.GetSolutionList(opts)
	//fmt.Println("total number:", number)
	//for _, solu := range ss {
	//	fmt.Println(solu)
	//}
	//

	//// 修改题解
	//solu := &v1.Solution{
	//	SchoolId: 123,
	//	Content:  "修改题解内容",
	//}
	//solu.ID = 10
	//s.Update(solu)

	////添加评论
	//s.AddComment(1, &v1.Comment{
	//	SchoolId: 123,
	//	Content:  "这是评论1",
	//})
	//s.AddComment(1, &v1.Comment{
	//	SchoolId: 123,
	//	Content:  "这是评论2",
	//})

	// 删除评论
	s.DelComment(1, 3)
	// 获取题解
	solution, _ := s.GetSolution(1)
	fmt.Println(solution.Comments)

}
