package rollbar

import (
	"golang.org/x/mod/modfile"
	"io/ioutil"
)

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		println("Unable to get module name")
		return ""
	}

	modName := modfile.ModulePath(goModBytes)

	return modName
}
