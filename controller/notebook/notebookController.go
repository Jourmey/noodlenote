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
	note.GetNotesIDByBookID(bookId)

	n := make([]string, 0, 0) //NoteBook.Path
	c.JSON(utils.SUCCESS, n)
}

//notebook
//笔记本路径
//是
//String
func Create(c *gin.Context) {
	book := new(model.NoteBook)
	err := c.ShouldBind(&book)
	book.Create()
	if err != nil {
		c.JSON(utils.SUCCESS, utils.CreateNoteBookSUCCESS)
	} else {
		c.JSON(utils.ERROR, utils.CreateNoteBookERROR)
	}
}

func Delete(c *gin.Context) {
	bookId := utils.Int64Parse(c.Param("id"))

	note := new(model.NoteBookLink)
	//links := note.GetNotesIDByBookID(bookId)
	note.DeleteByBookID(bookId)

	book := new(model.NoteBook)
	book.ID = bookId
	book.Delete()

	c.JSON(utils.SUCCESS, utils.CreateNoteBookSUCCESS)
}
