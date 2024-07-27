package settlements

import (
	"net/http"

	"basic.web/internal/usecase/settlements"
	"github.com/labstack/echo/v4"
)

type settlementsHandler struct {
	SettlementsUsecase settlements.SettlementsUsecase
}

func NewSettlementsHandler(SettlementsUsecase settlements.SettlementsUsecase) SettlementsHandler {
	return &settlementsHandler{
		SettlementsUsecase: SettlementsUsecase,
	}
}

func (h *settlementsHandler) GetSettlements(c echo.Context) error {
	ctx := c.Request().Context()
	s, err := h.SettlementsUsecase.GetSettlements(ctx, settlements.GetSettlementsParams{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": GetSettlementsResponse{ID: s.ID, Code: s.Code, TotalAmount: s.TotalAmount, CreatedAt: s.CreatedAt}})
}

func (h *settlementsHandler) GetOrders(c echo.Context) error {
	ctx := c.Request().Context()
	orders, err := h.SettlementsUsecase.GetOrders(ctx, settlements.GetAllOrdersParams{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": orders})
}
