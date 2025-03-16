package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MAXIMO int = 200

var matriz [MAXIMO][MAXIMO]float64
var termosIndependentes [MAXIMO]float64
var solucoes [MAXIMO]float64
var tamanho int

func gerarMatriz() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			matriz[i][j] = rand.Float64() * 100
		}
		termosIndependentes[i] = rand.Float64() * 100
	}
}

func eliminacaoGaussiana() {
	for i := 0; i < tamanho; i++ {
		pivo := matriz[i][i]
		if math.Abs(pivo) < 1e-9 {
			fmt.Println("Erro: pivô próximo de zero.")
			return
		}
		for j := i; j < tamanho; j++ {
			matriz[i][j] /= pivo
		}
		termosIndependentes[i] /= pivo

		for k := i + 1; k < tamanho; k++ {
			fator := matriz[k][i]
			for j := i; j < tamanho; j++ {
				matriz[k][j] -= fator * matriz[i][j]
			}
			termosIndependentes[k] -= fator * termosIndependentes[i]
		}
	}

	for i := tamanho - 1; i >= 0; i-- {
		solucoes[i] = termosIndependentes[i]
		for j := i + 1; j < tamanho; j++ {
			solucoes[i] -= matriz[i][j] * solucoes[j]
		}
	}
}

func imprimirMatriz() {
	fmt.Println("Matriz aumentada:")
	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			fmt.Printf("%.2f ", matriz[i][j])
		}
		fmt.Printf("| %.2f\n", termosIndependentes[i])
	}
	fmt.Println()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: go run gauss.go <tamanho da matriz>")
		return
	}

	tamanhoConvertido, err := strconv.Atoi(os.Args[1])
	if err != nil || tamanhoConvertido <= 0 || tamanhoConvertido > MAXIMO {
		fmt.Println("Erro: forneça um tamanho válido entre 1 e", MAXIMO)
		return
	}

	tamanho = tamanhoConvertido

	gerarMatriz()
	fmt.Println("Matriz gerada aleatoriamente:")
	imprimirMatriz()

	eliminacaoGaussiana()

	fmt.Println("Soluções do sistema:")
	for i := 0; i < tamanho; i++ {
		fmt.Printf("x%d = %.5f\n", i, solucoes[i])
	}
}
