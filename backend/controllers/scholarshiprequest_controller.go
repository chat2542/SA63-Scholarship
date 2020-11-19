package controllers

import (
	"context"
	"strconv"

	"github.com/chat2542/app/ent"
	"github.com/chat2542/app/ent/scholarshipinformation"
	"github.com/chat2542/app/ent/scholarshiptype"
	"github.com/chat2542/app/ent/semester"
	"github.com/chat2542/app/ent/user"
	"github.com/gin-gonic/gin"
)

// ScholarshiprequestController creates and registers handles for the scholarshiprequest controller
type ScholarshiprequestController struct {
	client *ent.Client
	router gin.IRouter
}

// Scholarshiprequest creates and registers handles for the scholarshiprequest
type Scholarshiprequest struct {
	User                   int
	Scholarshiptype        int
	Scholarshipinformation int
	Semester               int
}

// CreateScholarshiprequest handles POST requests for adding scholarshiprequest entities
// @Summary Create scholarshiprequest
// @Description Create scholarshiprequest
// @ID create-scholarshiprequest
// @Accept   json
// @Produce  json
// @Param scholarshiprequest body Scholarshiprequest true "ScholarshipRequest entity"
// @Success 200 {object} ent.ScholarshipRequest
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshiprequests [post]
func (ctl *ScholarshiprequestController) CreateScholarshiprequest(c *gin.Context) {
	obj := Scholarshiprequest{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "scholarshiprequest video binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	st, err := ctl.client.Scholarshiptype.
		Query().
		Where(scholarshiptype.IDEQ(int(obj.Scholarshiptype))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "scholarshiptype not found",
		})
		return
	}

	si, err := ctl.client.Scholarshipinformation.
		Query().
		Where(scholarshipinformation.IDEQ(int(obj.Scholarshipinformation))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "scholarshipinformation not found",
		})
		return
	}

	s, err := ctl.client.Semester.
		Query().
		Where(semester.IDEQ(int(obj.Semester))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "semester not found",
		})
		return
	}

	sq, err := ctl.client.ScholarshipRequest.
		Create().
		SetUser(u).
		SetScholarshiptype(st).
		SetScholarshipinformation(si).
		SetSemester(s).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   sq,
	})
}

// ListScholarshiprequest handles request to get a list of scholarshiprequest entities
// @Summary List scholarshiprequest entities
// @Description list scholarshiprequest entities
// @ID list-scholarshiprequest
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ScholarshipRequest
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshiprequests [get]
func (ctl *ScholarshiprequestController) ListScholarshiprequest(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	scholarshiprequest, err := ctl.client.ScholarshipRequest.
		Query().
		WithUser().
		WithScholarshipinformation().
		WithScholarshiptype().
		WithSemester().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, scholarshiprequest)
}

// NewScholarshiprequestController creates and registers handles for the scholarshiprequest controller
func NewScholarshiprequestController(router gin.IRouter, client *ent.Client) *ScholarshiprequestController {
	pvc := &ScholarshiprequestController{
		client: client,
		router: router,
	}

	pvc.register()

	return pvc

}

func (ctl *ScholarshiprequestController) register() {
	scholarshiprequest := ctl.router.Group("/scholarshiprequests")

	scholarshiprequest.POST("", ctl.CreateScholarshiprequest)
	scholarshiprequest.GET("", ctl.ListScholarshiprequest)

}
