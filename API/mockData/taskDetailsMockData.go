package mockData

import "project_pulse/models"

var TaskDetailsMockData = models.TaskDetailsModel{
	ID:          1,
	Title:       "Test Task",
	Description: "This is a test task",
	ProjectName: "Test Project",
	Board:       1,
	Assignee:    "Test Assignee",
	Reporter:    "Test Reporter",
	StoryPoints: 5,
	Priority:    "High",
	Status:      "In Progress",
	Comments: []models.CommentModel{
		{
			ID:        1,
			Author:    "Test Author",
			Text:      "Test comment",
			TimeStamp: "2020-04-20 12:00:00",
		},
	},
}
