package actions

import (
	"encoding/json"
	"net/http"
	"project_pulse/mockData"
	"strings"
)

func (as *ActionSuite) Test_Task_GetTaskDetails() {
	res := as.JSON("/task/216").Get()

	as.Equal(http.StatusOK, res.Code)

	// Unmarshal the mockData
	expectedResponse, _ := json.Marshal(mockData.TaskDetailsMockData)

	// Trim the newline character from the response body
	actualResponse := strings.TrimSpace(res.Body.String())

	// Check the body content
	as.Equal(string(expectedResponse), actualResponse)

}
