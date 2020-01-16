
/*
package main

import (
	"fmt"
)

func estudo() {

	x := 134
	x1 := 98349

	fmt.Println(x + x1)

}

// encapsulamento de estruturas

package main

import "fmt"

type Carro struct {
	Nome     string
	Idade    int
	ModelNum int
}

func (c Carro) Print() {
	fmt.Println(c)

}

func (c Carro) Dirigir() {
	fmt.Println("Dirigindo... vruumm")
}

func (c Carro) GetNome() string {

	return c.Nome
}

func main() {

	c := Carro{
		Nome:     "Chevy",
		Idade:    13,
		ModelNum: 34898234,
	}
	c.Print()
	c.Dirigir()
	fmt.Println(c.GetNome())
}


------------------------------------------------------------------------------------

package main

import (
	"fmt"
)

func Qualquercoisa(qualquercoisa interface{}) {

	fmt.Println(qualquercoisa)
}

func main() {

	fmt.Println("vim-go")
	Qualquercoisa(2.44)
	Qualquercoisa("Cílios")
	Qualquercoisa(struct{}{})

	//struct{ Pessoa string }{"Eae"}
	meumapa := make(map[string]interface{})
	meumapa["nome"] = "JIS83AS8JSAJ"
	meumapa["idade"] = 20
	fmt.Println(meumapa)

}

---------------------------------------------------------------------
func soma(){

		x := 124 int
		x1 :=432 int
        xf := ()int
		
		x1 - x 

		for xf <= 300 {
			fmt.Println("KKKKK")
		} else if {
			fmt.Println("WWWWWWW")
		} else {
			break
		}
		
	}

FAHRENHEIT
	*func main() {

	var f, n, n1, n2 float64

	fmt.Println("Insira uma temperatura em Fahrenheit (ºF) \n")
	fmt.Scanln(&f)

	n = f - 32
	n1 = (5 / 9)  
	n2 = n * n1

	fmt.Println(n, "\n", n1, "\n", n2)

}
