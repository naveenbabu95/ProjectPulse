package mockData

import "project_pulse/models"

var SprintBoardMockData = []models.SprintBoardModel{
	{
		ID:          1,
		Name:        "Task 1",
		Description: "A sprint board is a tool for planning and tracking work during a project. Sprint boards are usually set up in a team room or other open space where the team and stakeholders can see them.",
		ProjectName: "Sprint Board",
		Assignee:    "Ravi",
		Reporter:    "Alice",
		StoryPoints: 5,
		Status:      "IN_PROGRESS",
	},
	{
		ID:          2,
		Name:        "Task 2",
		Description: "A sprint board is a tool for planning and tracking work during a project. Sprint boards are usually set up in a team room or other open space where the team and stakeholders can see them.",
		ProjectName: "Sprint Board",
		Assignee:    "Bob",
		Reporter:    "Becky",
		StoryPoints: 8,
		Status:      "DRAFT",
	},
}
