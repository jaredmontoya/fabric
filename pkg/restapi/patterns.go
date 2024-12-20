package restapi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/danielmiessler/fabric/pkg/plugins/db/fsdb"
)

// PatternsHandler defines the handler for patterns-related operations
type PatternsHandler struct {
	*StorageHandler[fsdb.Pattern]
	patterns *fsdb.PatternsEntity
}

// NewPatternsHandler creates a new PatternsHandler
func NewPatternsHandler(r *gin.Engine, patterns *fsdb.PatternsEntity) (ret *PatternsHandler) {
	ret = &PatternsHandler{
		StorageHandler: NewStorageHandler[fsdb.Pattern](r, "patterns", patterns), patterns: patterns}

	// TODO: Add custom, replacement routes here
	//r.GET("/patterns/:name", ret.Get)
	return
}

// Get handles the GET /patterns/:name route
func (h *PatternsHandler) Get(c *gin.Context) {
	name := c.Param("name")
	variables := make(map[string]string) // Assuming variables are passed somehow
	pattern, err := h.patterns.GetApplyVariables(name, variables)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, pattern)
}
