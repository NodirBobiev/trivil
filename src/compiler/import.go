package compiler

import (
	"fmt"

	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

func (cc *compileContext) importModule(i *ast.Import) {

	env.Normalizer.Process(i.Path)
	if env.Normalizer.Err != nil {
		env.AddError(i.Pos, "ОКР-ОШ-ПУТЬ-ИМПОРТА", i.Path, env.Normalizer.Err.Error())
		return
	}

	var npath = env.Normalizer.NPath

	m, ok := cc.imported[npath]
	if ok {
		if m == nil {
			env.AddError(i.Pos, "СЕМ-ЦИКЛ-ИМПОРТА", i.Path)
			return
		}

		// Модуль уже был импортирован
		i.Mod = m
		//fmt.Printf("already imported %s\n", i.Path)
		return
	}

	var err = env.EnsureFolder(npath)
	if err != nil {
		env.AddError(i.Pos, "ОКР-ИМПОРТ-НЕ-ПАПКА", npath, err.Error())
		return
	}

	var list = env.GetFolderSources(i.Path, npath)

	if len(list) == 0 {
		env.AddError(i.Pos, "ОКР-ИМПОРТ-ПУСТАЯ-ПАПКА", i.Path, npath)
		return
	}

	if len(list) == 1 && list[0].Err != nil {
		env.AddError(i.Pos, "ОКР-ОШ-ЧТЕНИЕ-ИСХОДНОГО", list[0].FilePath, list[0].Err.Error())
		return
	}

	if cc.testModulePath == i.Path {
		var tlist = collectTestSources(i.Path, npath)

		if env.ErrorCount() != 0 {
			return
		}

		list = append(list, tlist...)
	}

	cc.imported[npath] = nil // начали обработку
	i.Mod = cc.parseModule(false, list)
	cc.imported[npath] = i.Mod
}
