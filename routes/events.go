package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"project.com/API/models"
)

func getEvent(context *gin.Context){
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
		return
	}

	event, err := models.GetEventById(eventid)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, event	)
}

func getEvents(context *gin.Context){
	events, err := models.GetAllEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
} 

func createEvent(context *gin.Context){
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
		return
	}
	userId := context.GetInt64("userID")
	event.UserID = userId
	err = event.Save()
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message":"Event created!", "event": event})
}

func updateEvent(context *gin.Context){
	eventid, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
		return
	}
	userId := context.GetInt64("userID")
	event, err := models.GetEventById(eventid)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fech the event."})
		return
	}

	if event.UserID != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message":"Not authorized to update event."})
		return
	}

	var updatedEvent models.Event 
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request data."})
		return 
	}

	updatedEvent.ID = eventid
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not update event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"Event update successfully!"})
}

func deletEvent(context *gin.Context){
	eventId , err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse event id."})
	}

	userId := context.GetInt64("userID")
	event, err := models.GetEventById(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fech the event."})
		return
	}
	
	if event.UserID != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message":"Not authorized to delete event."})
		return
	}
	
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not delete the event."})
	}

	context.JSON(http.StatusOK, gin.H{"message":"event deleted successfully!"})


}