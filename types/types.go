package types

//Task is the struct used to identify tasks
type Task struct {
	ID      int
	Title   string
	Content string
	Created string
}

//Context is the struct passed to templates
type Context struct {
	Tasks      []Task
	Navigation string
	Search     string
	Message    string
}
