package main

import "log"

type (
	Human struct {
		field1 string
		field2 string
	}

	Action1 struct {
		humanActions HumanActions
	}

	Action11 struct {
		HumanActions
	}

	Action2 struct {
		human *Human
	}

	Action22 struct {
		*Human
	}

	HumanActions interface {
		Method1()
		Method2()
	}
)

func NewAction_1(actions HumanActions) *Action1 {
	return &Action1{
		humanActions: actions,
	}
}

func NewAction_11(actions HumanActions) *Action11 {
	return &Action11{actions}
}

func NewAction_2(human *Human) *Action2 {
	return &Action2{human: human}
}

func NewAction_22(human *Human) *Action22 {
	return &Action22{human}
}

func NewHuman(f1, f2 string) *Human {
	return &Human{f1, f2}
}

func NewHumanI(f1, f2 string) HumanActions {
	return &Human{f1, f2}
}

func (h *Human) Method1() {
	log.Println("I'm method 1")
}

func (h *Human) Method2() {
	log.Println("I'm method 2")
}

func main() {
	humanI := NewHumanI("1", "2")

	// Первый вариант - написать интерфейс которому должна соответсвовать структура Human
	// В этом случае Action_1 унаследует методы структуры Human
	actions1 := NewAction_1(humanI)
	actions1.humanActions.Method1()
	actions1.humanActions.Method2()

	actions11 := NewAction_11(humanI)
	actions11.Method1()
	actions11.Method2()

	// Второй вариант менее удачный, потому что наследует ещё и поля родительской структуры.
	// К полям лучше написать Get & Set методы и использовать первый вариант.
	human := NewHuman("1", "2")
	actions2 := NewAction_2(human)
	actions2.human.Method1()
	actions2.human.Method2()

	actions22 := NewAction_22(human)
	actions22.Method1()
	actions22.Method2()
}
