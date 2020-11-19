package controllers

import (
	"context"
	"strconv"

	"github.com/chat2542/app/ent"
	"github.com/chat2542/app/ent/scholarshiptype"
	"github.com/gin-gonic/gin"
)
// ScholarshiptypeController creates and registers handles for the scholarshiptype controller
type ScholarshiptypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateScholarshiptype handles POST requests for adding scholarshiptype entities
// @Summary Create scholarshiptype
// @Description Create scholarshiptype
// @ID create-scholarshiptype
// @Accept   json
// @Produce  json
// @Param scholarshiptype body ent.Scholarshiptype true "Scholarshiptype entity"
// @Success 200 {object} ent.Scholarshiptype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshiptype [post]
func (ctl *ScholarshiptypeController) CreateScholarshiptype(c *gin.Context) {
	obj := ent.Scholarshiptype{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Scholarshiptype binding failed",
		})
		return
	}

	st, err := ctl.client.Scholarshiptype.
		Create().
		SetTypeName(obj.TypeName).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, st)
}

// GetScholarshiptype handles GET requests to retrieve a scholarshiptype entity
// @Summary Get a scholarshiptype entity by ID
// @Description get scholarshiptype by ID
// @ID get-scholarshiptype
// @Produce  json
// @Param id path int true "Scholarshiptype ID"
// @Success 200 {object} ent.Scholarshiptype
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshiptype/{id} [get]
func (ctl *ScholarshiptypeController) GetScholarshiptype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	st, err := ctl.client.Scholarshiptype.
		Query().
		Where(scholarshiptype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, st)
}

// ListScholarshiptype handles request to get a list of scholarshiptype entities
// @Summary List scholarshiptype entities
// @Description list scholarshiptype entities
// @ID list-scholarshiptype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Scholarshiptype
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /scholarshiptype [get]
func (ctl *ScholarshiptypeController) ListScholarshiptype(c *gin.Context) {
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

	scholarshiptype, err := ctl.client.Scholarshiptype.
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

	c.JSON(200, scholarshiptype)
}

// NewScholarshiptypeController creates and registers handles for the scholarshiptype controller
func NewScholarshiptypeController(router gin.IRouter, client *ent.Client) *ScholarshiptypeController {
	stc := &ScholarshiptypeController{
		client: client,
		router: router,
	}

	stc.register()

	return stc

}

func (ctl *ScholarshiptypeController) register() {
	scholarshiptype := ctl.router.Group("/scholarshiptype")

	scholarshiptype.POST("", ctl.CreateScholarshiptype)
	scholarshiptype.GET(":id", ctl.GetScholarshiptype)
	scholarshiptype.GET("", ctl.ListScholarshiptype)

}
