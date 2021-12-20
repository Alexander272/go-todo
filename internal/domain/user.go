package domain

type User struct {
	Id           string       `json:"id" bson:"_id,omitempty"`
	UserUrl      string       `json:"userUrl" bson:"userUrl"`
	Name         string       `json:"name" bson:"name,omitempty"`
	Email        string       `json:"email" bson:"email"`
	Password     string       `json:"password" bson:"password"`
	Role         string       `json:"role" bson:"role"`
	RegisteredAt int64        `json:"registeredAt" bson:"registeredAt"`
	LastVisitAt  int64        `json:"-" bson:"lastVisitAt"`
	Verification Verification `json:"-" bson:"verification"`
	// todo можно потом добавить настройки для пользователя
}

type Verification struct {
	Code     string `json:"code" bson:"code"`
	Verified bool   `json:"verified" bson:"verified"`
	Expires  int64  `json:"expires" bson:"expires"`
}

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

func NewUser(dto CreateUserDTO) User {
	return User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

type UpdateUserDTO struct {
	UserId   string
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	UserUrl  string `form:"userUrl" json:"userUrl"`
}

func UpdateUser(dto UpdateUserDTO) User {
	return User{
		Id:       dto.UserId,
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		UserUrl:  dto.UserUrl,
	}
}
