package domain

type TodoItem struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	UserId      string `json:"userId" bson:"userId,omitempty"`
	ListId      string `json:"listId" bson:"listId,omitempty"`
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	CreatedAt   int64  `json:"createdAt" bson:"createdAt,omitempty"`
	CompletedAt int64  `json:"completedAt" bson:"completedAt,omitempty"`
	StartAt     int64  `json:"startAt" bson:"startAt,omitempty"`
	Done        bool   `json:"done" bson:"done,omitempty"`
	Priority    int    `json:"priority" bson:"priority,omitempty"`
}

type CreateTodoDTO struct {
	ListId      string `json:"listId" binding:"required"`
	Title       string `json:"title" binding:"required,min=3,max=128"`
	Description string `json:"description"`
	StartAt     int64  `json:"startAt"`
	Priority    int    `json:"priority"`
}

func NewTodo(dto CreateTodoDTO) TodoItem {
	return TodoItem{
		ListId:      dto.ListId,
		Title:       dto.Title,
		Description: dto.Description,
		StartAt:     dto.StartAt,
		Priority:    dto.Priority,
	}
}

type UpdateTodoDTO struct {
	Id          string `json:"id"`
	ListId      string `json:"listId" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartAt     int64  `json:"startAt"`
	Priority    int    `json:"priority"`
	Done        bool   `json:"done"`
}

func UpdateTodo(dto UpdateTodoDTO) TodoItem {
	return TodoItem{
		Id:          dto.Id,
		ListId:      dto.ListId,
		Title:       dto.Title,
		Description: dto.Description,
		StartAt:     dto.StartAt,
		Priority:    dto.Priority,
		Done:        dto.Done,
	}
}
