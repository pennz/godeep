package main

/* Need to define the notations
 * J -> the cost function
 * need to review ....
 */

import (
	"fmt"
	"math/rand"
)

const (
	r1 = 0.0784 // really need to add comments, I just cannot understand it now
	r2 = 0.8116
	r3 = 0.1100

	rate = 0.01
)

// MALG struct contains parameters for the model, which the algorithm will optimize
type MALG struct {
	w1 float64
	w2 float64
	// w3 float64

	f1 float64
	f2 float64
	f3 float64

	// R and J need update
	// TODO The J matix thing, the inverse
}

var malg MALG

func (m *MALG) UpdateJ() {
	//c1 = w1*f1 ...
	tri := m.f1 + m.f2 + m.f3
	tri1 := -m.f1/m.w1 + m.f3/m.w3
	tri2 := -m.f2/m.w2 + m.f3/m.w3

	c1 := m.f1 * m.w1
	c2 := m.f2 * m.w2
	c3 := m.f3 * m.w3

	trip2 := tri * tri

	J11 = (-c1*tri - c1*m.w1*tri1) / (m.w1 * m.w1) / trip2
	J21 = (-c1 * tri2) / m.w1 / trip2

	J12 = (-c2 * tri1) / m.w2 / trip2
	J22 = (-c2*tri - c2*m.w2*tri2) / (m.w2 * m.w2) / trip2
}
func (m *MALG) Calculate() float64 {
	return m.w1*m.f1 + m.w2*m.f2 + m.w3*m.f3
}

func (m *MALG) SetWeight(a, b, c float64) {
	m.w1, m.w2, m.w3 = a, b, c
}
func (m *MALG) RandWeight() {
	m.w1 = 0.02 * float64(rand.Intn(10))

	m.w2 = rand.Float64()
	m.w3 = 1 - m.w1 - m.w2
}
func (m *MALG) RandConfig() {
	m.f1 = 4
	m.f2 = float64(rand.Intn(40)) + 30
	m.f3 = 100 - m.f2
}

func (m *MALG) UpdateWeight() {
	m.w1 += -rate * (R1*1 + R2*1)
	m.w2 += -rate * (R1*1 + R2*1)

	m.w3 = 1 - m.w1 - m.w2
}

// need to know w1, w2, w3 / R1,R2,R3 / a struct will be better, easy to handle
// for first run it, then improve it.

func main() {
	malg.RandConfig()
	malg.SetWeight()
	fmt.Println(malg.Calculate())

	malg.RandConfig()
	malg.SetWeight()
	fmt.Println(malg.Calculate())
}
