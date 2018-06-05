package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type calcMatrix struct {
	mutex  sync.Mutex
	matrix [][]float64
	size   int
	factor []float64
}

func (cm *calcMatrix) allRandInit(n int) {
	cm.size = n
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		r := make([]float64, cm.size)
		cm.matrix = append(cm.matrix, r)
	}
	cm.factor = make([]float64, cm.size)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go cm.rowRandInit(i, &wg)
	}
	wg.Add(1)
	go cm.factorRandInit(&wg)
	wg.Wait()

}

func (cm *calcMatrix) rowRandInit(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	for j := 0; j < cm.size; j++ {
		cm.matrix[i][j] = rand.Float64() * 10.0
	}
}

func (cm *calcMatrix) factorRandInit(wg *sync.WaitGroup) {
	defer wg.Done()
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	for j := 0; j < cm.size; j++ {
		cm.factor[j] = rand.Float64() * 10.0
	}
}

func (cm *calcMatrix) Print() {
	for i := 0; i < cm.size; i++ {
		for j := 0; j < cm.size; j++ {
			fmt.Printf("%.5f ", cm.matrix[i][j])
		}
		fmt.Println()
	}

	for j := 0; j < cm.size; j++ {
		fmt.Printf("x%d: %.5f\n", j, cm.factor[j])
	}
}

func (cm *calcMatrix) solveMatrix() {

	cm.solveForward()
	cm.solveBackward()
	cm.solveDivide()

}

func (cm *calcMatrix) solveForward() {
	for i := 0; i < cm.size-1; i++ {
		for j := i + 1; j < cm.size; j++ {
			q := cm.matrix[j][i] / cm.matrix[i][i]
			for k := i; k < cm.size; k++ {
				cm.matrix[j][k] -= q * cm.matrix[i][k]
			}

			cm.factor[j] -= q * cm.factor[i]
		}
	}
}

func (cm *calcMatrix) solveBackward() {
	for i := cm.size - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			q := cm.matrix[j][i] / cm.matrix[i][i]
			cm.matrix[j][i] -= q * cm.matrix[i][i]
			cm.factor[j] -= q * cm.factor[i]
		}
	}
}

func (cm *calcMatrix) solveDivide() {
	for i := 0; i < cm.size; i++ {
		if cm.matrix[i][i] != 0.0 {
			cm.factor[i] /= cm.matrix[i][i]
			cm.matrix[i][i] /= cm.matrix[i][i]
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	mat := &calcMatrix{}
	mat.allRandInit(5)

	mat.solveMatrix()

	mat.Print()
}
