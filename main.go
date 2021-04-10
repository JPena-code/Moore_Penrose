package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type MyMatrix []int

func read(reader *bufio.Reader, row, col int) []float64 {
	a := make([]float64, row*col)
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("A[%d][%d] = ", i, j)
			fmt.Fscan(reader, &a[count])
			count++
		}
	}
	fmt.Println()
	return a
}

func printMatrix(A mat.Matrix, s string) {
	pre := strings.Join(make([]string, len(s)+4), " ")
	fa := mat.Formatted(A, mat.Prefix(pre), mat.Squeeze())
	fmt.Printf("%s = %.2g\n\n", s, fa)
}

func main() {
	var row, col int
	fmt.Println("Ingrese el Numero de Filas y Columnas respectivamente: ")
	fmt.Scan(&row, &col)
	fmt.Println("Ingrese los valores de la matriz respectivamente: ")
	reader := bufio.NewReader(os.Stdin)
	data := read(reader, row, col)
	A := mat.NewDense(row, col, data)
	AproT := mat.NewDense(row, row, nil)
	AproT.Product(A.T(), A)
	y, x := AproT.Dims()
	invAT := mat.NewDense(y, x, nil)
	err := invAT.Inverse(AproT)
	if err != nil {
		log.Fatalf("No tiene inversa %v", err)
	}
	realInv := mat.NewDense(A.RawMatrix().Rows, A.RawMatrix().Cols, nil)
	realInv.Product(invAT, A.T())
	printMatrix(A, "A")
	printMatrix(realInv, "Inv")
}
