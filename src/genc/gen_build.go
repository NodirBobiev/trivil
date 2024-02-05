package genc

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	//	"path"
	"runtime"
	"strings"

	"trivil/ast"
	"trivil/env"
)

const (
	conf_file_name = "config/genc.txt"
	place_files    = "#files#"
	place_target   = "#target#"
	place_runtime  = "#runtime#"
	place_genc     = "#genc#"
)

var _ = fmt.Printf

func BuildExe(modules []*ast.Module) {
	//fmt.Printf("build: %s\n", runtime.GOOS)

	//=== setup command
	var command = findTemplate(runtime.GOOS + "-build")
	if command == "" {
		return
	}
	var names = make([]string, len(modules))
	for i, m := range modules {
		names[i] = env.OutName(m.Name) + ".c"
	}

	var target = env.OutName(modules[len(modules)-1].Name)

	command = strings.ReplaceAll(command, place_files, strings.Join(names, " "))
	command = strings.ReplaceAll(command, place_runtime, env.RuntimePath())
	command = strings.ReplaceAll(command, place_target, target)

	var folder = env.PrepareOutFolder()
	gencAbsolute, _ := filepath.Abs(folder)
	command = strings.ReplaceAll(command, place_genc, gencAbsolute)

	//=== write script file
	var script = findTemplate(runtime.GOOS + "-script")
	if script != "" {
		var lines = make([]string, 1)
		lines[0] = command

		writeFileExecutable(folder, script, "", lines)
	}

	var arg string
	var mainCmd string

	switch runtime.GOOS {
	case "windows":
		mainCmd = "cmd"
		arg = fmt.Sprintf("[/c cd %s & call %s ]", folder, script)
		//fmt.Printf("arg %v\n", arg)
	case "linux", "darwin", "freebsd":
		absoluteFolder, _ := filepath.Abs(folder)
		mainCmd = "bash"
		arg = path.Join(absoluteFolder, script)
	default:
		panic("build not implemented for " + runtime.GOOS)
	}

	var cmd = exec.Command(mainCmd, arg)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Exec failed: %s\n%s\n", err.Error(), string(out))
	} else {
		fmt.Printf("Execute: ./%s  Rebuild C code: %s/%s\n", target, folder, script)
	}
}

//=== genc configuration

var settings []string

func findTemplate(name string) string {

	if settings == nil {

		buf, err := os.ReadFile(env.SettingsRelativePath(conf_file_name))
		if err != nil {
			env.AddProgramError("ГЕН-ОШ-КОНФ-ФАЙЛА", err.Error())
			return ""
		}

		settings = strings.Split(string(buf[:]), "\n")
	}

	for _, s := range settings {
		if strings.HasPrefix(s, name) {
			var pair = strings.SplitN(s, ":", 2)
			if len(pair) == 2 && strings.TrimSpace(pair[1]) != "" {
				return strings.TrimSpace(pair[1])
			}
			env.AddProgramError("ГЕН-ОШ-НАСТРОЙКА", conf_file_name, name)
			return ""
		}
	}

	env.AddProgramError("ГЕН-ОШ-НАСТРОЙКА", conf_file_name, name)
	return ""
}

//====
