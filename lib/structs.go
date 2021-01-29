package lib

import "sync"

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

func NewVotingMap() (r *VoteStruct)  {
	return &VoteStruct{voteList: make(map[string]int)}
}
