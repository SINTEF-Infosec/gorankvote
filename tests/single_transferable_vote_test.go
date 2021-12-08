package gorankvote_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorankvote"
	"testing"
)

func TestSimpleSTV(t *testing.T) {
	stay := gorankvote.NewCandidate("Stay")
	soft := gorankvote.NewCandidate("Soft Brexit")
	hard := gorankvote.NewCandidate("Hard Brexit")

	candidates := []*gorankvote.Candidate{
		stay, soft, hard,
	}

	ballots := []*gorankvote.Ballot{
		gorankvote.NewBallot([]*gorankvote.Candidate{soft, stay}),
		gorankvote.NewBallot([]*gorankvote.Candidate{stay, soft}),
		gorankvote.NewBallot([]*gorankvote.Candidate{stay, soft}),
		gorankvote.NewBallot([]*gorankvote.Candidate{hard, soft}),
		gorankvote.NewBallot([]*gorankvote.Candidate{hard, stay, soft}),
	}

	electionResults, err := gorankvote.SingleTransferableVote(candidates, ballots, 1, gorankvote.DefaultSingleTransferableVoteOptions())
	if err != nil {
		t.Errorf("election failed: %v", err)
	}
	winners := electionResults.GetWinners()
	assert.Equal(t, 1, len(winners), "election should return winner")

	winner := winners[0]
	assert.Equal(t, stay, winner, "winner should be stay")
}

func TestSimpleSTV2(t *testing.T) {
	per := gorankvote.NewCandidate("Per")
	paal := gorankvote.NewCandidate("Pål")
	askeladden := gorankvote.NewCandidate("Askeladden")

	candidates := []*gorankvote.Candidate{
		per, paal, askeladden,
	}

	ballots := []*gorankvote.Ballot{
		gorankvote.NewBallot([]*gorankvote.Candidate{askeladden, per}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{paal, per}),
		gorankvote.NewBallot([]*gorankvote.Candidate{paal, per, askeladden}),
	}

	electionResults, err := gorankvote.SingleTransferableVote(candidates, ballots, 1, gorankvote.DefaultSingleTransferableVoteOptions())
	if err != nil {
		t.Errorf("election failed: %v", err)
	}
	winners := electionResults.GetWinners()
	assert.Equal(t, 1, len(winners), "election should return winner")

	winner := winners[0]
	assert.Equal(t, per, winner, "winner should be Per")
}

func TestCase1Simple(t *testing.T) {
	per := gorankvote.NewCandidate("Per")
	paal := gorankvote.NewCandidate("Pål")
	askeladden := gorankvote.NewCandidate("Askeladden")

	candidates := []*gorankvote.Candidate{
		per, paal, askeladden,
	}

	ballots := []*gorankvote.Ballot{
		gorankvote.NewBallot([]*gorankvote.Candidate{askeladden, per}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{paal, per}),
		gorankvote.NewBallot([]*gorankvote.Candidate{paal, per, askeladden}),
	}

	electionResults, err := gorankvote.SingleTransferableVote(candidates, ballots, 2, gorankvote.DefaultSingleTransferableVoteOptions())
	if err != nil {
		t.Errorf("election failed: %v", err)
	}
	winners := electionResults.GetWinners()
	assert.Equal(t, 2, len(winners), "election should return winner")

	winner1 := winners[0]
	winner2 := winners[1]

	assert.Equal(t, per, winner1, "winner1 should be Per")
	assert.Equal(t, paal, winner2, "winner2 should be Pål")
}

func TestCase2(t *testing.T) {
	per := gorankvote.NewCandidate("Per")
	paal := gorankvote.NewCandidate("Pål")
	maria := gorankvote.NewCandidate("Maria")
	ingrid := gorankvote.NewCandidate("Ingrid")

	candidates := []*gorankvote.Candidate{
		per, paal, maria, ingrid,
	}

	ballots := []*gorankvote.Ballot{
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{maria, ingrid}),
		gorankvote.NewBallot([]*gorankvote.Candidate{ingrid, maria}),
		gorankvote.NewBallot([]*gorankvote.Candidate{ingrid, maria}),
	}

	electionResults, err := gorankvote.SingleTransferableVote(candidates, ballots, 2, gorankvote.DefaultSingleTransferableVoteOptions())
	if err != nil {
		t.Errorf("election failed: %v", err)
	}
	winners := electionResults.GetWinners()
	assert.Equal(t, 2, len(winners), "election should return winner")

	winner1 := winners[0]
	winner2 := winners[1]

	assert.Equal(t, per, winner1, "winner1 should be Per")
	assert.Equal(t, paal, winner2, "winner2 should be Pål")

	votesRound := make([]float64, 0)
	for _, candidateResults := range electionResults.Rounds[0].CandidateResults {
		votesRound = append(votesRound, candidateResults.NumberOfVotes)
	}
	assertListsAlmostEqual(t, []float64{7, 2, 1, 0}, votesRound, 0.02)

	votesRound = make([]float64, 0)
	for _, candidateResults := range electionResults.Rounds[1].CandidateResults {
		votesRound = append(votesRound, candidateResults.NumberOfVotes)
	}
	assertListsAlmostEqual(t, []float64{3.33, 3.67, 1, 2}, votesRound, 0.02)
}

func TestCase3(t *testing.T) {
	per := gorankvote.NewCandidate("Per")
	paal := gorankvote.NewCandidate("Pål")
	maria := gorankvote.NewCandidate("Maria")
	ingrid := gorankvote.NewCandidate("Ingrid")

	candidates := []*gorankvote.Candidate{
		per, paal, maria, ingrid,
	}

	ballots := []*gorankvote.Ballot{
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{per, paal}),
		gorankvote.NewBallot([]*gorankvote.Candidate{maria, ingrid}),
		gorankvote.NewBallot([]*gorankvote.Candidate{ingrid, maria}),
		gorankvote.NewBallot([]*gorankvote.Candidate{ingrid, maria}),
	}

	electionResults, err := gorankvote.SingleTransferableVote(candidates, ballots, 2, gorankvote.DefaultSingleTransferableVoteOptions())
	if err != nil {
		t.Errorf("election failed: %v", err)
	}
	winners := electionResults.GetWinners()
	assert.Equal(t, 2, len(winners), "election should return winner")

	winner1 := winners[0]
	winner2 := winners[1]

	assert.Equal(t, per, winner1, "winner1 should be Per")
	assert.Equal(t, paal, winner2, "winner2 should be Pål")
}

func TestCase4(t *testing.T) {
	popularModerate := gorankvote.NewCandidate("William, popular moderate")
	moderate2 := gorankvote.NewCandidate("John, moderate")
	moderate3 := gorankvote.NewCandidate("Charles, moderate")
	farLeft := gorankvote.NewCandidate("Thomas, far-left")

	candidates := []*gorankvote.Candidate{
		popularModerate, moderate2, moderate3, farLeft,
	}

	ballots := []*gorankvote.Ballot{
		gorankvote.NewBallot([]*gorankvote.Candidate{popularModerate, moderate2, moderate3, farLeft}),
		gorankvote.NewBallot([]*gorankvote.Candidate{popularModerate, moderate2, moderate3, farLeft}),
		gorankvote.NewBallot([]*gorankvote.Candidate{popularModerate, moderate3, moderate2, farLeft}),
		gorankvote.NewBallot([]*gorankvote.Candidate{popularModerate, moderate3, moderate2, farLeft}),

		gorankvote.NewBallot([]*gorankvote.Candidate{moderate2, popularModerate, moderate3, farLeft}),
		gorankvote.NewBallot([]*gorankvote.Candidate{moderate2, popularModerate, moderate3, farLeft}),

		gorankvote.NewBallot([]*gorankvote.Candidate{farLeft, popularModerate, moderate2, moderate3}),
		gorankvote.NewBallot([]*gorankvote.Candidate{farLeft, popularModerate, moderate2, moderate3}),
		gorankvote.NewBallot([]*gorankvote.Candidate{farLeft, moderate2, popularModerate, moderate3}),
		gorankvote.NewBallot([]*gorankvote.Candidate{farLeft, moderate2, popularModerate, moderate3}),
	}

	electionResults, err := gorankvote.SingleTransferableVote(candidates, ballots, 2, gorankvote.DefaultSingleTransferableVoteOptions())
	if err != nil {
		t.Errorf("election failed: %v", err)
	}
	winners := electionResults.GetWinners()
	assert.Equal(t, 2, len(winners), "election should return 2 winner")

	winner1 := winners[0]
	winner2 := winners[1]

	assert.Equal(t, popularModerate, winner1, "winner1 should be Per")
	assert.Equal(t, farLeft, winner2, "winner2 should be Pål")
}

func assertListsAlmostEqual(t *testing.T, list1, list2 []float64, error float64) {
	assert.Equal(t, len(list1), len(list2), "lists length should be equal")

	for i := 0; i < len(list1); i++ {
		if list1[i] - list2[i] >= error {
			assert.Fail(t, fmt.Sprintf("values are too different: %f / %f", list1[i], list2[i]))
		}
	}
}
