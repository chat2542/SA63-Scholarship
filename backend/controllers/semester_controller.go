package controllers

import (
	"context"
	"strconv"

	"github.com/chat2542/app/ent"
	"github.com/chat2542/app/ent/semester"
	"github.com/gin-gonic/gin"
)
// SemesterController creates and registers handles for the semester controller
type SemesterController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSemester handles POST requests for adding semester entities
// @Summary Create semester
// @Description Create semester
// @ID create-semester
// @Accept   json
// @Produce  json
// @Param semester body ent.Semester true "Semester entity"
// @Success 200 {object} ent.Semester
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /resolutions [post]
func (ctl *SemesterController) CreateSemester(c *gin.Context) {
	obj := ent.Semester{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Semester binding failed",
		})
		return
	}

	s, err := ctl.client.Semester.
		Create().
		SetSemester(obj.Semester).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetSemester handles GET requests to retrieve a semester entity
// @Summary Get a semester entity by ID
// @Description get semester by ID
// @ID get-semester
// @Produce  json
// @Param id path int true "Semester ID"
// @Success 200 {object} ent.Semester
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /semester/{id} [get]
func (ctl *SemesterController) GetSemester(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := ctl.client.Semester.
		Query().
		Where(semester.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListSemester handles request to get a list of semester entities
// @Summary List semester entities
// @Description list semester entities
// @ID list-semester
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Semester
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /semester [get]
func (ctl *SemesterController) ListSemester(c *gin.Context) {
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

	semester, err := ctl.client.Semester.
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

	c.JSON(200, semester)
}

// NewSemesterController creates and registers handles for the semester controller
func NewSemesterController(router gin.IRouter, client *ent.Client) *SemesterController {
	sc := &SemesterController{
		client: client,
		router: router,
	}

	sc.register()

	return sc

}

func (ctl *SemesterController) register() {
	semester := ctl.router.Group("/semester")

	semester.POST("", ctl.CreateSemester)
	semester.GET(":id", ctl.GetSemester)
	semester.GET("", ctl.ListSemester)

}
