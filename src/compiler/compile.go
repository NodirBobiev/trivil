package compiler

import (
	"fmt"
	"io"
	"os"
	"path"
	"trivil/ast"
	"trivil/env"
	"trivil/genjava"
	"trivil/jasmin"
	"trivil/semantics"
)

var _ = fmt.Printf
var sourcePath string

type compileContext struct {
	main     *ast.Module
	imported map[string]*ast.Module // map[folder]

	testModulePath string // импорт путь для тестируемого модуля

	// упорядоченный список для обработки
	// головной модуль - в конце
	list   []*ast.Module
	status map[*ast.Module]int

	// Путь к папке для модуля, только для создания интерфейса модуля
	folders map[*ast.Module]string
}

func Compile(spath string) {
	sourcePath = spath
	var list = env.GetSources(spath)
	var src = list[0]
	if src.Err != nil {
		env.FatalError("ОКР-ОШ-ЧТЕНИЕ-ИСХОДНОГО", spath, src.Err.Error())
		return
	}

	var cc = &compileContext{
		imported: make(map[string]*ast.Module),
		folders:  make(map[*ast.Module]string),
	}

	cc.main = cc.parseModule(true, list)

	if env.ErrorCount() != 0 {
		return
	}

	cc.build()
}

func (cc *compileContext) build() {
	cc.orderedList()

	var gen *jasmin.Jasmin
	for _, m := range cc.list {

		if env.ErrorCount() != 0 {
			break
		}

		if *env.TraceCompile {
			fmt.Printf("Анализ и генерация: '%s'\n", m.Name)
			//fmt.Printf("Анализ и генерация: '%s' %p\n", m.Name, m)
		}

		gen = cc.process(m)
	}
	if gen != nil {
		destDir := *env.GenOut
		if destDir != "" {
			copyBuiltins(destDir)
			files := gen.Save(destDir)
			fmt.Printf("Generated: %d files\n", len(files))
			for _, f := range files {
				fmt.Printf("[%s]\n", f)
			}
		}
	}
	//if env.ErrorCount() == 0 && *env.DoGen && *env.BuildExe {
	//	genc.BuildExe(cc.list)
	//}
}

//=== process

func (cc *compileContext) process(m *ast.Module) *jasmin.Jasmin {

	ast.CurHost = m
	semantics.Analyse(m)
	ast.CurHost = nil

	if env.ErrorCount() != 0 {
		return nil
	}

	if *env.ShowAST >= 2 {
		fmt.Println(ast.SExpr(m))
	}

	if *env.MakeDef && m != cc.main {
		makeDef(m, cc.folders[m])
	}

	java := genjava.Generate(m, m == cc.main)
	return java
	//if *env.DoGen {
	//	genc.Generate(m, m == cc.main)
	//}
}

func (cc *compileContext) orderedList() {

	cc.list = make([]*ast.Module, 0)
	cc.status = make(map[*ast.Module]int)

	cc.traverse(cc.main, cc.main.Pos)
}

const (
	processing = 1
	processed  = 2
)

//=== traverse

func (cc *compileContext) traverse(m *ast.Module, pos int) {

	s, ok := cc.status[m]
	if ok {
		if s == processing {
			env.AddError(pos, "СЕМ-ЦИКЛ-ИМПОРТА", m.Name)
			cc.status[m] = processed
		}
		return
	}

	cc.status[m] = processing

	for _, i := range m.Imports {
		cc.traverse(i.Mod, i.Pos)
	}

	cc.status[m] = processed
	cc.list = append(cc.list, m)
}

//=== copy builtins

func copyFile(src, dst string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		panic("copyFile: src Open: " + err.Error())
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		panic("copyFile: dst Open: " + err.Error())
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		panic("copyFile: io Copy: " + err.Error())
	}
	err = destinationFile.Sync()
	if err != nil {
		panic("copyFile: dst file Sync: " + err.Error())
	}
}

func copyBuiltins(destDir string) {
	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		panic(fmt.Sprintf("copy builtins: make dir: %q", destDir))
	}
	sourceDir := "jasmin_runtime"
	filesToCopy := []string{
		"builtins_Print.j",
		"builtins_Scan.j",
	}
	for _, f := range filesToCopy {
		copyFile(path.Join(sourceDir, f), path.Join(destDir, f))
	}
}
