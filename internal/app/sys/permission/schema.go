package permission

import "jobor/internal/models"

var tbName = models.Permission

type ShowData struct {
	ID          uint   `json:"id"`
	Name        string `json:"name" form:"name"`
	Url         string `json:"url"`
	Method      string `json:"Method"`
}
func (ShowData) TableName() string {
	return tbName
}

type PostSchema struct {
	Name        string `json:"name" binding:"required"`
	Url         string `json:"url"`
	Method      string `json:"Method"`
	//ByUpdate    string `json:"by_update,-"`
}
func (PostSchema) TableName() string {
	return tbName
}

type PutSchema struct {
	Name        *string `json:"name,omitempty"`
	Url         *string `json:"url,omitempty"`
	Method      *string `json:"method,omitempty"`
	//ByUpdate    *string `json:"by_update,-"`
}

func (PutSchema) TableName() string {
	return tbName
}

type DeleteSchema struct {
	Rows []uint `json:"rows"`
}