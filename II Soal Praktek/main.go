package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	score   int
	dice    []int
	nDice   int
	IsAlive bool
}

func main() {
	var nPlayer, nDice int
	fmt.Print("Masukkan jumlah pemain: ")
	fmt.Scanln(&nPlayer)
	fmt.Print("Masukkan jumlah dadu: ")
	fmt.Scanln(&nDice)

	fmt.Println()

	players := make([]Player, nPlayer)
	for i := range players {
		players[i].dice = make([]int, nDice)
		players[i].nDice = nDice
		players[i].IsAlive = true
	}

	gameOver := false
	round := 1
	alive := nPlayer

	for !gameOver && alive > 1 {
		fmt.Println("=============")
		fmt.Printf("Giliran %d lempar dadu\n", round)
		gameOver = true

		for i, p := range players {
			if len(p.dice) > 0 {
				gameOver = false
				fmt.Printf("Pemain #%d lempar dadu (%d): ", (i + 1), p.score)
				for j := range p.dice {
					players[i].dice[j] = rand.Intn(6) + 1
					fmt.Printf("%d ", players[i].dice[j])
				}
				fmt.Println()
			}
		}

		for i, p := range players {
			newDice := []int{}
			for j := 0; j < p.nDice; j++ {
				switch p.dice[j] {
				case 1:
					iNextPlayer := (i + 1) % nPlayer
					for players[iNextPlayer].nDice == 0 {
						iNextPlayer = (iNextPlayer + 1) % nPlayer
					}
					players[iNextPlayer].dice = append(players[iNextPlayer].dice, 1)
					if i == len(players)-1 {
						players[iNextPlayer].nDice += 1
					}
				case 6:
					players[i].score++
				default:
					newDice = append(newDice, p.dice[j])
				}
			}
			players[i].dice = newDice
			players[i].dice = append(players[i].dice, p.dice[p.nDice:]...)
			players[i].nDice = len(players[i].dice)
		}

		fmt.Println("Setelah evaluasi:")
		for i := 0; i < nPlayer; i++ {
			fmt.Printf("Pemain #%d (%d): ", i+1, players[i].score)
			if len(players[i].dice) == 0 && players[i].IsAlive {
				players[i].IsAlive = false
				alive--
			}
			for _, d := range players[i].dice {
				fmt.Printf("%d ", d)
			}
			fmt.Println()
		}

		round++
		fmt.Println("=============")
	}

	maxScore := -1
	winner := -1
	for i, p := range players {
		if p.score > maxScore {
			maxScore = p.score
			winner = i
		}
	}
	fmt.Printf("\nGame berakhir. Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", winner+1)

}
