package genc

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"

	"trivil/ast"
	"trivil/env"
)

var _ = fmt.Printf

type genContext struct {
	module      *ast.Module
	outname     string
	declNames   map[ast.Decl]string
	genTypes    bool
	autoNo      int // used for additional names
	header      []string
	code        []string
	globals     []string
	init        []string        // используется для типов
	initGlobals []ast.Decl      // глобалы, которые надо инициализировать во входе
	sysAPI      map[string]bool // true - exported
}

func Generate(m *ast.Module, main bool) {

	var genc = &genContext{
		module:      m,
		outname:     env.OutName(m.Name),
		declNames:   make(map[ast.Decl]string),
		header:      make([]string, 0),
		code:        make([]string, 0),
		globals:     make([]string, 0),
		init:        make([]string, 0),
		initGlobals: make([]ast.Decl, 0),
		sysAPI:      make(map[string]bool),
	}

	genc.genModule(main)
	genc.finishCode()

	//genc.show()
	genc.save()

}

func (genc *genContext) h(format string, args ...interface{}) {
	genc.header = append(genc.header, fmt.Sprintf(format, args...))
}

func (genc *genContext) c(format string, args ...interface{}) {
	genc.code = append(genc.code, fmt.Sprintf(format, args...))
}

func (genc *genContext) g(format string, args ...interface{}) {
	genc.globals = append(genc.globals, fmt.Sprintf(format, args...))
}

func (genc *genContext) finishCode() {
	var hname = fmt.Sprintf("_%s_H", genc.outname)

	// header file
	var lines = genc.header

	genc.header = make([]string, 0)
	genc.h("#ifndef %s", hname)
	genc.h("#define %s", hname)

	genc.includeSysAPI(genc.header, true)
	genc.h("")

	genc.header = append(genc.header, lines...)

	genc.h("#endif")

	// code
	lines = genc.code
	genc.code = make([]string, 0)

	genc.c("#include \"rt_api.h\"")
	genc.c("#include \"%s\"", genc.outname+".h")
	genc.includeSysAPI(genc.code, false)
	genc.c("")

	if len(genc.globals) != 0 {
		genc.c("//--- globals")
		genc.code = append(genc.code, genc.globals...)
		genc.c("//--- end globals")
		genc.c("")
	}

	genc.code = append(genc.code, lines...)
}

func (genc *genContext) includeSysAPI(list []string, header bool) {
	for name, exported := range genc.sysAPI {
		if header == exported {
			if header {
				genc.h("#include \"%s.h\"", name)
			} else {
				genc.c("#include \"%s.h\"", name)
			}
		}
	}
}

//====

func (genc *genContext) save() {
	var folder = env.PrepareOutFolder()

	writeFile(folder, genc.outname, ".h", genc.header)
	writeFile(folder, genc.outname, ".c", genc.code)
}

func writeFile(folder, name, ext string, lines []string) {
    writeFileCommon(folder, name, ext, lines, 0644)
}

func writeFileExecutable(folder, name, ext string, lines []string) {
    writeFileCommon(folder, name, ext, lines, 0755)
}

func writeFileCommon(folder, name, ext string, lines []string, perm fs.FileMode) {

	var filename = path.Join(folder, name+ext)

	var out = strings.Join(lines, "\n")

	var err = os.WriteFile(filename, []byte(out), perm)

	if err != nil {
		panic("Ошибка записи файла " + filename + ": " + err.Error())
	}
}

//====

func (genc *genContext) show() {
	fmt.Println("---- header ----")
	for _, s := range genc.header {
		fmt.Println(s)
	}
	fmt.Println("---- c code ----")
	for _, s := range genc.code {
		fmt.Println(s)
	}
	fmt.Println("---- end c  ----")
}
