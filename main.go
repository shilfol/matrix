package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type calcMatrix struct {
	mutex  sync.Mutex
	matrix [][]float64
	size   int
	factor []float64
}

func (cm *calcMatrix) fillSlice(n int) {
	cm.size = n
	for i := 0; i < n; i++ {
		r := make([]float64, n)
		cm.matrix = append(cm.matrix, r)
	}
	cm.factor = make([]float64, n)

}

func (cm *calcMatrix) AllRandInit(n int) {
	var wg sync.WaitGroup

	cm.fillSlice(n)

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

func (cm *calcMatrix) SolveMatrix() {

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

func (cm *calcMatrix) ReadFileInit(n int, f *os.File) {
	cm.fillSlice(n)
	scanner := bufio.NewScanner(f)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			return
		}
		spl := strings.Fields(scanner.Text())
		for j := 0; j <= n; j++ {
			floats, _ := strconv.ParseFloat(spl[j], 64)
			if j == n {
				cm.factor[i] = floats
				continue
			}
			cm.matrix[i][j] = floats
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	mat := &calcMatrix{}
	N := 3
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil {
			N = n
		}
	}
	if len(os.Args) > 2 {
		file, err := os.Open(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer file.Close()

		mat.ReadFileInit(N, file)
	} else {
		mat.AllRandInit(N)
	}
	mat.Print()

	mat.SolveMatrix()

	mat.Print()
}
