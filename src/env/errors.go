package env

import (
	"fmt"
	"os"
	"strings"
)

const errorsPath = "config/errors.txt"

type Error struct {
	id     string
	source *Source
	line   int
	col    int
	text   string
}

var (
	errors   []*Error
	messages map[string]string
)

func initErrors() {
	errors = make([]*Error, 0)
	messages = make(map[string]string)

	buf, err := os.ReadFile(SettingsRelativePath(errorsPath))
	if err != nil {
		fmt.Printf("! error reading %v file: %v\n", errorsPath, err.Error())
		return
	}

	var lines = strings.Split(string(buf[:]), "\n")

	for _, s := range lines {
		pair := strings.SplitN(s, ": ", 2)
		if len(pair) == 2 {
			messages[pair[0]] = strings.ReplaceAll(pair[1], "$;", "%v")
		}
	}
}

func PosString(pos int) string {
	source, line, col := SourcePos(pos)
	return fmt.Sprintf("%s:%d:%d", source.FilePath, line, col)
}

func AddError(pos int, id string, args ...interface{}) string {

	source, line, col := SourcePos(pos)

	var err = &Error{
		id:     id,
		source: source,
		line:   line,
		col:    col,
	}

	template, ok := messages[id]
	msg := ""

	if ok {
		msg = fmt.Sprintf(template, args...)
	} else {
		msg = fmt.Sprintf("сообщение для ошибки '%s' не задано!", id)
	}

	err.text = fmt.Sprintf("%s:%d:%d:%s: %s", source.FilePath, line, col, id, msg)

	errors = append(errors, err)

	return err.text
}

func AddProgramError(id string, args ...interface{}) string {

	var err = &Error{
		id:     id,
		source: nil,
		line:   0,
		col:    0,
	}

	template, ok := messages[id]
	msg := ""

	if ok {
		msg = fmt.Sprintf(template, args...)
	} else {
		msg = fmt.Sprintf("сообщение для ошибки '%s' не задано!", id)
	}

	err.text = fmt.Sprintf("%s: %s", id, msg)

	errors = append(errors, err)

	return err.text
}

func FatalError(id string, args ...interface{}) {
	AddProgramError(id, args...)
	ShowErrors()
	os.Exit(1)
}

func ShowErrors() {
	for _, e := range errors {
		fmt.Println(e.text)
	}
}

func ErrorCount() int {
	return len(errors)
}

func GetError(i int) string {
	return errors[i].text
}

func GetErrorId(i int) string {
	return errors[i].id
}

func ClearErrors() {
	errors = make([]*Error, 0)
}
