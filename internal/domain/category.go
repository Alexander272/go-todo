package domain

type Category struct {
	Id     string `json:"id" bson:"_id,omitempty"`
	UserId string `json:"userId" bson:"userId,omitempty"`
	Title  string `json:"title" bson:"title,omitempty"`
}

type CategoryWithLists struct {
	Id    string          `json:"id" bson:"_id,omitempty"`
	Title string          `json:"title" bson:"title,omitempty"`
	Lists []TodoListShort `json:"lists" bson:"lists"`
}

type CreateCategoryDTO struct {
	UserId string `json:"userId"`
	Title  string `json:"title" binding:"required"`
}

func NewCategory(dto CreateCategoryDTO) Category {
	return Category{
		UserId: dto.UserId,
		Title:  dto.Title,
	}
}

type UpdateCategoryDTO struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Title string `json:"title" binding:"required"`
}

func UpdateCategory(dto UpdateCategoryDTO) Category {
	return Category{
		Id:    dto.Id,
		Title: dto.Title,
	}
}
