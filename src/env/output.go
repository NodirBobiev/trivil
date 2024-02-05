package env

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var _ = fmt.Printf

// Подготовка папку для записи кода
func PrepareOutFolder() string {

	name := "./_genc"

	err := os.Mkdir(name, 0750)
	if err != nil && !os.IsExist(err) {
		panic("ошибка создания папки: " + name)
	}

	return name
}

// Возвращает имя, заменяя русские буквы и пробелы
func OutName(name string) string {
	var b strings.Builder
	var s string
	var origin rune

	var c = func(x string) {
		if unicode.IsUpper(origin) {
			s = strings.ToUpper(x)
		} else {
			s = x
		}
	}

	for _, origin = range name {
		switch unicode.ToLower(origin) {
		case ' ', '-':
			s = "_"
		case '№':
			s = "N_"
		case '?':
			s = "Qm"
		case '!':
			s = "Em"
		case 'а':
			c("a")
		case 'б':
			c("b")
		case 'в':
			c("v")
		case 'г':
			c("g")
		case 'д':
			c("d")
		case 'е':
			c("e")
		case 'ё':
			c("e")
		case 'ж':
			c("zh")
		case 'з':
			c("z")
		case 'и':
			c("i")
		case 'й':
			c("yi")
		case 'к':
			c("k")
		case 'л':
			c("l")
		case 'м':
			c("m")
		case 'н':
			c("n")
		case 'о':
			c("o")
		case 'п':
			c("p")
		case 'р':
			c("r")
		case 'с':
			c("s")
		case 'т':
			c("t")
		case 'у':
			c("u")
		case 'ф':
			c("f")
		case 'х':
			c("x")
		case 'ц':
			c("tc")
		case 'ч':
			c("ch")
		case 'ш':
			c("sh")
		case 'щ':
			c("shch")
		case 'ъ':
			c("_")
		case 'ы':
			c("y")
		case 'ь':
			c("_")
		case 'э':
			c("e")
		case 'ю':
			c("yu")
		case 'я':
			c("ya")
		default:
			s = string(origin)
		}
		b.WriteString(s)
	}
	return b.String()
}
