package settlements

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

type JobSpec struct {
	APIVersion string       `json:"apiVersion"`
	Kind       string       `json:"kind"`
	Metadata   Metadata     `json:"metadata"`
	Spec       SpecTemplate `json:"spec"`
}

type Metadata struct {
	Name string `json:"name"`
}

type Container struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Command []string `json:"command"`
}

type Spec struct {
	Containers    []Container `json:"containers"`
	RestartPolicy string      `json:"restartPolicy"`
}

type Template struct {
	Spec Spec `json:"spec"`
}

type SpecTemplate struct {
	Template `json:"template"`
}

func createJob(ctx context.Context) error {

	job := JobSpec{
		APIVersion: "batch/v1",
		Kind:       "Job",
		Metadata: Metadata{
			Name: "example-job-2",
		},
		Spec: SpecTemplate{
			Template: Template{
				Spec: Spec{
					Containers: []Container{
						{
							Name:    "example-container",
							Image:   "busybox",
							Command: []string{"echo", "Hello, Kubernetes!"},
						},
					},
					RestartPolicy: "Never",
				},
			},
		},
	}

	jobBytes, err := json.Marshal(job)
	if err != nil {
		return fmt.Errorf("failed to marshal job: %v", err)
	}

	// <kubernetes-api-server> => minikube ip
	// <your-access-token>
	req, err := http.NewRequestWithContext(ctx, "POST", "http://localhost:8001/apis/batch/v1/namespaces/default/jobs", bytes.NewBuffer(jobBytes))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6InBFa1M0OTVIeWtuOWZCSWNNdEpFbU16R1ozREg4S05nWURYX2R6NE8taTgifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6Im15LXNlcnZpY2UtYWNjb3VudC10b2tlbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJteS1zZXJ2aWNlLWFjY291bnQiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiI3MDc5YmI0Mi02NDdkLTQzOTktOTZiNS0zMjRlNjgwZGRkOGUiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6ZGVmYXVsdDpteS1zZXJ2aWNlLWFjY291bnQifQ.TlG9ZOJg2MweTBpRk82_FiZ-VsOUQYKIiZn8dZSwTKhNqnDhppgitooPtY3fhxYrFXwkxivh8irryz3SHwbvpHcwrO1wY55KkURYrQ3wtosFl8bjRf8VsK3rMVJzoyov635Hy3Zn7uyVcRefDo1j2h03E22jy0kbtWJlp9R8sZPJCBz9AQe8-IgLv-_F10NxwM_pOZxUUVRdCTEqg5eDJ_5WBIunoCtYT-kGDgnBPkjBUnGHy_P8tQLlUF1gizkBLcjBt654K0R_N0gH7faDAjAbkUJWE1CllJvcaroU0nLNmhCSRX6WlTOXf43kl7zyN3pVFhfkdy8qTDud-siHsg")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create job, status code: %d", resp.StatusCode)
	}

	fmt.Println("Job created successfully")
	return nil
}

func (h *settlementsHandler) CreateJob(c echo.Context) error {
	ctx := c.Request().Context()

	if err := createJob(ctx); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})

	}

	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "message": "success", "data": nil})
}
