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
	G1, G2 []Time
}

func sorteioGrupos(times []Time) Grupos {
	rand.Seed(time.Now().UnixNano())

	rand.Shuffle(len(times), func(i, j int) {
		times[i], times[j] = times[j], times[i]
	})

	meio := len(times) / 2

	return Grupos{
		Nome: "Grupos A e B",
		G1:   times[:meio],
		G2:   times[meio:],
	}
}

func jogosGrupo(grupo []Time, dadosTime map[Time]int) {
	for i := 0; i < len(grupo)-1; i++ {
		for j := i + 1; j < len(grupo); j++ {
			jogosSimulados(grupo[i], grupo[j], dadosTime)
		}
	}
}

func jogosSimulados(casa, fora Time, dadosTime map[Time]int) {
	golsCasa := rand.Intn(5) // Gols do time da casa (0 a 4 gols)
	golsFora := rand.Intn(5) // Gols do time visitante (0 a 4 gols)

	if golsCasa > golsFora {
		dadosTime[casa] += 3
	} else if golsCasa < golsFora {
		dadosTime[fora] += 3
	} else {
		dadosTime[casa]++
		dadosTime[fora]++
	}
}

func campeonatoSimulado(vantagenss []VantagensTimes) map[Time]int {
	dadosTime := make(map[Time]int)
	for _, time := range vantagenss {
		dadosTime[time.Time] = 0
	}

	return dadosTime
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

	times := make([]Time, len(vantagens))
	for i, v := range vantagens {
		times[i] = v.Time
	}

	grupos := sorteioGrupos(times)

	dadosTime := campeonatoSimulado(vantagens)

	jogosGrupo(grupos.G1, dadosTime)
	jogosGrupo(grupos.G2, dadosTime)

	fmt.Printf("%s\n", grupos.G1)
	//Arruma separar organizar
	fmt.Println("Grupo A", grupos.G1, dadosTime)
	fmt.Println("Grupo B", grupos.G2, dadosTime)

	fmt.Println("Pontuação final do campeonato:")
	for time, points := range dadosTime {
		fmt.Printf("%s:%d ", time, points)
		fmt.Printf("%s: %d pontos\n", time, points)
	}
}
