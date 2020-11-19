package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chat2542/app/ent"
	"github.com/chat2542/app/ent/scholarshipinformation"
	"github.com/gin-gonic/gin"
)

// ScholarshipinformationController creates and registers handles for the Scholarshipinformation controller
type ScholarshipinformationController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateScholarshipinformation handles POST requests for adding scholarshipinformation entities
// @Summary Create scholarshipinformation
// @Description Create scholarshipinformation
// @ID create-scholarshipinformation
// @Accept   json
// @Produce  json
// @Param playlist body ent.Scholarshipinformation true "Scholarshipinformation entity"
// @Success 200 {object} ent.Scholarshipinformation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshipinformation [post]
func (ctl *ScholarshipinformationController) CreateScholarshipinformation(c *gin.Context) {
	obj := ent.Scholarshipinformation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "scholarshipinformation binding failed",
		})
		return
	}

	si, err := ctl.client.Scholarshipinformation.
		Create().
		SetScholarshipName(obj.ScholarshipName).
		Save(context.Background())

	fmt.Println(err)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, si)
}

// GetScholarshipinformation handles GET requests to retrieve a scholarshipinformation entity
// @Summary Get a scholarshipinformation entity by ID
// @Description get scholarshipinformation by ID
// @ID get-scholarshipinformation
// @Produce  json
// @Param id path int true "Scholarshipinformation ID"
// @Success 200 {object} ent.Scholarshipinformation
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshipinformation/{id} [get]
func (ctl *ScholarshipinformationController) GetScholarshipinformation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	si, err := ctl.client.Scholarshipinformation.
		Query().
		Where(scholarshipinformation.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, si)
}

// ListScholarshipinformation handles request to get a list of scholarshipinformation entities
// @Summary List scholarshipinformation entities
// @Description list scholarshipinformation entities
// @ID list-scholarshipinformation
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Scholarshipinformation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshipinformation [get]
func (ctl *ScholarshipinformationController) ListScholarshipinformation(c *gin.Context) {
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

	scholarshipinformation, err := ctl.client.Scholarshipinformation.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, scholarshipinformation)
}

// NewScholarshipinformationController creates and registers handles for the Scholarshipinformation controller
func NewScholarshipinformationController(router gin.IRouter, client *ent.Client) *ScholarshipinformationController {
	sic := &ScholarshipinformationController{
		client: client,
		router: router,
	}

	sic.register()

	return sic

}

func (ctl *ScholarshipinformationController) register() {
	scholarshipinformation := ctl.router.Group("/scholarshipinformation")

	scholarshipinformation.POST("", ctl.CreateScholarshipinformation)
	scholarshipinformation.GET(":id", ctl.GetScholarshipinformation)
	scholarshipinformation.GET("", ctl.ListScholarshipinformation)

}
