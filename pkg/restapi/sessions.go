package restapi

import (
	"github.com/gin-gonic/gin"

	"github.com/danielmiessler/fabric/pkg/plugins/db/fsdb"
)

// SessionsHandler defines the handler for sessions-related operations
type SessionsHandler struct {
	*StorageHandler[fsdb.Session]
	sessions *fsdb.SessionsEntity
}

// NewSessionsHandler creates a new SessionsHandler
func NewSessionsHandler(r *gin.Engine, sessions *fsdb.SessionsEntity) (ret *SessionsHandler) {
	ret = &SessionsHandler{
		StorageHandler: NewStorageHandler[fsdb.Session](r, "sessions", sessions), sessions: sessions}
	return ret
}
