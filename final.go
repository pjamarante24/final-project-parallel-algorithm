package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

var arrayUnordered []int
var arrayOrdered []int

func generateArrayWithRandomValues() {
	min := 0
	max := 10000
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100000; i++ {
		arrayUnordered = append(arrayUnordered, rand.Intn(max-min)+min)
	}
}

func printResult(title string, startTime time.Time) {
	var duration time.Duration = time.Since(startTime)
	fmt.Println(title + ": Terminado en " + fmt.Sprint(duration.Seconds()) + " segundos")
}

func printResultNotFound(title string, startTime time.Time) {
	var duration time.Duration = time.Since(startTime)
	fmt.Println(title + ": Terminado en " + fmt.Sprint(duration.Seconds()) + " segundos (Resultado no encontrado)")
}

func bubbleSort(array []int, print bool) []int {
	var startTime time.Time = time.Now()

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	if print {
		printResult("Ordenamiento por Bubble Sort", startTime)
	}
	return array
}

func quickSort(array []int, print bool) []int {
	var startTime time.Time = time.Now()

	quickSortHelper(array, 0, len(array)-1, print)

	if print {
		printResult("Ordenamiento por Quick Sort", startTime)
	}
	return array
}

func quickSortHelper(array []int, low, high int, print bool) {
	if low < high {
		var pivotIndex int = partition(array, low, high)
		quickSortHelper(array, low, pivotIndex-1, print)
		quickSortHelper(array, pivotIndex+1, high, print)
	}
}

func partition(array []int, low, high int) int {
	var pivotIndex int = low
	var pivotValue int = array[low]
	for i := low + 1; i <= high; i++ {
		if array[i] < pivotValue {
			pivotIndex++
			array[i], array[pivotIndex] = array[pivotIndex], array[i]
		}
	}
	array[low], array[pivotIndex] = array[pivotIndex], array[low]
	return pivotIndex
}

func insertionSort(array []int, print bool) []int {
	var startTime time.Time = time.Now()

	for i := 1; i < len(array); i++ {
		var j int = i
		var temp int = array[i]
		for j > 0 && array[j-1] > temp {
			array[j] = array[j-1]
			j--
		}
		array[j] = temp
	}

	if print {
		printResult("Ordenamiento por Insertion Sort", startTime)
	}
	return array
}

func linearSearch(array []int, value int) {
	var startTime time.Time = time.Now()

	for i := 0; i < len(array); i++ {
		if array[i] == value {
			printResult("Busqueda Secuenial", startTime)
			return
		}
	}

	printResultNotFound("Busqueda Secuenial", startTime)
}

func binarySearch(array []int, value int) {
	var startTime time.Time = time.Now()

	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2
		if array[mid] == value {
			printResult("Busqueda Binaria", startTime)
			return
		} else if array[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	printResultNotFound("Busqueda Binaria", startTime)
}

func main() {
	// Establece el numero de CPUs que pueden
	// ejecutar simultaneamente. En este caso 4
	// procesadores lógicos (esto es lo que hace
	// que el programa sea paralelo)
	runtime.GOMAXPROCS(4)

	generateArrayWithRandomValues()
	arrayOrdered = insertionSort(arrayUnordered, false)

	// for i := 0; i < len(arrayOrdered); i++ {
	// 	fmt.Printf("[%v]", arrayOrdered[i])
	// }

	go linearSearch(arrayOrdered, 413)
	go binarySearch(arrayOrdered, 413)
	go bubbleSort(arrayUnordered, true)
	go quickSort(arrayUnordered, true)
	go insertionSort(arrayUnordered, true)

	// Lee el teclado para detener la ejecucion
	a := 0
	fmt.Scanf("%d", &a)
}
