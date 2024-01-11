package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type List struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

type Students struct {
	List

	Students []Student `json:"students"`
}

type Teachers struct {
	List
	Teachers []Teacher `json:"teachers"`
}

type AllPersonals struct {
	List
	Teachers []Teacher  `json:"teachers"`
	Students []Students `json:"students"`
}

type Personal struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	Student bool   `json:"student"`
	Teacher bool   `json:"teacher"`
}

type Student struct {
	Personal
	Major      string   `json:"major"`
	Gpa        float64  `json:"gpa"`
	Courses    []string `json:"courses"`
	Professors []string `json:"professors"`
}

type Staff struct {
	Personal
	Wage            float64 `json:"wage"`
	StartDate       string  `json:"start-date"`
	ContractEndDate string  `json:"contract-end-date"`
}

type Teacher struct {
	Staff
	CourseTeaching []string `json:"course-teaching"`
	Rating         float32  `json:"rating"`
	Filed          string   `json:"filed"`
}

var teachersDataBase = []Teacher{

	{Rating: 3.2, CourseTeaching: []string{"English 1000", "English 2000"}, Filed: "English", Staff: Staff{
		Personal: Personal{
			ID:      "1",
			Name:    "Jack",
			Age:     45,
			Gender:  "Male",
			Student: false,
			Teacher: true,
		},
		Wage:            90000,
		StartDate:       "1-2-2018",
		ContractEndDate: "23-9-2031",
	}},
}

var studentDataBase = []Student{}

func (students Students) allStudentsList() {
	for _, student := range students.Students {
		studentDataBase = append(studentDataBase, student)
	}

}

func (teachers Teachers) allTeachersList() {
	for _, teacher := range teachers.Teachers {
		fmt.Println(teacher)
		teachersDataBase = append(teachersDataBase, teacher)
	}

}

func getTeachersList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teachersDataBase)
}

func getStudentList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, studentDataBase)
}

func getAllList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teachersDataBase)
	c.IndentedJSON(http.StatusOK, studentDataBase)
}

func createStudent(c *gin.Context) {
	var newStudent Student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	studentDataBase = append(studentDataBase, newStudent)

	c.IndentedJSON(http.StatusCreated, newStudent)

}

func createTeacher(c *gin.Context) {
	var newTeacher Teacher

	if err := c.BindJSON(&newTeacher); err != nil {
		return
	}

	teachersDataBase = append(teachersDataBase, newTeacher)
	c.IndentedJSON(http.StatusCreated, newTeacher)
}

func createUser(c *gin.Context) {

}

func createStudentList(c *gin.Context) {
	var studentsList Students
	if err := c.BindJSON(&studentsList); err != nil {
		return
	}

	studentsList.allStudentsList()

	c.IndentedJSON(http.StatusCreated, studentsList)

}

func createTeacherList(c *gin.Context) {
	var teachersList Teachers

	if err := c.BindJSON(&teachersList); err != nil {
		return
	}

	teachersList.allTeachersList()

	c.IndentedJSON(http.StatusCreated, teachersList)

}

// need some type of internet connection to run **

// curl localhost:8080/TeacherDataBase --include --header "Content-Type: application/json" -d @teacherBody.json --request "POST"
func main() {
	fmt.Println("Launch API")
	router := gin.Default()

	router.GET("/AllDataBase", getAllList)

	router.POST("/AllDataBase", createUser)

	router.GET("/TeacherDataBase", getTeachersList)

	router.POST("/TeacherDataBase/addTeacher", createTeacher)

	router.POST("/TeacherDataBase/addTeacherList", createTeacherList)

	router.GET("/StudentDataBase", getStudentList)

	router.POST("/StudentDataBase/addStudent", createStudent)

	router.POST("/StudentDataBase/addStudentsList", createStudentList)

	router.Run("localhost:8080")
}
