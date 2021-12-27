package domain

//todo Можно добавить возможность делиться задачами и листами

type TodoList struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	UserId      string `json:"userId" bson:"userId,omitempty"`
	CategoryId  string `json:"categoryId" bson:"categoryId,omitempty"`
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	CreatedAt   int64  `json:"createdAt" bson:"createdAt,omitempty"`
	// Tags        []string  `json:"tags" bson:"tags"`
}

type TodoListShort struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	Title       string `json:"title" bson:"title,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	CreatedAt   int64  `json:"createdAt" bson:"createdAt,omitempty"`
	Comlited    int    `json:"comlited"`
	Count       int    `json:"count"`
	// Tags        []string  `json:"tags" bson:"tags"`
}

// todo можно потом добавить Files File

type CreateListDTO struct {
	Title       string `json:"title" binding:"required,min=3,max=128"`
	UserId      string `json:"userId"`
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
}

func NewTodoList(dto CreateListDTO) TodoList {
	return TodoList{
		Title:       dto.Title,
		UserId:      dto.UserId,
		Description: dto.Description,
		CategoryId:  dto.CategoryId,
	}
}

type UpdateListDTO struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  string `json:"categoryId"`
}

func UpdateTodoList(dto UpdateListDTO) TodoList {
	return TodoList{
		Id:          dto.Id,
		Title:       dto.Title,
		Description: dto.Description,
		CategoryId:  dto.CategoryId,
	}
}
