package main

import (
	"trivil/genjava/tests/common"
	"trivil/genjava/tests/scenarious"
)

var tests = []common.Test{
	&scenarious.SimpleTest{
		ProgramName: "cчётчик",
		PackagePath: "/home/cyrus/trivil/examples/simples/cчётчик",
		OutputPath:  "/home/cyrus/trivil/examples/simples/cчётчик/output.txt",
	},
	&scenarious.SimpleTest{
		ProgramName: "арифметика",
		PackagePath: "/home/cyrus/trivil/examples/simples/арифметика",
		OutputPath:  "/home/cyrus/trivil/examples/simples/арифметика/output.txt",
	},
	&scenarious.SimpleTest{
		ProgramName: "вещ64-тест",
		PackagePath: "/home/cyrus/trivil/examples/simples/вещ64-тест",
		OutputPath:  "/home/cyrus/trivil/examples/simples/вещ64-тест/output.txt",
	},
	&scenarious.SimpleTest{
		ProgramName: "вывод",
		PackagePath: "/home/cyrus/trivil/examples/simples/вывод",
		OutputPath:  "/home/cyrus/trivil/examples/simples/вывод/output.txt",
	},
	&scenarious.SimpleTest{
		ProgramName: "импортер",
		PackagePath: "/home/cyrus/trivil/examples/simples/импортер",
		OutputPath:  "/home/cyrus/trivil/examples/simples/импортер/output.txt",
	},
	&scenarious.SimpleTest{
		ProgramName: "перемена",
		PackagePath: "/home/cyrus/trivil/examples/simples/перемена",
		OutputPath:  "/home/cyrus/trivil/examples/simples/перемена/output.txt",
	},
	&scenarious.SimpleTest{
		ProgramName: "простой-класс",
		PackagePath: "/home/cyrus/trivil/examples/simples/простой-класс",
		OutputPath:  "/home/cyrus/trivil/examples/simples/простой-класс/output.txt",
	},
}

func main() {
	for _, t := range tests {
		err := t.Run()
		if err != nil {
			panic(err)
		}
	}
}
