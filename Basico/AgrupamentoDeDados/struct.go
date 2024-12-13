package main

import "fmt"

type pessoas struct {
	nome      string
	sobrenome string
	trabalha  bool
	telefone  int
}

func meuStruct() {

	pessoa1 := pessoas{
		nome:      "Guilherme",
		sobrenome: "Sousa",
		trabalha:  true,
		telefone:  9879654123,
	}

	pessoa2 := pessoas{
		nome:      "Antonio",
		sobrenome: "Lima",
		trabalha:  false,
		telefone:  9879654113,
	}

	fmt.Println(pessoa1, pessoa2)
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	cargo   string
	salario int
}

func structEmbutido() {
	pessoa1 := pessoa{
		nome:  "Alfredo",
		idade: 30,
	}
	pessoa2 := profissional{
		pessoa: pessoa{
			nome:  "Joao",
			idade: 32,
		},
		cargo:   "Assistente",
		salario: 1600,
	}

	pessoa3 := profissional{pessoa{"Mauricio", 30}, "Politico", 1000000}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.nome)
	fmt.Println(pessoa3.nome)
	fmt.Println(pessoa3)

}

func structAnonimo() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Guilherme",
		idade: 30,
	}

	fmt.Println(x)
}
