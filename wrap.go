package wrap

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	APIResponse struct {
		Status      int         `json:"status"`
		Message     string      `json:"message"`
		Description string      `json:"description,omitempty"`
		Data        interface{} `json:"data,omitempty"`
	}
)

func OK(
	c echo.Context,
	data interface{},
) error {
	return wrapResponse(c, http.StatusOK, "OK", "", data)
}

func Error(
	c echo.Context,
	code int,
	msg string,
	desc string,
) error {
	return wrapResponse(c, code, msg, desc, nil)
}

func InternalError(
	c echo.Context,
	err error,
) error {
	return wrapResponse(c, http.StatusInternalServerError, err.Error(), "Ada kendala.", nil)
}

func wrapResponse(
	c echo.Context,
	code int,
	msg string,
	desc string,
	data interface{},
) error {
	return c.JSONPretty(code, buildResponseAPI(code, msg, desc, data), "    ")
}

func buildResponseAPI(
	status int,
	msg string,
	desc string,
	data interface{},
) *APIResponse {
	return &APIResponse{
		Status:      status,
		Message:     msg,
		Description: desc,
		Data:        data,
	}
}
