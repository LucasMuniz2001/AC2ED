# AC2ED
package main

import (
    "fmt"
    "math"
    "geometria"
)


func criaVetor() {
    vetor := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    fmt.Println("Elementos do vetor:")
    for _, elemento := range vetor {
        fmt.Println(elemento)
    }
}

func inverterString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}

type Ponto struct {
    X, Y float64
}

func (p Ponto) DistanciaOrigem() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}


func main() {
   
    criaVetor()

   
    var entrada string
    fmt.Println("Digite uma string:")
    fmt.Scanln(&entrada)
    fmt.Println("String invertida:", inverterString(entrada))

    ponto := Ponto{3, 4}
    fmt.Printf("Distância até a origem: %.2f\n", ponto.DistanciaOrigem())

    
    var base, altura float64
    fmt.Println("Digite a base do retângulo:")
    fmt.Scanln(&base)
    fmt.Println("Digite a altura do retângulo:")
    fmt.Scanln(&altura)

    area := geometria.Area(base, altura)
    perimetro := geometria.Perimetro(base, altura)

    fmt.Printf("Área do retângulo: %.2f\n", area)
    fmt.Printf("Perímetro do retângulo: %.2f\n", perimetro)
}
