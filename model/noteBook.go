package model

import "time"

type NoteBook struct {
	BaseModel
	Name     string `gorm:"column:name" json:"name"`           // “笔记本1”, // 笔记本的名称
	NotesNum int16  `gorm:"column:notes_num" json:"notes_num"` // “3” // 该笔记本中笔记的数目
}

func (b *NoteBook) Create() {
	b.CreatedAt = time.Now()
	db.Create(b)
}

func (b *NoteBook) Delete() {
	db.Delete(b)
}

func (b NoteBook) GetAll() (books []NoteBook) {
	db.Find(&books)
	return books
}
