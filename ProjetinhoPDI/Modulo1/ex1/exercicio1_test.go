package main

import "testing"

func TestSumation2(t *testing.T) {
	type input struct {
		Description string
		Soma        int
		Num         int
		Expect      int
		ExpectError bool
	}

	cases := []input{
		{
			Description: "sucesso: passando 10 o valor deve ser 55",
			Soma:        0,
			Num:         10,
			Expect:      55,
			ExpectError: false,
		},
		{
			Description: "sucesso: passando 30 o valore deve ser 465",
			Soma:        0,
			Num:         30,
			Expect:      465,
			ExpectError: false,
		},
		{
			Description: "falha: passando 10 porem esperando um numero errado 3",
			Soma:        0,
			Num:         10,
			Expect:      3,
			ExpectError: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Description, func(t *testing.T) {
			sum := Sumation2(c.Soma, c.Num)
			t.Logf("o valor que recebi foi %d", sum)
			if sum != c.Expect && !c.ExpectError {
				t.Errorf("o valor esperado e %d, porem recebi %d", c.Expect, sum)
			}
		})
	}
}
