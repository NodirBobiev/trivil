package env

import (
	goerr "errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var _ = fmt.Printf

const (
	separator = "::"
	stdName   = "стд"
)

// Обработка входного пути для модуля:
// Поиск по кодовой базе или нормализация
type NormalizePath struct {
	Root  string // "", если не задана кодовая база
	Err   error  // ошибка при обработке кодовой базы или нормализации
	NPath string // Найденный или нормализованный путь
}

// Объект для поиска
var Normalizer NormalizePath

// Словарь кодовых баз
var sourceRoots = make(map[string]string)

//====

func (np *NormalizePath) Process(fpath string) {

	np.Root = ""
	np.Err = nil

	var parts = strings.SplitN(fpath, separator, 2)
	if len(parts) != 2 || len(parts) >= 1 && parts[0] == "" {

		np.NPath, np.Err = filepath.Abs(fpath)

		return
	}

	np.Root = parts[0]
	rootPath, ok := sourceRoots[np.Root]
	if !ok {
		np.Err = goerr.New(fmt.Sprintf("путь для кодовой базы '%s' не задан", np.Root))
		return
	}
	np.NPath = path.Join(rootPath, parts[1])
}

//====

var baseFolder = ""

// Возвращает ошибку, если путь не указывает на папку
func EnsureFolder(fpath string) error {
	fi, err := os.Stat(fpath)
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return goerr.New("это файл")
	}
	return nil
}

// Папка с настройками компилятора
func SettingsFolder() string {
	return baseFolder
}

func SettingsRelativePath(filename string) string {
	return path.Join(baseFolder, filename)
}

func RuntimePath() string {
	return path.Join(baseFolder, "runtime")
}

func initFolders() {
	var err error
	//var dir = filepath.Dir(os.Args[0])
	//baseFolder, err = filepath.Abs(dir)
	//if err != nil {
	//	panic(fmt.Sprintf("filepath.Abs(%s): %s", dir, err.Error()))
	//}
	//
	//baseFolder = filepath.ToSlash(baseFolder)
	baseFolder = "/home/cyrus/trivil"
	var stdPath = path.Join(baseFolder, stdName)

	err = EnsureFolder(stdPath)
	if err == nil {
		sourceRoots[stdName] = stdPath
	}

	// TODO: прочитать файл со списком кодовых баз
}
