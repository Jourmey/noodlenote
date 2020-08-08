package model

type NoteBookLink struct {
	BaseModel
	NoteBookID int64 `gorm:"column:note_book_id" json:"note_book_id"` // 对应笔记的ID
	NoteID     int64 `gorm:"column:note_id" json:"note_id"`           // 对应笔记的ID
}

func (l NoteBookLink) GetNotesIDByBookID(bookID int64) []int64 {
	roots := make([]NoteBookLink, 0)
	db.Where("note_book_id=?", bookID).Select("note_id").Find(&roots)

	noteIDs := make([]int64, 0, len(roots))
	for _, root := range roots {
		noteIDs = append(noteIDs, root.NoteID)
	}

	return noteIDs
}

func (l NoteBookLink) DeleteByBookID(bookID int64) {
	db.Where("note_book_id=?", bookID).Delete(l)
}

func (l *NoteBookLink) Create() {
	db.Create(l)
}

func (l *NoteBookLink) GetByNoteID(noteId int64) {
	db.First(l, "note_id=?", noteId)
}

func (l *NoteBookLink) Update() {
	db.Save(l)
}

func (l *NoteBookLink) Delete() {
	db.Delete(l)
}
