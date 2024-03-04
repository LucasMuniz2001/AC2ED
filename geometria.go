package geometria

// Area calcula a área de um retângulo
func Area(base, altura float64) float64 {
    return base * altura
}

// Perimetro calcula o perímetro de um retângulo
func Perimetro(base, altura float64) float64 {
    return 2 * (base + altura)
}
