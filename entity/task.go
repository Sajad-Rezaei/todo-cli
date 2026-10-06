package entity


type Task struct {
	Id         int
	Title      string
	DueDate    string
	CategoryID int
	IsDone     bool
	UserID     int
}
