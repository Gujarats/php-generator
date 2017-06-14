package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "PHP CODE GENERATOR"
	app.Usage = "Generate code from variables in php files"
	app.Version = "1.0.0"

	// flags for option command
	var prefix string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "prefix",
			Value:       "public",
			Usage:       "specify prefix modifier",
			Destination: &prefix,
		},
	}

	app.Commands = []cli.Command{

		//first command load column data
		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate getter and setter",
			Action: func(c *cli.Context) error {
				fileName := c.Args().Get(0)
				if fileName == "" {
					fmt.Println("Please specify fileName")
					return nil
				}

				generateSetGet(fileName, prefix)

				return nil
			},
		},
	}

	app.Run(os.Args)

}

func generateSetGet(fileName, prefix string) {
	// path
	filePath := path.Join(getAbsulutePath(), fileName)
	// open file
	file, err := os.OpenFile(filePath, os.O_RDWR, 0777)
	checkErr(err)

	// read data with specific len data
	data := make([]byte, 1024)

	startVar := "$"
	endVar := ";"
	endClass := "}"
	foundVar := false

	// getting the varible name from file
	var vName []byte
	// hold string getter and setter result from all varible
	var template string

	// finding the class string
	n, err := file.Read(data)
	checkErr(err)
	if n == 0 {
		return
	}

	// contine to find variable's name.
	for index, char := range data {
		// finding the varible
		if char == startVar[0] {
			// write template getter setter
			if vName != nil {
				if len(template) == 0 {
					template = getSetterGetter(prefix, string(vName))
				} else {
					template = template + "\n" + getSetterGetter(prefix, string(vName))
				}
			}

			// reset vName to empty byte, because we found new variable
			vName = nil
			foundVar = true

			// avoid $ being added to vName
			continue
		}

		// end of variable
		if char == endVar[0] {
			foundVar = false
		}

		// getting the varible name.
		if foundVar {
			vName = append(vName, char)
		}

		// write all template to a file
		if char == endClass[0] {
			// we call this again for our last variable
			if vName != nil {
				if len(template) == 0 {
					template = getSetterGetter(prefix, string(vName))
				} else {
					template = template + "\n" + getSetterGetter(prefix, string(vName))
				}
			}
			template = template + "\n}"
			// insert all getSet before the end of data
			_, err := file.WriteAt([]byte(template), int64(index))
			checkErr(err)
		}

	}

	fmt.Println("Succes generate getter and setter")

}

func getVarNames(datas []byte) []string {
	startVar := "$"
	endVar := ";"
	endClass := "}"
	foundVar := false

	// getting the varible name from file
	var vName []byte
	var allName []string

	for _, data := range datas {
		switch data {
		case startVar[0]:
			if vName != nil {
				allName = append(allName, string(vName))
			}
			foundVar = true
			vName = nil
			continue
		case endVar[0]:
			foundVar = false
		case endClass[0]:
			if vName != nil {
				allName = append(allName, string(vName))
			}
			vName = nil
			break
		}

		if foundVar {
			vName = append(vName, data)
		}

	}

	return allName
}

// get The current absolute path on the active shell
func getAbsulutePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}

	return dir
}

// Creating getter and setter template.
func getSetterGetter(prefix, name string) string {
	space := "    "
	// change the firsChar to UpperCase
	upper := strings.ToUpper(string(name[0]))[0]
	Upname := string(upper) + string(name[1:])

	getTemplate := space + prefix + " function get" + Upname + "(){\n" +
		space + space + "return $this->" + name + ";\n" + space + "}\n"

	setTemplate := space + prefix + " function set" + Upname + "($" + name + "){\n" +
		space + space + "$this->" + name + "=$" + name + ";\n" + space + "}\n"

	result := getTemplate + setTemplate
	return result

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
