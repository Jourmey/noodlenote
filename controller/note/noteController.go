package noteController

import (
	"github.com/gin-gonic/gin"
	"noodlenote/model"
	"noodlenote/utils"
)

func Create(c *gin.Context) {
	bookId := utils.Int64Parse(c.DefaultQuery("id", "0"))
	note := new(model.Note)
	err := c.ShouldBind(&note)
	if err == nil && note != nil {
		note.Created()

		link := new(model.NoteBookLink)
		link.NoteBookID = bookId
		link.NoteID = note.ID
		link.Create()

		c.JSON(utils.SUCCESS, gin.H{"msg": utils.CreateNoteSuccess, "note_id": note.ID})
	} else {
		c.JSON(utils.SUCCESS, utils.CreateNoteErr)
	}
}
func Get(c *gin.Context) {
	noteId := utils.Int64Parse(c.Param("id"))
	note := new(model.Note)
	note.GetByID(noteId)
	c.JSON(utils.SUCCESS, note)
}

func Update(c *gin.Context) {
	note := new(model.Note)
	err := c.ShouldBind(&note)
	if err != nil {
		c.JSON(utils.SUCCESS, utils.CreateNoteErr)
		return
	}
	note.Update()
	c.JSON(utils.SUCCESS, utils.CreateNoteSuccess)
}

func Move(c *gin.Context) {
	noteId := utils.Int64Parse(c.Param("id"))
	bookId := utils.Int64Parse(c.DefaultQuery("id", "0"))
	if noteId == bookId {
		c.JSON(utils.SUCCESS, utils.MoveNoteErr)
		return
	}

	noteLink := new(model.NoteBookLink)
	noteLink.GetByNoteID(noteId)
	noteLink.NoteBookID = bookId
	noteLink.Update()

	c.JSON(utils.SUCCESS, utils.MoveNoteSuccess)
}
func Delete(c *gin.Context) {
	noteId := utils.Int64Parse(c.Param("id"))
	noteLink := new(model.NoteBookLink)
	noteLink.GetByNoteID(noteId)
	note := new(model.Note)
	note.GetByID(noteId)

	note.Delete()
	noteLink.Delete()
	c.JSON(utils.SUCCESS, utils.DeleteNoteSuccess)
}
