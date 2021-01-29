package main

import (
	"github.com/khaosdoctor/go-vote-api/lib"
	"github.com/labstack/echo/v4"
	"net/http"
)

var votes *lib.VoteStruct

func main() {
	votes = lib.NewVotingMap()
	e := echo.New()
	e.GET("/votes/:id", voteHandler)
	e.GET("/votes", listVotesHandler)
	e.GET("/total", totalHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func voteHandler (c echo.Context) error {
	voteId := c.Param("id")
	votes.AddVote(voteId)
	return c.JSON(http.StatusOK, votes.GetVotes())
}

func totalHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, votes.GetTotalVotes())
}

func listVotesHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, votes.GetVotes())
}
