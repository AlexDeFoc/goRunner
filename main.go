package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"

	//"log"
	"os"
	//"os/exec"
	"strings"
	"time"
)

// SplitAt splits a string at a delimiter
func SplitAt(delim string) bufio.SplitFunc {
				return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
								if atEOF && len(data) == 0 {
												return 0, nil, nil
								}
								if i := strings.Index(string(data), delim); i >= 0 {
												return i + len(delim), data[0:i], nil
								}
								if atEOF {
												return len(data), data, nil
								}
								return 0, nil, nil
				}
}

func main() {
				//Get the name of program file from config
				configFile, err := os.ReadFile("config")
				if err != nil {
								fmt.Println("Error reading config file:", err)
								return
				}

				programFile := string(configFile)
				programFile = strings.ReplaceAll(programFile, `"`, "")
				programFile = strings.ReplaceAll(programFile, "\r\n", "")

				//Get DATA from program file
				data, err := os.Open(programFile)
				if err != nil {
								fmt.Println("Error opening program file:", err)
								return
				}
				defer data.Close()

				// Separate lines
				lineScanner := bufio.NewScanner(data)
				var lines []string
				for lineScanner.Scan() {
								lines = append(lines, lineScanner.Text())
				}

				// Clean lines
				for i := range lines {
								lines[i] = strings.ReplaceAll(lines[i], "\"", "")
				}

				// Separate fields
				field := make([][]string, len(lines))
				for i := range field {
								field[i] = make([]string, 3)
				}

				for lineIndex, line := range lines {
								appScanner := bufio.NewScanner(strings.NewReader(line))
								appScanner.Split(SplitAt(", "))

								fieldCount := 0
								for appScanner.Scan() {
												if fieldCount < 3 {
																field[lineIndex][fieldCount] = appScanner.Text()
																fieldCount++
												}
								}
				}

				// Clean fields
				fieldName := map[int]string{
								0: "name:",
								1: "path:",
								2: "alias:",
				}

				for lineCount := range lines {
								for fieldIndex := range field[lineCount] {
												if strings.Contains(field[lineCount][fieldIndex], fieldName[fieldIndex]) {
																field[lineCount][fieldIndex] = strings.ReplaceAll(field[lineCount][fieldIndex], fieldName[fieldIndex], "")
												}
								}
				}

				fmt.Printf("Number of fields: %v\n", len(field))
				fmt.Printf("Number of lines: %v\n", len(lines))

				// Debug print the fields
				for i := range field {
								fmt.Printf("Field %d: %v\n", i, field[i][0])
				}

				nameField := 0
				pathField := 1
				aliasField := 2

/*				appsNotFound := []string{}
				t := time.Now()
				for _, argValue := range os.Args[1:] {
								found := false
								for lineIndex := range field {
												if argValue == field[lineIndex][aliasField] || argValue == field[lineIndex][nameField] {
																fmt.Printf("Index: %v got Value: %v\n", lineIndex, argValue)
																cmd := exec.Command(field[lineIndex][pathField])
																err := cmd.Start()
																if err != nil {
																				log.Printf("err")
																}
																found = true
																break
												}
								}
								if !found {
												appsNotFound = append(appsNotFound, argValue)
								}
				}


				if len(appsNotFound) > 0 {
								fmt.Println("Apps not found:", appsNotFound)
				}

				fmt.Printf("For loop took: %v", time.Since(t).Nanoseconds())

*/

				t2 := time.Now()
				// Create a map for fast lookups
				appMap := make(map[string][]string)
				for _, app := range field {
								appMap[app[nameField]] = app
								appMap[app[aliasField]] = app
				}

				apps2NotFound := []string{}

				for _, argValue := range os.Args[1:] {
								app, found := appMap[argValue]
								if found {
												fmt.Printf("Got Value: %v\n", argValue)
												cmd := exec.Command(app[pathField])
												err := cmd.Start()
												if err != nil {
																log.Printf("err")
												}
								} else {
												apps2NotFound = append(apps2NotFound, argValue)
								}
				}


				if len(apps2NotFound) > 0 {
								fmt.Println("Apps not found:", apps2NotFound)
				}
				fmt.Printf("For loop took: %v", time.Since(t2).Nanoseconds())


}









/*



func rum(){
				//CHECK if User Inputed at least one argument
				configFile, _ := os.ReadFile("config")
				programFile := string(configFile)
				programFile = strings.ReplaceAll(programFile, `"`, "")
				programFile = strings.ReplaceAll(programFile, "\r\n", "")

				f, _ := os.Open(programFile)

				//Separate lines
				lineScanner := bufio.NewScanner(f)
				line := []string{}
				for lineScanner.Scan() {
								line = append(line, lineScanner.Text())
				}

				//Clean lines
				for n := range len(line) {
								line[n] = strings.ReplaceAll(line[n], "\"", "")
								//fmt.Printf("%v\n", line[n]) //DEBUG
				}
				
				//Separate fields
				//fieldScanner := bufio.NewScanner(strings.NewReader(line[0]))
				//field := make([][]string, len(line))
				//appScanner.Split(SplitAt(", "))
				appScanner := make([]bufio.Scanner, len(line))

				for lineCount := range len(line) {
								appScanner[lineCount] = *bufio.NewScanner(strings.NewReader(line[lineCount]))
				}

				/*for lineCount := range len(line) {
								appScanner[lineCount].Split(SplitAt(", "))
								for appScanner[lineCount].Scan() {
												field[0] = append(field[lineCount], appScanner[lineCount].Text())
								}
				}*/
/*
				field := make([][]string, 3)
				for lineIndex := range len(line){
								for fieldIndex := range len(field){
												appScanner[lineIndex].Split(SplitAt(", "))
												for appScanner[lineIndex].Scan() {
																field[fieldIndex][lineIndex] = append(field[fieldIndex][lineIndex], appScanner[lineIndex].Text())
												}
								}
				}

				fmt.Printf("%v", len(field))
				fmt.Printf("%v", len(line))

/*
				//Clean fields
				fieldName := map[int]string {
								0 : "name:",
								1 : "path:",
								2 : "alias:",
				}

				for lineCount := range len(line) {
								for fieldIndex := range len(field) {
												if strings.Contains(field[lineCount][fieldIndex], fieldName[fieldIndex]) {
																field[lineCount][fieldIndex] = strings.ReplaceAll(field[lineCount][fieldIndex], fieldName[fieldIndex], "")
												}
								}
				}




				nameField := 0
				pathField := 1
				aliasField := 2

				appsNotFound := []string{}

				for _, argValue := range os.Args[1:] {
								found := false
								for lineIndex := range field {
												if argValue == field[lineIndex][aliasField] || argValue == field[lineIndex][nameField] {
																//fmt.Printf("Index: %v got Value: %v\n", lineIndex, argValue)
																cmd := exec.Command(field[lineIndex][pathField])
																err := cmd.Start()
																if err != nil {
																				log.Panic(err)
																}
																found = true
																break
												}
								}
								if !found {
												appsNotFound = append(appsNotFound, argValue)
								}
				}

				if len(appsNotFound) > 0 {
								fmt.Println("Apps not found:", appsNotFound)
				}
/*


func run() {
				rawConfig, err := os.ReadFile("config")

				// Error Handling
				if err != nil {
								log.Panic("Programs file not found")
				}

				config := string(rawConfig)
				config = strings.TrimSpace(config)

				rawFile, err := os.Open(config)
				defer rawFile.Close()

				// Error Handling
				if err != nil {
								log.Panic("Programs file not found")
				}

				lineScanner := bufio.NewScanner(rawFile)
				lineData := []string{}

				for lineScanner.Scan() {
								lineData = append(lineData, lineScanner.Text())
				}

				fieldData := make([][]string, len(lineData))

				for i, line := range lineData {
								// Process each line manually
								fields := strings.Split(line, ",")
								for j, field := range fields {
												field = strings.TrimSpace(field)
												field = strings.TrimSuffix(field, ",")
												if strings.Contains(field, "name:") {
																field = strings.ReplaceAll(field, "name:", "")
												} else if strings.Contains(field, "path:") {
																field = strings.ReplaceAll(field, "path:", "")
																field = strings.ReplaceAll(field, "\\", "/")
												} else if strings.Contains(field, "alias:") {
																field = strings.ReplaceAll(field, "alias:", "")
												}
												field = strings.Trim(field, `\"`)
												fields[j] = field
												//fmt.Println(field) //DEBUG
								}
								fieldData[i] = fields
				}

				// Get user flags
				var panicStatus bool
				for _, arg := range os.Args[1:] {
								for n := range fieldData {
												if arg != fieldData[n][2] && arg != fieldData[n][0] {
																cmd := exec.Command(fieldData[n][1])
																err := cmd.Start()
																//fmt.Printf("Going to run: %v\n", fieldData[n][1]) // DEBUG

																if err != nil {
																				log.Printf("Couldn't start: %v at path: %v\n", fieldData[n][0], fieldData[n][1])
																				//log.Printf("ERR: %v", err) //DEBUG
																				panicStatus = true
																}
												}
								}
				}

				if panicStatus == true {
								log.Panic()
				}
}
*/
