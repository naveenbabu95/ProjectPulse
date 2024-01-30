package models

type SprintBoardModel struct {
	ID          int
	Name        string
	Description string
	ProjectName string
	Assignee    string
	Reporter    string
	StoryPoints int
	Status      string
}
