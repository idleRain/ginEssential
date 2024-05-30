package model

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Name      string `json:"name" gorm:"type:varchar(50);not null"`
	CreatedAt Time   `json:"createdAt" gorm:"type:timestamp"`
	UpdatedAt Time   `json:"updateAd" gorm:"type:timestamp"`
}

// Category 文章
//func Category() {
//
//}
