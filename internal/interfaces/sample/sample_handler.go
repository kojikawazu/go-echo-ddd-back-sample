package interfaces_sample

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SampleHandler struct {
}

func NewSampleHandler() *SampleHandler {
	return &SampleHandler{}
}

// 動作確認用
func (h *SampleHandler) ExecSample(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
