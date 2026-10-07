package deliveryparam

type Request struct{
	Command string
	CreateTaskRequest CreateTaskRequest 
	CreateCategoryRequest CreateCategoryRequest
}

type CreateTaskRequest struct{
	Title string
	DueDate string
	CategoryID int
}

type CreateCategoryRequest struct{
	Title  string
	Color  string
}