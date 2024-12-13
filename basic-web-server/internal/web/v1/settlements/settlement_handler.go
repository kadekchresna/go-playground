package settlements

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": s})
}

func (h *settlementsHandler) GetOrders(c echo.Context) error {
	ctx := c.Request().Context()
	orders, err := h.SettlementsUsecase.GetOrders(ctx, settlements.GetAllOrdersParams{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": orders})
}

func (h *settlementsHandler) GetLogs(c echo.Context) error {
	ctx := c.Request().Context()
	logs, err := h.SettlementsUsecase.GetLogs(ctx, settlements.GetAllOrdersParams{})
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": logs})
}

func (h *settlementsHandler) GetLogsChan(c echo.Context) error {
	ctx := c.Request().Context()
	var params settlements.GetAllOrdersParams
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	logsChan, err := h.SettlementsUsecase.GetLogsChan(ctx, params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	c.Response().Header().Set(echo.HeaderContentType, "text/csv")
	c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment;filename=log-%d.csv", time.Now().Unix()))
	writer := c.Response().Writer
	for log := range logsChan {
		content, _ := json.Marshal(log)
		if _, err := writer.Write(content); err != nil {
			return err
		}

		c.Response().Flush()
	}

	return nil
}

func (h *settlementsHandler) GetLogsPaginate(c echo.Context) error {
	ctx := c.Request().Context()

	var params settlements.GetAllOrdersParams
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	logs, err := h.SettlementsUsecase.GetLogsPaginate(ctx, params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": logs})
}
