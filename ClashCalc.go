package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

type State struct {
	ACoins int
	BCoins int
}

type BuffDebuff struct {
	Power int
	Coin  int
}

func calculateHeadChance(sanity int) float64 {
	return 0.5 + float64(sanity)*0.01
}

func calculatePower(basePower int, buffsDebuffs BuffDebuff, coins []bool, declaredDuel int, paralyze int, tremorChains int) int {
	power := basePower + buffsDebuffs.Power + declaredDuel
	for i, coin := range coins {
		if i < paralyze {
			continue
		}
		if coin {
			power += buffsDebuffs.Coin
		}
	}
	if tremorChains >= 10 {
		power -= tremorChains / 10
		if power < basePower-3 {
			power = basePower - 3
		}
	}
	return power
}

func createTransitionMatrix(sanityA, sanityB int, basePowerA, basePowerB int, buffsA, buffsB BuffDebuff, declaredDuelA, declaredDuelB, paralyzeA, paralyzeB, tremorChainsA, tremorChainsB int) ([][]float64, map[State]int) {
	states := []State{
		{2, 2}, {2, 1}, {1, 2}, {1, 1},
	}
	stateIndex := make(map[State]int)
	for i, state := range states {
		stateIndex[state] = i
	}

	matrix := make([][]float64, len(states))
	for i := range matrix {
		matrix[i] = make([]float64, len(states)+2)
	}

	headChanceA := calculateHeadChance(sanityA)
	headChanceB := calculateHeadChance(sanityB)

	rand.Seed(time.Now().UnixNano())

	for _, state := range states {
		if state.ACoins > 0 && state.BCoins > 0 {
			coinsA := make([]bool, state.ACoins)
			coinsB := make([]bool, state.BCoins)
			for i := range coinsA {
				coinsA[i] = rand.Float64() < headChanceA
			}
			for i := range coinsB {
				coinsB[i] = rand.Float64() < headChanceB
			}

			powerA := calculatePower(basePowerA, buffsA, coinsA, declaredDuelA, paralyzeA, tremorChainsA)
			powerB := calculatePower(basePowerB, buffsB, coinsB, declaredDuelB, paralyzeB, tremorChainsB)

			if powerA > powerB {
				nextState := State{state.ACoins, state.BCoins - 1}
				matrix[stateIndex[state]][stateIndex[nextState]] = 1.0
			} else {
				nextState := State{state.ACoins - 1, state.BCoins}
				matrix[stateIndex[state]][stateIndex[nextState]] = 1.0
			}
		}
	}

	matrix[stateIndex[State{1, 0}]][len(states)] = 1.0
	matrix[stateIndex[State{0, 1}]][len(states)+1] = 1.0

	return matrix, stateIndex
}

func invertMatrix(Q [][]float64) *mat.Dense {
	qr := mat.NewDense(len(Q), len(Q[0]), nil)
	for i := range Q {
		for j := range Q[i] {
			qr.Set(i, j, Q[i][j])
		}
	}

	I := mat.NewDense(len(Q), len(Q), nil)
	for i := 0; i < len(Q); i++ {
		I.Set(i, i, 1.0)
	}
	I.Sub(I, qr)

	var invI mat.Dense
	err := invI.Inverse(I)
	if err != nil {
		panic("Matrix is singular and cannot be inverted")
	}

	return &invI
}

func calculateAbsorptionProbabilities(N *mat.Dense, R [][]float64) []float64 {
	r := mat.NewDense(len(R), len(R[0]), nil)
	for i := range R {
		for j := range R[i] {
			r.Set(i, j, R[i][j])
		}
	}

	var B mat.Dense
	B.Mul(N, r)

	probabilities := make([]float64, B.RawMatrix().Cols)
	for i := 0; i < B.RawMatrix().Cols; i++ {
		probabilities[i] = B.At(0, i)
	}

	return probabilities
}

func main() {
	var SanAl, SanIn, MoedasAl, MoedasIn, PoderBaseAl, PoderBaseIn, PoderMoedaAl, PoderMoedaIn, Anormalidade int
	var DistintorAl [6]int
	var DistintorIn [7]int

	fmt.Println("What`s the ally sanity?")
	fmt.Scan(&SanAl)
	fmt.Println("How many coins are in the attack?")
	fmt.Scan(&MoedasAl)
	fmt.Println("What`s the base power of the attack?")
	fmt.Scan(&PoderBaseAl)
	fmt.Println("What`s the coin power of the attack?")
	fmt.Scan(&PoderMoedaAl)
	fmt.Println("Power/Level Up? (0 for none)")
	fmt.Scan(&DistintorAl[0])
	fmt.Println("Coin Power Up? (including skill effects) (0 for none)")
	fmt.Scan(&DistintorAl[1])
	fmt.Println("Declared Duel? (0 for no, 1 for yes)")
	fmt.Scan(&DistintorAl[2])
	fmt.Println("Power/Level Down? (0 for none)")
	fmt.Scan(&DistintorAl[3])
	fmt.Println("Coin Power Down? (0 for none)")
	fmt.Scan(&DistintorAl[4])
	fmt.Println("Paralyze? (0 for none)")
	fmt.Scan(&DistintorAl[5])

	fmt.Println("Does the enemy have sanity (1 for yes, 2 for no)")
	fmt.Scan(&Anormalidade)
	if Anormalidade == 1 {
		fmt.Println("What`s the enemy sanity?")
		fmt.Scan(&SanIn)
	} else {
		SanIn = 10
	}
	fmt.Println("How many coins are in the attack?")
	fmt.Scan(&MoedasIn)
	fmt.Println("What`s the base power of the attack?")
	fmt.Scan(&PoderBaseIn)
	fmt.Println("What`s the coin power of the attack?")
	fmt.Scan(&PoderMoedaIn)
	fmt.Println("Power/Level Up? (0 for none)")
	fmt.Scan(&DistintorIn[0])
	fmt.Println("Coin Power Up? (including skill effects) (0 for none)")
	fmt.Scan(&DistintorIn[1])
	fmt.Println("Power/Level Down? (0 for none)")
	fmt.Scan(&DistintorIn[2])
	fmt.Println("Coin Power Down? (0 for none)")
	fmt.Scan(&DistintorIn[3])
	fmt.Println("Paralyze? (0 for none)")
	fmt.Scan(&DistintorIn[4])
	fmt.Println("Echoes of The Manor? (0 para não, 1 para sim)")
	fmt.Scan(&DistintorIn[5])
	fmt.Println("Tremor - Chains? (0 for none)")
	fmt.Scan(&DistintorIn[6])

	if DistintorIn[5] == 1 {
		SanIn -= 10
	}
	if SanIn < 0 {
		SanIn = 0
	}

	buffsA := BuffDebuff{Power: DistintorAl[0] - DistintorAl[3], Coin: PoderMoedaAl + DistintorAl[1] - DistintorAl[4]}
	buffsB := BuffDebuff{Power: DistintorIn[0] - DistintorIn[2], Coin: PoderMoedaIn + DistintorIn[1] - DistintorIn[3]}

	matrix, stateIndex := createTransitionMatrix(SanAl, SanIn, PoderBaseAl, PoderBaseIn, buffsA, buffsB, DistintorAl[2], 0, DistintorAl[5], DistintorIn[4], 0, DistintorIn[6])

	Q := [][]float64{
		{matrix[0][0], matrix[0][1], matrix[0][2], matrix[0][3]},
		{matrix[1][0], matrix[1][1], matrix[1][2], matrix[1][3]},
		{matrix[2][0], matrix[2][1], matrix[2][2], matrix[2][3]},
		{matrix[3][0], matrix[3][1], matrix[3][2], matrix[3][3]},
	}

	R := [][]float64{
		{matrix[0][4], matrix[0][5]},
		{matrix[1][4], matrix[1][5]},
		{matrix[2][4], matrix[2][5]},
		{matrix[3][4], matrix[3][5]},
	}

	N := invertMatrix(Q)

	probabilities := calculateAbsorptionProbabilities(N, R)
	fmt.Printf("Status index: %v\n", stateIndex)
	fmt.Printf("Ally victory probability: %.2f%%\n", probabilities[1]*100)
	fmt.Printf("Enemy victory probability: %.2f%%\n", probabilities[0]*100)
}
