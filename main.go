package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var packageName = flag.String("package", "", "Name of the package to create")
var functionName = flag.String("function", "", "Name of the function")

func main() {
	flag.Parse()

	if *packageName == "" {
		panic("package flag must be defined")
	}

	if *functionName == "" {
		*functionName = "solution"
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
				fmt.Sprintf("func %s(N int) int { \n\n", *functionName),
				"} \n",
			},
		},
		"test": {
			fileName: path.Join(packagePath, fmt.Sprintf("%s_test.go", *packageName)),
			lines: []string{
				fmt.Sprintf("package %s\n\n", *packageName),
				"import (\n",
				"	\"github.com/stretchr/testify/assert\"\n",
				"	\"testing\"\n",
				")\n\n",
				"func TestSolution(t *testing.T) { \n",
				"	t.Run(\"Test Solution\", func(t *testing.T) {\n",
				fmt.Sprintf("		res := %s(0)\n", *functionName),
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
