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
	Vantagens float64
}

type Grupos struct {
	Nome   string
	G1, G2 []string
}

func gruposDefinidos() {}

func jogosSimulados(casa, fora Time, teamStats map[Time]int) {
	golsCasa := rand.Intn(5) // Gols do time da casa (0 a 4 gols)
	golsFora := rand.Intn(5) // Gols do time visitante (0 a 4 gols)

	if golsCasa > golsFora {
		teamStats[casa] += 3
	} else if golsCasa < golsFora {
		teamStats[fora] += 3
	} else {
		teamStats[casa]++
		teamStats[fora]++
	}
}

func campeonatoSimulado(vantagenss []VantagensTimes) map[Time]int {
	teamStats := make(map[Time]int)
	for _, team := range vantagenss {
		teamStats[team.Time] = 0
	}

	for i := 0; i < len(vantagenss)-1; i++ {
		for j := i + 1; j < len(vantagenss); j++ {
			jogosSimulados(vantagenss[i].Time, vantagenss[j].Time, teamStats)
		}
	}

	return teamStats
}

func main() {
	rand.Seed(time.Now().UnixNano())

	vantagens := []VantagensTimes{
		{Time: Flamengo, Vantagens: 1.2},
		{Time: Fluminense, Vantagens: 1.0},
		{Time: Palmeiras, Vantagens: 1.1},
		{Time: River, Vantagens: 0.9},
		{Time: Internacional, Vantagens: 0.8},
		{Time: Atletico, Vantagens: 1.0},
	}

	teamStats := campeonatoSimulado(vantagens)

	fmt.Println("Pontuação final do campeonato:")
	for team, points := range teamStats {
		fmt.Printf("%s: %d pontos\n", team, points)
	}
}
