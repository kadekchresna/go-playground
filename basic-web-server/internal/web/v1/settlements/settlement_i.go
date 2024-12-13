package settlements

import "github.com/labstack/echo/v4"

type SettlementsHandler interface {
	GetSettlements(c echo.Context) error
	GetOrders(c echo.Context) error
	GetLogs(c echo.Context) error
	GetLogsChan(c echo.Context) error
	GetLogsPaginate(c echo.Context) error
}
