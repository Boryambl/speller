package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"go.uber.org/zap"
	"homework/internal/mistakes"
	"net/http"
)

type routes struct {
	l *zap.Logger
}

func (r routes) correctMistakes(c *gin.Context) {
	var text mistakes.Text
	err := json.NewDecoder(c.Request.Body).Decode(&text)
	if err != nil {
		r.l.Error("text not found in request " + "routes - correctMistakes")
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("text not found in request"))
		return
	}
	result := mistakes.CorrectMistakes(text)
	c.JSON(http.StatusOK, result)
}
