package main

import (
	"errors"
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

	// commands
	app.Commands = []cli.Command{

		//generate code
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

				generateCode(fileName, prefix)

				return nil
			},
		},
	}

	app.Run(os.Args)

}

// generate code for the specify file with prefix -> {public,private,protected}.
func generateCode(fileName, prefix string) error {
	filePath := path.Join(getAbsulutePath(), fileName)
	file, err := os.OpenFile(filePath, os.O_RDWR, 0777)
	checkErr(err)

	// read data with specific len data
	data := make([]byte, 1024)

	// for searching variable's name
	//startVar := "$"
	//endVar := ";"
	//endClass := "}"
	//foundVar := false

	// getting the varible name from file
	//var vName []byte
	// hold string getter and setter result from all varible
	//var template string

	// reading file's data to store to our variable => data
	n, err := file.Read(data)
	checkErr(err)
	if n == 0 {
		return errors.New("Empty selected file")
	}

	// get variable's name => $varNames
	// and the last index of => '}' from our data
	varNames, indexEndClass := getVarNames(data)
	var code string

	// generate our code
	code = constructor(varNames)
	for _, varName := range varNames {
		code += getSetterGetter(prefix, varName)
	}
	code += "\n}"

	// write generated code to file

	_, err = file.WriteAt([]byte(code), int64(indexEndClass))
	checkErr(err)

	//continue to find variable's name.
	//for index, char := range data {
	//	// finding the varible
	//	switch char {
	//	case startVar[0]:
	//		// write template getter setter
	//		if vName != nil {
	//			if len(template) == 0 {
	//				template = getSetterGetter(prefix, string(vName))
	//			} else {
	//				template = template + "\n" + getSetterGetter(prefix, string(vName))
	//			}
	//		}

	//		// reset vName to empty byte, because we found new variable
	//		vName = nil
	//		foundVar = true

	//		// avoid $ being added to vName
	//		continue
	//	case endVar[0]:
	//		foundVar = false
	//	case endClass[0]:
	//		// we call this again for our last variable
	//		if vName != nil {
	//			if len(template) == 0 {
	//				template = getSetterGetter(prefix, string(vName))
	//			} else {
	//				template = template + "\n" + getSetterGetter(prefix, string(vName))
	//			}
	//		}
	//		template = template + "\n}"
	//		// insert all getSet before the end of data
	//		_, err := file.WriteAt([]byte(template), int64(index))
	//		checkErr(err)
	//		break
	//	}

	//	// getting the varible name.
	//	if foundVar {
	//		vName = append(vName, char)
	//	}

	//}

	fmt.Println("Succes generate code")
	return nil

}

// getting all variable names in the classs.
// return those names and the indexEndClass --> '}'
func getVarNames(datas []byte) ([]string, int) {
	startVar := "$"
	endVar := ";"
	endClass := "}"
	foundVar := false

	var indexEndClass int

	// getting the varible name from file
	var vName []byte
	var allName []string

	for index, data := range datas {
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
			indexEndClass = index
			break
		}

		if foundVar {
			vName = append(vName, data)
		}

	}

	return allName, indexEndClass
}

// Create constructor
func constructor(arguments []string) string {
	space := "    "
	var result string

	result = space + "public function __construct(" + getArgments(arguments) + "){\n" +
		setConstructor(arguments) + space + "}\n"
	return result
}

// Create $his->variable for the constructor
func setConstructor(arguments []string) string {
	space := "    "
	var result string
	for _, argument := range arguments {
		result += space + space + "$this->" + argument + "=$" + argument + ";\n"
	}
	return result
}

// create arguments for constructor
func getArgments(arguments []string) string {
	var result string
	for index, argument := range arguments {
		if index != len(arguments)-1 {
			result += "$" + argument + ","
		} else {
			result += "$" + argument // last argument without comma.
		}
	}
	return result
}

// get The current absolute path on the active shell
func getAbsulutePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
	}

	return dir
}

// Creating getter and setter template. with prefix is the first modifier such as public or private or protected.
// name is the name of the function
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
