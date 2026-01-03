package main

import (
	"io"
	"os"
	"slices"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, v := range results {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func (l *League) MatchResult(firstTeamName string, scoreFirstTeam int, secondTeamName string, scoreSecondTeam int) {
	// Security to check that both teams exists
	if _, ok := l.Teams[firstTeamName]; !ok {
		return
	}
	if _, ok := l.Teams[secondTeamName]; !ok {
		return
	}
	if scoreFirstTeam == scoreSecondTeam {
		return
	}
	if scoreFirstTeam > scoreSecondTeam {
		l.Wins[firstTeamName] += 1
		// or
		// l.Wins[firstTeamName]++
	} else {
		l.Wins[secondTeamName] += 1
	}
}

func (league League) Ranking() []string {
	// It creates an slice from Wins keys
	ranking := make([]string, 0, len(league.Wins))
	for k := range league.Wins {
		ranking = append(ranking, k)
	}

	// Sorts the keys using the wins for each team
	slices.SortFunc(ranking, func(a, b string) int {
		valA := league.Wins[a]
		valB := league.Wins[b]

		// Sort descendent
		if valB != valA {
			return valB - valA
		}

		// In case vals are equal, sort alphabetically
		if a < b {
			return -1
		}
		return 1
	})

	// Book solution
	// Being names the keys slice
	// sort is an imported library instead of slices
	// sort.Slice(names, func(i, j int) bool {
	// 	return l.Wins[names[i]] > l.Wins[names[j]]
	// })

	return ranking
}

func main() {
	firstTeam := Team{
		Name:    "Real Madrid",
		Players: []string{"Vinicius", "Rodrygo", "Mbappe"},
	}

	secondTeam := Team{
		Name:    "Chelsea",
		Players: []string{"Drogba", "Lampard", "Terry"},
	}

	// Important to start every map with an
	// empty literal map instead of a nil map
	// With the nil map, the lines adding the teams
	// panics.
	league := League{
		Teams: map[string]Team{},
		Wins:  map[string]int{},
	}

	league.Teams[firstTeam.Name] = firstTeam
	league.Teams[secondTeam.Name] = secondTeam

	league.Wins[firstTeam.Name] = 1
	league.Wins[secondTeam.Name] = 0

	league.MatchResult("Real Madrid", 1, "Chelsea", 0)
	RankPrinter(league, os.Stdout)
}
