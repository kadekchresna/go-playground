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

	// OOM using data dummy
	// https://github.com/morenoh149/postgresDBSamples/blob/master/dellstore2-normal-1.0/dellstore2-normal-1.0.sql
	v1.GET("/order", settlements.GetOrders)
	v1.GET("/log", settlements.GetLogs)
	v1.GET("/log/chunk", settlements.GetLogsChan)
	v1.GET("/log/paginate", settlements.GetLogsPaginate)

}
