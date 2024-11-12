package restapi

import (
	"github.com/gin-gonic/gin"

	"github.com/danielmiessler/fabric/pkg/plugins/db/fsdb"
)

// ContextsHandler defines the handler for contexts-related operations
type ContextsHandler struct {
	*StorageHandler[fsdb.Context]
	contexts *fsdb.ContextsEntity
}

// NewContextsHandler creates a new ContextsHandler
func NewContextsHandler(r *gin.Engine, contexts *fsdb.ContextsEntity) (ret *ContextsHandler) {
	ret = &ContextsHandler{
		StorageHandler: NewStorageHandler[fsdb.Context](r, "contexts", contexts), contexts: contexts}
	return
}