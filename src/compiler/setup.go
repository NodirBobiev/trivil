package compiler

import (
	"fmt"

	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

func (cc *compileContext) setup(setuped *ast.Module) []*ast.Module {

	var setting = setuped.Setting

	env.Normalizer.Process(setting.Path)
	if env.Normalizer.Err != nil {
		env.AddError(setting.Pos, "ОКР-ОШ-ПУТЬ-ИМПОРТА", setting.Path, env.Normalizer.Err.Error())
		return nil
	}

	var npath = env.Normalizer.NPath

	var err = env.EnsureFolder(npath)
	if err != nil {
		env.AddError(setting.Pos, "ОКР-НЕ-ПАПКА", npath, err.Error())
		return nil
	}

	var list = env.GetFolderSources(setting.Path, npath)

	if len(list) == 0 {
		// TODO: изменить ошибку
		env.AddError(setting.Pos, "ОКР-ИМПОРТ-ПУСТАЯ-ПАПКА", setting.Path)
		return nil
	}

	if len(list) == 1 && list[0].Err != nil {
		env.AddError(setting.Pos, "ОКР-ОШ-ЧТЕНИЕ-ИСХОДНОГО", list[0].FilePath, list[0].Err.Error())
		return nil
	}

	var mods = cc.parseList(false, list)

	for _, m := range mods {
		m.Name = setuped.Name
	}

	return mods
}
