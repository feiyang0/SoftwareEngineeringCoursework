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
	opts := &v1.ProblemListOption{
		Category: "填空",
		//CourseName: "学科分类",
		////Orders: []v1.Order{
		////	{OrderBy: "cnt", SortOrder: "asc"},
		////},
		Tag:            "t1",
		SearchKeyWords: "题目",
		Limit:          10,
		Offset:         0,
	}
	ps, err := p.GetAllWithTag(11, opts)
	for _, problem := range ps {
		fmt.Println(problem)
	}
}
