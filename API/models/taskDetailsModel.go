package models

type TaskDetailsModel struct {
	ID          int
	Title       string
	Description string
	ProjectName string
	Board       int
	Assignee    string
	Reporter    string
	StoryPoints int
	Priority    string
	Status      string
	Comments    []CommentModel
}
