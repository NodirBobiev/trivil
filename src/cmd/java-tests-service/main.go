package main

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"trivil/genjava/tests/ansi"
	"trivil/genjava/tests/common"
	"trivil/genjava/tests/scenarious"

	"trivil/genjava/tests/logger"
)

var tests = []common.Test{
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-cчётчик",
		},
		ProgramName: "cчётчик",
		PackagePath: "/home/cyrus/trivil/examples/simples/cчётчик",
		OutputPath:  "/home/cyrus/trivil/examples/simples/cчётчик/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-арифметика",
		},
		ProgramName: "арифметика",
		PackagePath: "/home/cyrus/trivil/examples/simples/арифметика",
		OutputPath:  "/home/cyrus/trivil/examples/simples/арифметика/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-арифметика-сумма",
		},
		ProgramName: "арифметика-сумма",
		PackagePath: "/home/cyrus/trivil/examples/simples/арифметика/сумма",
		OutputPath:  "/home/cyrus/trivil/examples/simples/арифметика/сумма/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/арифметика/сумма/input.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-арифметика-разница",
		},
		ProgramName: "арифметика-разница",
		PackagePath: "/home/cyrus/trivil/examples/simples/арифметика/разница",
		OutputPath:  "/home/cyrus/trivil/examples/simples/арифметика/разница/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/арифметика/разница/input.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-арифметика-площадь",
		},
		ProgramName: "арифметика-площадь",
		PackagePath: "/home/cyrus/trivil/examples/simples/арифметика/площадь",
		OutputPath:  "/home/cyrus/trivil/examples/simples/арифметика/площадь/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/арифметика/площадь/input.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-вещ64-тест",
		},
		ProgramName: "вещ64-тест",
		PackagePath: "/home/cyrus/trivil/examples/simples/вещ64-тест",
		OutputPath:  "/home/cyrus/trivil/examples/simples/вещ64-тест/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-вывод",
		},
		ProgramName: "вывод",
		PackagePath: "/home/cyrus/trivil/examples/simples/вывод",
		OutputPath:  "/home/cyrus/trivil/examples/simples/вывод/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-импортер",
		},
		ProgramName: "импортер",
		PackagePath: "/home/cyrus/trivil/examples/simples/импортер",
		OutputPath:  "/home/cyrus/trivil/examples/simples/импортер/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-перемена",
		},
		ProgramName: "перемена",
		PackagePath: "/home/cyrus/trivil/examples/simples/перемена",
		OutputPath:  "/home/cyrus/trivil/examples/simples/перемена/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-класс",
		},
		ProgramName: "простой-класс",
		PackagePath: "/home/cyrus/trivil/examples/simples/простой-класс",
		OutputPath:  "/home/cyrus/trivil/examples/simples/простой-класс/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-если",
		},
		ProgramName: "простой-если",
		PackagePath: "/home/cyrus/trivil/examples/simples/простой-если",
		OutputPath:  "/home/cyrus/trivil/examples/simples/простой-если/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-рекурсия-диапазон",
		},
		ProgramName: "рекурсия-диапазон",
		PackagePath: "/home/cyrus/trivil/examples/simples/рекурсия/диапазон",
		OutputPath:  "/home/cyrus/trivil/examples/simples/рекурсия/диапазон/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-рекурсия-фибоначи",
		},
		ProgramName: "рекурсия-фибоначи",
		PackagePath: "/home/cyrus/trivil/examples/simples/рекурсия/фибоначи",
		OutputPath:  "/home/cyrus/trivil/examples/simples/рекурсия/фибоначи/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-пока-интервал",
		},
		ProgramName: "пока-интервал",
		PackagePath: "/home/cyrus/trivil/examples/simples/пока/интервал",
		OutputPath:  "/home/cyrus/trivil/examples/simples/пока/интервал/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-пока-фибоначи",
		},
		ProgramName: "пока-фибоначи",
		PackagePath: "/home/cyrus/trivil/examples/simples/пока/фибоначи",
		OutputPath:  "/home/cyrus/trivil/examples/simples/пока/фибоначи/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-пока-н-й-фибоначи",
		},
		ProgramName: "пока-н-й-фибоначи",
		PackagePath: "/home/cyrus/trivil/examples/simples/пока/н-й-фибоначи",
		OutputPath:  "/home/cyrus/trivil/examples/simples/пока/н-й-фибоначи/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/пока/н-й-фибоначи/input.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-наследование-вызов-метода",
		},
		ProgramName: "наследование-вызов-метода",
		PackagePath: "/home/cyrus/trivil/examples/simples/наследование/вызов-метода",
		OutputPath:  "/home/cyrus/trivil/examples/simples/наследование/вызов-метода/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-наследование-вызов-метода-переопределение",
		},
		ProgramName: "наследование-вызов-метода-переопределение",
		PackagePath: "/home/cyrus/trivil/examples/simples/наследование/переопределение",
		OutputPath:  "/home/cyrus/trivil/examples/simples/наследование/переопределение/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-прервать-стоп",
		},
		ProgramName: "прервать-стоп",
		PackagePath: "/home/cyrus/trivil/examples/simples/прервать/стоп",
		OutputPath:  "/home/cyrus/trivil/examples/simples/прервать/стоп/output.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-эхо-цел64",
		},
		ProgramName: "эхо-цел64",
		PackagePath: "/home/cyrus/trivil/examples/simples/эхо/цел64",
		OutputPath:  "/home/cyrus/trivil/examples/simples/эхо/цел64/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/эхо/цел64/input.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-эхо-вещ64",
		},
		ProgramName: "эхо-вещ64",
		PackagePath: "/home/cyrus/trivil/examples/simples/эхо/вещ64",
		OutputPath:  "/home/cyrus/trivil/examples/simples/эхо/вещ64/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/эхо/вещ64/input.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-эхо-строка",
		},
		ProgramName: "эхо-строка",
		PackagePath: "/home/cyrus/trivil/examples/simples/эхо/строка",
		OutputPath:  "/home/cyrus/trivil/examples/simples/эхо/строка/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/эхо/строка/input.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-эхо-линия",
		},
		ProgramName: "эхо-линия",
		PackagePath: "/home/cyrus/trivil/examples/simples/эхо/линия",
		OutputPath:  "/home/cyrus/trivil/examples/simples/эхо/линия/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/эхо/линия/input.txt",
	},
	&scenarious.SimpleTest{
		TestBase: common.TestBase{
			Name: "простой-эхо-текст",
		},
		ProgramName: "эхо-текст",
		PackagePath: "/home/cyrus/trivil/examples/simples/эхо/текст",
		OutputPath:  "/home/cyrus/trivil/examples/simples/эхо/текст/output.txt",
		InputPath:   "/home/cyrus/trivil/examples/simples/эхо/текст/input.txt",
	},
}

func runTest(t common.Test) error {
	log := logger.InitLogger(io.Discard)
	log = log.With(
		zap.String("test_name", t.GetName()),
	)
	zap.ReplaceGlobals(log)
	return t.Run()
}

func main() {

	failedTests := 0
	for _, t := range tests {
		err := runTest(t)
		if err != nil {
			failedTests++
			fmt.Println(ansi.Red.Wrap(" ✗ FAIL | %s\n\t  ERROR: %s", t.GetName(), err))
		} else {
			fmt.Println(ansi.Green.Wrap(" ✓ OK   | %s", t.GetName()))
		}
	}

	success := 100.0 * float64(len(tests)-failedTests) / float64(len(tests))
	color := ansi.Green
	text := color.Wrap("\t%.1f%% passed", success)
	if failedTests > 0 {
		color = ansi.Yellow
		text = color.Wrap(" %.1f%% passed | %d out %d failed", success, failedTests, len(tests))
	}
	fmt.Println()
	fmt.Println(color.Wrap("============= Report ============="))
	fmt.Println(text)
	fmt.Println(color.Wrap("=================================="))

}
