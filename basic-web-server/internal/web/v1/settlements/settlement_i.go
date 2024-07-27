package settlements

import "github.com/labstack/echo/v4"

type SettlementsHandler interface {
	GetSettlements(c echo.Context) error

	GetOrders(c echo.Context) error
}
