package external

import("fmt")

type ExternalStructure struct {
	External int
	internal string
}

var ExternVar ExternalStructure

func init() {
	ExternVar.External  = 10
	ExternVar.internal  = "文字列"
}


func ExternalTest() {
	fmt.Println("ExternalTest")
}

func internalTest() {
	fmt.Println("internalTest")
}
