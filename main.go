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
}

func (cm *calcMatrix) allRandInit(n int) {
	cm.size = n
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		r := make([]float64, cm.size)
		cm.matrix = append(cm.matrix, r)
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go cm.rowRandInit(i, &wg)
	}
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

func (cm *calcMatrix) Print() {
	for i := 0; i < cm.size; i++ {
		for j := 0; j < cm.size; j++ {
			fmt.Printf("%.5f ", cm.matrix[i][j])
		}
		fmt.Println()
	}

}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	mat := &calcMatrix{}
	mat.allRandInit(3)

	mat.Print()
}
