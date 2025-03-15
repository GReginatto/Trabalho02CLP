package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const MAXN int = 2000

func initializeInputs(n int, a [][]float64, b []float64) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = rand.Float64()
		}
		b[i] = rand.Float64()
	}
}

func gauss(n int, a [][]float64, b, x []float64) {
	for norm := 0; norm < n-1; norm++ {
		for row := norm + 1; row < n; row++ {
			multiplier := a[row][norm] / a[norm][norm]
			for col := norm; col < n; col++ {
				a[row][col] -= a[norm][col] * multiplier
			}
			b[row] -= b[norm] * multiplier
		}
	}
	for row := n - 1; row >= 0; row-- {
		x[row] = b[row]
		for col := row + 1; col < n; col++ {
			x[row] -= a[row][col] * x[col]
		}
		x[row] /= a[row][row]
	}
}

func main() {
	var n int = 4
	if len(os.Args) > 1 {
		if num, err := strconv.Atoi(os.Args[1]); err == nil {
			n = num
		}
	}
	a := make([][]float64, n)
	for i := range a {
		a[i] = make([]float64, n)
	}
	b := make([]float64, n)
	x := make([]float64, n)

	initializeInputs(n, a, b)

	originalA := make([][]float64, n)
	for i := range originalA {
		originalA[i] = make([]float64, n)
		copy(originalA[i], a[i]) 
	}
	originalB := make([]float64, n)
	copy(originalB, b) 

	start := time.Now()
	gauss(n, a, b, x)
	fmt.Printf("Tempo decorrido: %v\n", time.Since(start))

	fmt.Println("Matriz A:")
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}

	fmt.Println("Vetor B:")
	fmt.Println(b)

	fmt.Println("Vetor Solução X:")
	fmt.Println(x)

	fmt.Println("Verificação (A * X):")
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += originalA[i][j] * x[j]
		}
		fmt.Printf("Resultado[%d] = %v, Esperado[%d] = %v\n", i, sum, i, originalB[i])
	}
	fmt.Println("Verificação (A * X) com tolerância:")

	tolerance := 1e-10
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := 0; j < n; j++ {
			sum += originalA[i][j] * x[j]
		}
		if math.Abs(sum-originalB[i]) < tolerance {
			fmt.Printf("Resultado[%d] = %v, Esperado[%d] = %v → CORRETO\n", i, sum, i, originalB[i])
		} else {
			fmt.Printf("Resultado[%d] = %v, Esperado[%d] = %v → INCORRETO\n", i, sum, i, originalB[i])
		}
	}
}
