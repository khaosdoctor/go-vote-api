package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type VoteStruct struct {
	mut      sync.RWMutex
	voteList map[string]int
}

type returnType struct{
	Total int `json:"total"`
	Votes map[string]int `json:"voteList"`
}

func (s *VoteStruct) AddVote (id string) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.voteList[id]++
}

func (s *VoteStruct) GetVotes () map[string]int {
	s.mut.RLock()
	defer s.mut.RUnlock()
	return s.voteList
}

func (s *VoteStruct) GetTotalVotes() returnType {
	s.mut.RLock()
	defer s.mut.RUnlock()
	finalAcc := returnType{
		Total: 0,
		Votes: s.GetVotes(),
	}
	for _, totalVotes := range s.GetVotes() {
		finalAcc.Total += totalVotes
	}
	return finalAcc
}

var votes VoteStruct

func main() {
	votes = VoteStruct{voteList: make(map[string]int)}
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
