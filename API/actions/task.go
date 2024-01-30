package actions

import (
	"fmt"
	"net/http"
	"project_pulse/mockData"

	"github.com/gobuffalo/buffalo"
)

func GetTaskDetails(c buffalo.Context) error {
	fmt.Println(c.Param("taskId")) // use taskId to fetch board details from db

	return c.Render(http.StatusOK, r.JSON(mockData.TaskDetailsMockData))

}
