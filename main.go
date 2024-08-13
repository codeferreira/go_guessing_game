package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo de Advinhação")
	fmt.Println("Um número aleatório será sorteado. Tente acertar. O número é um inteiro entre 0 e 100.")

	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual o seu chute?")
		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)
		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("Não é um número inteiro")
			return
		}

		switch {
		case chuteInt < x:
			fmt.Println("Você errou! O número sorteado é maior que", chuteInt)
		case chuteInt > x:
			fmt.Println("Você errou! O número sorteado é menor que", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabéns, você acertou. O número sorteado foi: %d\n"+
					"Você acertou em %d tentativas.\n"+
					"Essas foram os seus chutes: %v\n",
				x, i+1, chutes[:i],
			)

			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Infelizmente, você perdeu. O número sorteado foi %d\n"+
			"Você teve 10 tentativas.\n"+
			"Essas foram os seus chutes: %v\n",
		x, chutes,
	)
}
