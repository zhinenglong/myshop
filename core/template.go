package core

import (
	"github.com/gin-contrib/multitemplate"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func TempleFunc() template.FuncMap {
	return template.FuncMap{}
}

func LoadTemplates(template_dir string) multitemplate.Renderer {

	basedir := template_dir
	layoutdir := basedir + "/layout/**"
	itemsdir := basedir + "/items/**"
	pagedir := basedir + "/page/**"
	render := multitemplate.NewRenderer()
	files, err := filepath.Glob(template_dir + "/**")
	if err != nil {
		panic(err.Error())
	}
	for _, file := range files {
		file_type, err := os.Stat(file)
		if err != nil {
			panic(err.Error())
		}
		if !file_type.IsDir() {
			filename := filepath.Base(file)
			tmpinfo := strings.Split(filename, ".")
			fileslayout := GetPathFiles(layoutdir, tmpinfo[0])
			filespages := GetPathFiles(pagedir, tmpinfo[0])
			filesitems := GetPathFiles(itemsdir, tmpinfo[0])
			if len(filespages) > 0 {
				if len(fileslayout) > 0 {
					filespages = append(filespages, fileslayout...)
					if len(filesitems) > 0 {
						filespages = append(filespages, filesitems...)
					}
				} else {
					if len(filesitems) > 0 {
						filespages = append(filespages, filesitems...)
					}
				}
			} else {
				if len(fileslayout) > 0 {
					if len(filesitems) > 0 {
						fileslayout = append(fileslayout, filesitems...)
					}
					filespages = fileslayout
				} else {
					if len(filesitems) > 0 {
						filespages = filesitems
					}
				}
			}

			filespages = append([]string{file}, filespages...)
			render.AddFromFilesFuncs(filename, TempleFunc(), filespages...)
		}
	}
	basedir = ""
	layoutdir = ""
	itemsdir = ""
	pagedir = ""
	return render
}

func GetPathFiles(currentpath string, prefix string) []string {
	itemfiles, err := filepath.Glob(currentpath)
	if err != nil {
		panic(err.Error())
	}
	var files = make([]string, 0)
	for _, itemfile := range itemfiles {
		filestat, err := os.Stat(itemfile)
		if err != nil {
			panic(err.Error())
		}
		if !filestat.IsDir() {
			if strings.HasPrefix(filepath.Base(itemfile), prefix) {
				files = append(files, itemfile)

			}
		}
	}
	return files
}
