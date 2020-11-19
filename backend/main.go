package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chat2542/app/controllers"
	"github.com/chat2542/app/ent"
	"github.com/chat2542/app/ent/scholarshiptype"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	Email string
}

type Scholarshiptypes struct {
	Scholarshiptype []Scholarshiptype
}
type Scholarshiptype struct {
	TypeName string
}

type Scholarshipinformations struct {
	Scholarshipinformation []Scholarshipinformation
}
type Scholarshipinformation struct {
	ScholarshipName                       string
	ScholarshiptypeScholarshipinformation int
}
type Semesters struct {
	Semester []Semester
}
type Semester struct {
	Semester string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewScholarshipinformationController(v1, client)
	controllers.NewScholarshiptypeController(v1, client)
	controllers.NewSemesterController(v1, client)
	controllers.NewUserController(v1, client)
	controllers.NewScholarshiprequestController(v1, client)

	// Set User Data
	users := Users{
		User: []User{
			User{"Chatmongkhon2542@gmail.com"},
			User{"B6104887@g.sut.ac.th"},
		},
	}
	for _, u := range users.User {

		client.User.
			Create().
			SetEmail(u.Email).
			Save(context.Background())
	}

	// Set  Scholarshiptype Data
	scholarshiptypes := Scholarshiptypes{
		Scholarshiptype: []Scholarshiptype{
			Scholarshiptype{"แบบต่อเนื่อง"},
			Scholarshiptype{"แบบไม่ต่อเนื่อง"},
		},
	}
	for _, st := range scholarshiptypes.Scholarshiptype {
		client.Scholarshiptype.
			Create().
			SetTypeName(st.TypeName).
			Save(context.Background())
	}

	// Set Scholarshipinformation Data
	scholarshipinformations := Scholarshipinformations{
		Scholarshipinformation: []Scholarshipinformation{
			Scholarshipinformation{"ทุนกรอ.", 1},
			Scholarshipinformation{"ทุนกยศ.", 1},
			Scholarshipinformation{"ทุนผู้ยากไร้", 2},
		},
	}
	for _, si := range scholarshipinformations.Scholarshipinformation {

		st, err := client.Scholarshiptype.
			Query().
			Where(scholarshiptype.IDEQ(int(si.ScholarshiptypeScholarshipinformation))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		client.Scholarshipinformation.
			Create().
			SetScholarshipName(si.ScholarshipName).
			SetScholarshiptype(st).
			Save(context.Background())
	}
	// Set Semester Data
	semesters := Semesters{
		Semester: []Semester{
			Semester{"1/2563"},
			Semester{"2/2563"},
			Semester{"3/2563"},
		},
	}
	for _, s := range semesters.Semester {
		client.Semester.
			Create().
			SetSemester(s.Semester).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
