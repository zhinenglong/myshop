package tests

import (
	tfunc "myshop/core"
	"myshop/core/utils"
	"path/filepath"
	"testing"
)

func Test_template_function(t *testing.T) {
	cpath, _ := filepath.Abs(".")
	cpath = filepath.Dir(cpath)
	cpath = cpath + "/template"
	tfunc.LoadTemplates(cpath)
}

func Test_RootPath(T *testing.T) {
	utils.GetRootPath()
}
