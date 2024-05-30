package dto

import "idleRain.com/ginEssential/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Id        uint   `json:"id"`
}

// ToUserDto 封装用户信息的 dto
func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:        user.ID,
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
