package model

type Note struct {
	BaseModel
	Title      string `gorm:"column:title" json:"title"`               //“title” : “工作记录”, // 笔记标题
	Author     string `gorm:"column:author" json:"author"`             //“author” : “Tom”, // 笔记作者
	Size       string `gorm:"column:size" json:"size"`                 //“size” : “1024”, // 笔记大小，包括笔记正文及附件
	Content    string `gorm:"column:content" json:"content"`           //“content” : “<p>This is a test note</p> // 笔记正文
}
