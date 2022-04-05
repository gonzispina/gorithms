package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var packageName = flag.String("package", "", "Name of the package to create")

func main() {
	flag.Parse()

	if *packageName == "" {
		panic("package flag must be defined")
	}

	wd, _ := os.Getwd()
	*packageName = strings.ToLower(*packageName)
	packagePath := path.Join(wd, *packageName)
	if err := os.Mkdir(packagePath, os.ModePerm); err != nil {
		if os.IsExist(err) {
			panic(fmt.Sprintf("Package '%s' already exists", packagePath))
		}

		panic(err)
	}

	files := map[string]struct {
		fileName string
		lines    []string
	}{
		"main": {
			fileName: path.Join(packagePath, fmt.Sprintf("%s.go", *packageName)),
			lines: []string{
				fmt.Sprintf("package %s \n\n", *packageName),
				"func Solution(N int) int { \n\n",
				"} \n",
			},
		},
		"test": {
			fileName: path.Join(packagePath, fmt.Sprintf("%s_test.go", *packageName)),
			lines: []string{
				fmt.Sprintf("package %s_test\n\n", *packageName),
				"import (\n",
				fmt.Sprintf("	\"github.com/gonzispina/gorithms/%s\"\n", *packageName),
				"	\"github.com/stretchr/testify/assert\"\n",
				"	\"testing\"\n",
				")\n\n",
				"func TestSolution(t *testing.T) { \n",
				"	t.Run(\"Test Solution N= \", func(t *testing.T) {\n",
				fmt.Sprintf("		res := %s.Solution(0)\n", *packageName),
				"		assert.Equal(t, ,res)\n",
				"	})\n",
				"} \n",
			},
		},
		"readme": {
			fileName: path.Join(packagePath, "README.md"),
			lines: []string{
				fmt.Sprintf("# %s", *packageName),
			},
		},
	}

	for name, file := range files {
		f, err := os.Create(file.fileName)
		if err != nil {
			panic(fmt.Sprintf("Couldn't create %s file. Err %s", name, err))
		}

		for _, l := range file.lines {
			_, err = f.Write([]byte(l))
			if err != nil {
				return
			}
		}

		if err = f.Close(); err != nil {
			panic(fmt.Sprintf("Couldn't close %s file. Err %s", name, err))
		}
	}
}
