package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const AMOUNT_OF_USERS = 999_999

type UserList struct {
	Amount int    `json:"amount"`
	Type   string `json:"type"`
}

type Students struct {
	UserList
	Students []Student `json:"students"`
}

type Teachers struct {
	UserList
	Teachers []Teacher `json:"teachers"`
}

type AllPersonals struct {
	Teachers Teachers `json:"teachers-user"`
	Students Students `json:"students-user"`
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

var allIDDB = []string{}

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

var studentsDataBase = []Student{}

func (students *Students) allStudentsList() {
	for _, student := range students.Students {
		student.changeID()
		studentsDataBase = append(studentsDataBase, student)

	}

}

func (teachers *Teachers) allTeachersList() {
	for _, teacher := range teachers.Teachers {
		//fmt.Println(teacher)
		teacher.changeID()
		teachersDataBase = append(teachersDataBase, teacher)

	}

}

func (teacher *Teacher) changeID() {

	//fmt.Println(teacher.ID)
	teacher.ID = generateID()
	//fmt.Println(teacher.ID)

	//fmt.Println(teachersDataBase)
	allIDDB = append(allIDDB, teacher.ID)
}

func (student *Student) changeID() {
	student.ID = generateID()

	allIDDB = append(allIDDB, student.ID)

}

func checkID(id string) bool {

	for _, v := range allIDDB {
		if id == v {
			return true
		}
	}
	return false
}

func generateID() string {
	randomNumber := rand.Intn(AMOUNT_OF_USERS)
	randomNumberStr := strconv.Itoa(randomNumber)

	//fmt.Println(len(strconv.Itoa(AMOUNT_OF_USERS)))

	if len(randomNumberStr) < len(strconv.Itoa(AMOUNT_OF_USERS)) {
		buffer := len(strconv.Itoa(AMOUNT_OF_USERS)) - len(randomNumberStr)

		bufferStr := strings.Repeat("0", buffer)

		randomNumberStr = bufferStr + randomNumberStr

	}

	flag := checkID(randomNumberStr)

	if flag {
		return generateID()
	}

	return randomNumberStr

}

func getTeachersList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teachersDataBase)
}

func getStudentsList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, studentsDataBase)
}

func getAllList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, teachersDataBase)
	c.IndentedJSON(http.StatusOK, studentsDataBase)
}

func createStudent(c *gin.Context) {
	var newStudent Student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	studentsDataBase = append(studentsDataBase, newStudent)

	newStudent.changeID()

	c.IndentedJSON(http.StatusCreated, newStudent)

}

func createTeacher(c *gin.Context) {
	var newTeacher Teacher

	if err := c.BindJSON(&newTeacher); err != nil {
		return
	}
	newTeacher.changeID()
	teachersDataBase = append(teachersDataBase, newTeacher)
	c.IndentedJSON(http.StatusCreated, newTeacher)
}

func createUsers(c *gin.Context) {
	var allUsers AllPersonals
	fmt.Println("here")
	if err := c.BindJSON(&allUsers); err != nil {
		return
	}
	var newTeachers = allUsers.Teachers

	newTeachers.allTeachersList()

	c.IndentedJSON(http.StatusCreated, newTeachers)

	var newStudents = allUsers.Students

	newStudents.allStudentsList()

	c.IndentedJSON(http.StatusCreated, newStudents)

	c.IndentedJSON(http.StatusCreated, allUsers)

}

func createStudentsList(c *gin.Context) {
	var studentsList Students
	if err := c.BindJSON(&studentsList); err != nil {
		return
	}

	studentsList.allStudentsList()

	c.IndentedJSON(http.StatusCreated, studentsList)

}

func createTeachersList(c *gin.Context) {
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

	router.GET("/AllDataBases", getAllList)

	router.POST("/AllDataBases/createUsersList", createUsers)

	router.GET("/TeacherDataBase", getTeachersList)

	router.POST("/TeachersDataBase/addTeacher", createTeacher)

	router.POST("/TeachersDataBase/addTeachersList", createTeachersList)

	router.GET("/StudentsDataBase", getStudentsList)

	router.POST("/StudentsDataBase/addStudent", createStudent)

	router.POST("/StudentsDataBase/addStudentsList", createStudentsList)

	router.Run("localhost:8080")
}
