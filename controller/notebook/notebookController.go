package notebookController

import (
	"github.com/gin-gonic/gin"
	"noodlenote/model"
	"noodlenote/utils"
)

func All(c *gin.Context) {
	books := model.NoteBook{}.GetAll()
	c.JSON(utils.SUCCESS, books)
}

func List(c *gin.Context) {
	bookId := utils.Int64Parse(c.Param("id"))
	note := new(model.NoteBookLink)
	noteIds := note.GetNotesIDByBookID(bookId)
	c.JSON(utils.SUCCESS, noteIds)
}

//notebook
//笔记本路径
//是
//String
func Create(c *gin.Context) {
	book := new(model.NoteBook)
	err := c.ShouldBind(&book)
	if err != nil {
		book.Create()
		c.JSON(utils.SUCCESS, gin.H{"msg": utils.CreateNoteBookSuccess, "note_book_id": book.ID})

	} else {
		c.JSON(utils.SUCCESS, utils.CreateNoteBookErr)
	}
}

func Delete(c *gin.Context) {
	bookId := utils.Int64Parse(c.Param("id"))

	//删笔记本不删关系
	//note := new(model.NoteBookLink)
	//note.DeleteByBookID(bookId)

	book := new(model.NoteBook)
	book.ID = bookId
	book.Delete()

	c.JSON(utils.SUCCESS, utils.CreateNoteBookSuccess)
}
