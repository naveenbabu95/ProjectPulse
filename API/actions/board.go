package actions

import (
	"fmt"
	"net/http"
	"project_pulse/mockData"

	"github.com/gobuffalo/buffalo"
)

func GetBoardHandler(c buffalo.Context) error {
	fmt.Println(c.Param("boardId")) // use boardID to fetch board details from db

	return c.Render(http.StatusOK, r.JSON(mockData.SprintBoardMockData))

}
