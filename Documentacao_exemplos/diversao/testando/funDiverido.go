package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Time string

const (
	Flamengo      Time = "Flamengo"
	Fluminense    Time = "Fluminense"
	Palmeiras     Time = "Palmeiras"
	River         Time = "River"
	Internacional Time = "Internacional"
	Atletico      Time = "Atletico"
)

type VantagensTimes struct {
	Time      Time
	Advantage float64
}

func simulateGame(home, away Time, teamStats map[Time]int) {
	homeGoals := rand.Intn(5) // Gols do time da casa (0 a 4 gols)
	awayGoals := rand.Intn(5) // Gols do time visitante (0 a 4 gols)

	if homeGoals > awayGoals {
		teamStats[home] += 3
	} else if homeGoals < awayGoals {
		teamStats[away] += 3
	} else {
		teamStats[home]++
		teamStats[away]++
	}
}

func simulateChampionship(advantages []VantagensTimes) map[Time]int {
	teamStats := make(map[Time]int)
	for _, team := range advantages {
		teamStats[team.Time] = 0
	}

	for i := 0; i < len(advantages)-1; i++ {
		for j := i + 1; j < len(advantages); j++ {
			simulateGame(advantages[i].Time, advantages[j].Time, teamStats)
		}
	}

	return teamStats
}

func main() {
	rand.Seed(time.Now().UnixNano())

	vantagens := []VantagensTimes{
		{Time: Flamengo, Advantage: 1.2},
		{Time: Fluminense, Advantage: 1.0},
		{Time: Palmeiras, Advantage: 1.1},
		{Time: River, Advantage: 0.9},
		{Time: Internacional, Advantage: 0.8},
		{Time: Atletico, Advantage: 1.0},
	}

	teamStats := simulateChampionship(vantagens)

	fmt.Println("Pontuação final do campeonato:")
	for team, points := range teamStats {
		fmt.Printf("%s: %d pontos\n", team, points)
	}
}
