package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Definição dos times
type Time string

const (
	Flamengo      Time = "Flamengo"
	Fluminense    Time = "Fluminense"
	Palmeiras     Time = "Palmeiras"
	River         Time = "River"
	Internacional Time = "Internacional"
	Atletico      Time = "Atletico"
)

// Estrutura para guardar a vantagem de cada time
type VantagensTimes struct {
	Time      Time
	Advantage float64
}

// Estrutura que representa um time na tabela do campeonato
type DadosTimes struct {
	ID           int
	Time         Time
	Points       int
	GoalsFor     int
	GoalsAgainst int
}

// Função para simular o resultado de um jogo
func jogosSimulados(casa, fora Time, dadosTimes map[Time]int) {
	// Lógica para simular o jogo entre times casa e fora
	// Atualize os dados de dadosTimes com o resultado
}

// Função para simular o campeonato
func simulateChampionship(vantagens []VantagensTimes) map[Time]int {
	dadosTimes := make(map[Time]int)
	for _, time := range vantagens {
		dadosTimes[time.Time] = 0 // Inicializa os pontos de cada time
	}

	// Lógica para simular os jogos do campeonato
	// Utilize a função simulateGame e atualize dadosTimes com os resultados dos jogos

	return dadosTimes
}

func main() {
	// Seed para gerar números aleatórios
	rand.Seed(time.Now().UnixNano())

	// Defina a vantagem de cada time
	vantagens := []VantagensTimes{
		{Time: Flamengo, Advantage: 1.2},
		{Time: Fluminense, Advantage: 1.0},
		// Adicione os outros times com suas vantagens
	}

	// Simula o campeonato e obtém a pontuação final de cada time
	dadosTimes := simulateChampionship(vantagens)

	// Exibe a pontuação final de cada time
	fmt.Println("Pontuação final do campeonato:")
	for time, points := range dadosTimes {
		fmt.Printf("%s: %d pontos\n", time, points)
	}
}
