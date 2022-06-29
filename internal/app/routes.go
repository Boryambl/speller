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
		r.l.Error(fmt.Sprintf("text not found ") + "routes - correctMistakes")
		c.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Errorf("text not found"))
		return
	}
	result, err := mistakes.CorrectMistakes(text.Text, r.l)
	if err != nil {
		r.l.Error(err.Error() + "routes - correctMistakes")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, result)
}
