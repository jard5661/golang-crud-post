package structs

import "time"

//Post
type Post struct{
	Id 				int 			`json:"-"`
	Title			string			`json:"title" gorm:"size:200"`
	Content			string			`json:"content" gorm:"type:text"`
	Category		string			`json:"category" gorm:"size:100"`
	Created_date	time.Time		`json:"-" gorm:"type:timestamp;default:current_timestamp"`
	Updated_date	time.Time		`json:"-" gorm:"type:timestamp null	"`
	Status 			string			`json:"status" gorm:"size:100"`
}

// Result is an array of post
type Result struct {
	Code 	int 		`json:"code"`
	Data 	interface{} `json:"data,omitempty"`
	Message string 		`json:"message"`
}