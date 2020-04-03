package main

import "fmt"

type Polynomial *PolyNode

type PolyNode struct {
	coef  int
	expon int
	link  Polynomial
}

func main() {
	// f1 = 5x^10 + 10x^3 + 5
	// f2 = 10x^8 - 4x^7 + 5x^3 -1
	// sum = 5x^10 + 10x^8 - 4x^7 + 15x^3 4

	// input1  3 5 10 10 3 5 0
	// inout2  4 10 8 -4 7 5 3 -1 0
	input1 := []int{3, 5, 10, 10, 3, 5, 0}
	input2 := []int{4, 10, 8, -4, 7, 5, 3, -1, 0}

	fx1 := ReadOnly(input1)
	fx2 := ReadOnly(input2)

	sum := Add(fx1, fx2)
	for sum != nil {
		fmt.Printf(" %dx^%d ", sum.coef, sum.expon)
		sum = sum.link
	}
}

func IsEmpty(Pnode *PolyNode) bool {
	if Pnode.link != nil {
		return false
	}
	return true
}

func ReadOnly(input []int) *PolyNode {
	N := input[0]
	Pnode := &PolyNode{
		coef:  0,
		expon: 0,
		link:  nil,
	}
	i, j := 1, 1
	for ; i <= N; i++ {
		Attach(input[j], input[j+1], Pnode)
		j = j + 2
	}
	Pnode = Pnode.link
	return Pnode
}

func Attach(coef, expon int, Pnode *PolyNode) {
	P := &PolyNode{
		coef,
		expon,
		nil,
	}
	if !IsEmpty(Pnode) {
		for Pnode.link != nil {
			Pnode = Pnode.link
		}
	}
	Pnode.link = P
}

func Add(P1 Polynomial, P2 Polynomial) Polynomial {
	P := &PolyNode{
		coef:  0,
		expon: 0,
		link:  nil,
	}
	Rear := P
	for P1 != nil && P2 != nil {
		switch Compare(P1.expon, P2.expon) {
		case 1:
			Attach(P1.coef, P1.expon, Rear)
			P1 = P1.link
		case -1:
			Attach(P2.coef, P2.expon, Rear)
			P2 = P2.link
		case 0:
			fmt.Println(P1.coef, P2.coef)
			sum := P1.coef + P2.coef
			if sum != 0 {
				fmt.Println(sum)
				Attach(sum, P1.expon, Rear)
				P1 = P1.link
				P2 = P2.link
			}
		}
	}
	for P1 != nil {
		Attach(P1.coef, P1.expon, Rear)
		P1 = P1.link
	}
	for P2 != nil {
		Attach(P2.coef, P2.expon, Rear)
		P2 = P2.link
	}
	return Rear
}

func Compare(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
