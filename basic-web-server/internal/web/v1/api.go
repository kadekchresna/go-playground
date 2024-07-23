package v1

import (
	"basic.web/internal/web/v1/settlements"
	"github.com/labstack/echo/v4"
)

func InitAPI(
	e *echo.Echo,
	settlements settlements.SettlementsHandler,
) {

	v1 := e.Group("/api/v1")

	v1.GET("/settlement", settlements.GetSettlements)

}
