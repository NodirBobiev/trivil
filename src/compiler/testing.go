package compiler

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

// TODO: сделать конфигурационный файл (для всего компилятора)
const (
	testsFolderName   = "_тест_"
	testEntryPath     = "стд::тестирование/тест-вход"
	patternImportPath = "№путь-тестируемого№"
	patternModuleName = "№имя-тестируемого№"
)

func TestOne(tpath string) {

	env.Normalizer.Process(tpath)
	if env.Normalizer.Err != nil {
		env.AddProgramError("ОКР-ОШ-ПУТЬ", tpath, env.Normalizer.Err.Error())
		return
	}

	var npath = env.Normalizer.NPath

	var cc = &compileContext{
		imported:       make(map[string]*ast.Module),
		folders:        make(map[*ast.Module]string),
		testModulePath: tpath,
	}

	// Подготовка головного модуля тестирования
	var list = prepareTestEntry(tpath, npath)
	if env.ErrorCount() != 0 {
		return
	}

	cc.main = cc.parseModule(true, list)

	if env.ErrorCount() != 0 {
		return
	}

	cc.build()
}

//==== добавление тестов ====

func collectTestSources(tpath, npath string) []*env.Source {

	var list = getTestSources(path.Join(tpath, testsFolderName), path.Join(npath, testsFolderName))
	if list == nil {
		env.AddProgramError("ОКР-ОШ-НЕТ-ТЕСТОВ", path.Join(npath, testsFolderName))
		return nil
	}

	return list
}

func getTestSources(tpath, npath string) []*env.Source {
	var err = env.EnsureFolder(npath)
	if err != nil {
		return nil
	}

	var list = env.GetFolderSources(tpath, npath)

	if len(list) == 0 {
		return nil
	}

	if len(list) == 1 && list[0].Err != nil {
		return nil
	}

	return list
}

// ==== тест вход

// Получает пути к тестируемому модулю
func prepareTestEntry(tpath, npath string) []*env.Source {

	var list = getTestEntrySource()
	if list == nil {
		panic("нет входа в тестирование")
		return nil
	}

	if len(list) != 1 {
		panic("тест-вход - должен быть один исходный файл")
	}

	var text = string(list[0].Bytes)

	text = strings.ReplaceAll(text, patternImportPath, tpath)
	text = strings.ReplaceAll(text, patternModuleName, filepath.Base(npath))

	//fmt.Printf("!! %v %v\n", npath, filepath.Base(npath))

	list[0].Bytes = []byte(text)

	return list
}

func getTestEntrySource() []*env.Source {

	env.Normalizer.Process(testEntryPath)
	if env.Normalizer.Err != nil {
		env.AddProgramError("ОКР-НЕТ-ТЕСТ-ВХОДА", testEntryPath, env.Normalizer.Err.Error())
		return nil
	}

	var npath = env.Normalizer.NPath

	var err = env.EnsureFolder(npath)
	if err != nil {
		return nil
	}

	var list = env.GetFolderSources(testEntryPath, npath)

	if len(list) == 0 {
		return nil
	}

	if len(list) == 1 && list[0].Err != nil {
		return nil
	}

	return list
}
