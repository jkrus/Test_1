package main

import "log"

type (
	SlotForUSD struct {
		cash float64
	}

	SlotForEUR struct {
		cash float64
	}

	AnySlot interface {
		AddCash(float642 float64)
		GetCash() float64
	}

	RURSlotAdapter struct {
		course float64
		slot   AnySlot
	}

	SlotAdapterI interface {
		AddToSlot(cash float64)
		GetFromSlot() float64
	}
)

func NewSlotUSD(initCash float64) AnySlot {
	return &SlotForUSD{cash: initCash}
}

func (su *SlotForUSD) AddCash(cash float64) {
	su.cash += cash
}

func (su *SlotForUSD) GetCash() float64 {
	return su.cash
}

func NewSlotEUR(initCash float64) AnySlot {
	return &SlotForEUR{cash: initCash}
}

func (se *SlotForEUR) AddCash(cash float64) {
	se.cash += cash
}

func (se *SlotForEUR) GetCash() float64 {
	return se.cash
}

func NewSlotAdapter(slot AnySlot, course float64) SlotAdapterI {
	return &RURSlotAdapter{slot: slot, course: course}
}

func (sa *RURSlotAdapter) AddToSlot(cash float64) {
	sa.slot.AddCash(cash / sa.course)
}

func (sa *RURSlotAdapter) GetFromSlot() float64 {
	return sa.slot.GetCash() * sa.course
}

func main() {

	u := NewSlotUSD(100)

	rubUsd := NewSlotAdapter(u, 35)
	log.Println(rubUsd.GetFromSlot())

	rubUsd.AddToSlot(100)
	log.Println(rubUsd.GetFromSlot())

	e := NewSlotEUR(100)
	rubEur := NewSlotAdapter(e, 45)
	log.Println(rubEur.GetFromSlot())

	rubEur.AddToSlot(100)
	log.Println(rubEur.GetFromSlot())

}
