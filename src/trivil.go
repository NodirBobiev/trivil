package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"trivil/compiler"
	"trivil/env"
	"trivil/lexer"
)

const (
	versionCompiler = "0.71"
	versionLanguage = "0.9.0"
)

// команды
const (
	prefix  = "+"
	testOne = prefix + "тест"
)

func main() {
	// флаги определены в env.flags
	flag.Parse()
	var arg = flag.Arg(0)
	var justHelp = arg == ""

	var command = ""
	if !justHelp && strings.HasPrefix(arg, prefix) {
		switch arg {
		case testOne:
			command = arg
			arg = flag.Arg(1)
			if arg == "" {
				justHelp = true
			}
		default:
			justHelp = true
		}
	}

	if justHelp {
		fmt.Println("Использование:")
		fmt.Println("  построить: tric (folder | file.tri)")
		fmt.Println("  тестировать: tric +тест folder")
		os.Exit(1)
	}

	fmt.Printf("Тривиль-0 компилятор (Go) v%s (язык %s)\n", versionCompiler, versionLanguage)
	env.Init()

	//fmt.Printf("%v\n", src.Bytes)
	if *env.JustLexer {
		testLexer(arg)
	} else if command == testOne {
		compiler.TestOne(arg)
	} else {
		compiler.Compile(arg)
	}

	env.ShowErrors()

	if env.ErrorCount() == 0 {
		fmt.Printf("Без ошибок\n")
	} else {
		os.Exit(1)
	}

}

func testLexer(arg string) {

	var files = env.GetSources(arg)
	var src = files[0]
	if src.Err != nil {
		fmt.Printf("Ошибка чтения исходного файла '%s': %s\n", arg, src.Err.Error())
		os.Exit(1)
	}

	var lex = new(lexer.Lexer)
	lex.Init(src)

	for true {
		pos, tok, lit := lex.Scan()
		if tok == lexer.EOF {
			break
		}

		_, line, col := env.SourcePos(pos)
		if lit == "" {
			fmt.Printf("%d:%d %v\n", line, col, tok)
		} else {
			fmt.Printf("%d:%d %v '%s'\n", line, col, tok, lit)
		}
	}
}
