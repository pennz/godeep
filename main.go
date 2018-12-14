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

type _MALGModel struct {
	w1 float64
	w2 float64
	w3 float64

	f1 float64
	f2 float64
	f3 float64

	// R and J need update
	// TODO The J matix thing, the inverse
}

// MALG struct contains parameters for the model, which the algorithm will optimize
type MALG struct {
	J11, J21, J22, J12 float64
	R1, R2             float64
	_MALGModel
}

// updateJ will update the cost function J
func (m *MALG) updateJ() {
	//c1 = w1*f1 ...
	tri := m.f1 + m.f2 + m.f3
	tri1 := -m.f1/m.w1 + m.f3/m.w3
	tri2 := -m.f2/m.w2 + m.f3/m.w3

	c1 := m.f1 * m.w1
	c2 := m.f2 * m.w2
	//c3 := m.f3 * m.w3

	trip2 := tri * tri

	m.J11 = (-c1*tri - c1*m.w1*tri1) / (m.w1 * m.w1) / trip2
	m.J21 = (-c1 * tri2) / m.w1 / trip2

	m.J12 = (-c2 * tri1) / m.w2 / trip2
	m.J22 = (-c2*tri - c2*m.w2*tri2) / (m.w2 * m.w2) / trip2
}

// Calculate just uses the model to calculate final result/output
func (m *MALG) Calculate() float64 {
	return m.w1*m.f1 + m.w2*m.f2 + m.w3*m.f3
}

// SetWeight set initial weight?
func (m *MALG) SetWeight(a, b, c float64) {
	m.w1, m.w2, m.w3 = a, b, c
}

// setRandWeight will set random weight
func (m *MALG) setRandWeight() {
	m.w1 = 0.02 * float64(rand.Intn(10))

	m.w2 = rand.Float64()
	m.w3 = 1 - m.w1 - m.w2
}

// RandConfig sets random configuration for the algrithm?
func (m *MALG) RandConfig() {
	m.f1 = 4
	m.f2 = float64(rand.Intn(40)) + 30
	m.f3 = 100 - m.f2
}

func (m *MALG) updateWeight() {
	m.w1 += -rate * (m.R1*1 + m.R2*1)
	m.w2 += -rate * (m.R1*1 + m.R2*1)

	m.w3 = 1 - m.w1 - m.w2
}

func main() {
	// need to know w1, w2, w3 / R1,R2,R3 / a struct will be better, easy to handle
	// for first run it, then improve it.
	malg := new(MALG)
	malg.RandConfig()
	malg.setRandWeight()
	fmt.Println(malg.Calculate())
}
