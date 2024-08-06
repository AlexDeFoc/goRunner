package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// AppInfo holds information about an application
type AppInfo struct {
	Name  string
	Path  string
	Alias string
}

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

func cmdApp() {
	// Get the path of the executing script
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the absolute path of the executing script
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the directory of the executing script
	scriptDir := filepath.Dir(absPath)

	// Get the location of File A
	configFilePath := filepath.Join(scriptDir, "config")

    configFile, err := os.ReadFile(configFilePath)

    if err != nil {
        fmt.Printf("Error finding config file: %v", err)
    }

	programFile := strings.TrimSpace(strings.ReplaceAll(string(configFile), `"`, ""))
    programFilePath := filepath.Join(scriptDir, programFile)

	// Get DATA from program file
	data, err := os.Open(programFilePath)
	if err != nil {
		fmt.Println("Error opening program file:", err)
		return
	}
	defer data.Close()

	// Read and process lines
	scanner := bufio.NewScanner(data)
	appsMap := make(map[string]AppInfo)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ", ", 3)
		if len(parts) != 3 {
			//fmt.Println("Invalid line format:", line) Annoying bug warning that always runs and IDK how to fix it to run when actually needed
			continue
		}

		// Remove `name:`, `path:`, and `alias:` prefixes
		name := strings.TrimPrefix(parts[0], "name:\"")
		path := strings.TrimPrefix(parts[1], "path:\"")
		alias := strings.TrimPrefix(parts[2], "alias:\"")

		// Remove trailing quotes
		name = strings.TrimSuffix(name, "\"")
		path = strings.TrimSuffix(path, "\"")
		alias = strings.TrimSuffix(alias, "\"")

		appInfo := AppInfo{
			Name:  name,
			Path:  path,
			Alias: alias,
		}

		appsMap[appInfo.Name] = appInfo
		appsMap[appInfo.Alias] = appInfo
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	// Ensure there are sufficient arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: appName -a passedArgument | appName -c '+' -a arg1 arg2 ... | appName -p 'prefix' -c '+' -a arg1 arg2 ...")
		return
	}

	// Process the arguments
	appArg := os.Args[1]

	// Check if the app exists in the map
	app, found := appsMap[appArg]
	if !found {
		fmt.Println("App not found:", appArg)
		return
	}

	// Variables to hold flags and their arguments
	var prefix, combineSymbol string
	argsMap := make(map[string][]string)
	var currentFlag string

	// Iterate over the arguments to collect flags and their values
	for i := 2; i < len(os.Args); i++ {
		arg := os.Args[i]
		if strings.HasPrefix(arg, "-") {
			currentFlag = arg
			argsMap[currentFlag] = []string{}
		} else if currentFlag != "" {
			argsMap[currentFlag] = append(argsMap[currentFlag], arg)
		} else {
			fmt.Printf("Unexpected argument: %s\n", arg)
			return
		}
	}

	// Assign collected arguments to the respective variables
	if vals, ok := argsMap["-p"]; ok && len(vals) == 1 {
		prefix = vals[0]
	}
	if vals, ok := argsMap["-c"]; ok && len(vals) == 1 {
		combineSymbol = vals[0]
	}
	arguments := argsMap["-a"]

	// Combine the arguments based on the prefix and combine symbol if provided
	var combinedArgument string
	if combineSymbol != "" {
		combinedArgument = prefix + strings.Join(arguments, combineSymbol)
	} else {
		combinedArgument = prefix + strings.Join(arguments, " ")
	}

	// Run the command with the path and the combined argument
    cmd := exec.Command(app.Path, combinedArgument)
    cmd.Dir = filepath.Dir(app.Path)
    cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Printf("Error starting command: %v", err)
	}
}

func runApp() {
	// Get the path of the executing script
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the absolute path of the executing script
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the directory of the executing script
	scriptDir := filepath.Dir(absPath)

	// Get the location of File A
	configFilePath := filepath.Join(scriptDir, "config")

    configFile, err := os.ReadFile(configFilePath)

    if err != nil {
        fmt.Printf("Error finding config file: %v", err)
    }

	programFile := strings.TrimSpace(strings.ReplaceAll(string(configFile), `"`, ""))
    programFilePath := filepath.Join(scriptDir, programFile)

	// Get DATA from program file
	data, err := os.Open(programFilePath)
	if err != nil {
		fmt.Println("Error opening program file:", err)
		return
	}
	defer data.Close()

	// Read and process lines
	scanner := bufio.NewScanner(data)
	appsMap := make(map[string]AppInfo)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ", ", 3)
		if len(parts) != 3 {
			//fmt.Println("Invalid line format:", line)
			continue
		}

		// Remove `name:`, `path:`, and `alias:` prefixes
		name := strings.TrimPrefix(parts[0], "name:\"")
		path := strings.TrimPrefix(parts[1], "path:\"")
		alias := strings.TrimPrefix(parts[2], "alias:\"")

		// Remove trailing quotes
		name = strings.TrimSuffix(name, "\"")
		path = strings.TrimSuffix(path, "\"")
		alias = strings.TrimSuffix(alias, "\"")

		appInfo := AppInfo{
			Name:  name,
			Path:  path,
			Alias: alias,
		}

		appsMap[appInfo.Name] = appInfo
		appsMap[appInfo.Alias] = appInfo
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
		return
	}

	// Process command-line arguments
	appsNotFound := []string{}
	for _, argValue := range os.Args[1:] {
		app, found := appsMap[argValue]
		if found {
            cmd := exec.Command(app.Path)
            cmd.Dir = filepath.Dir(app.Path)
            cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				log.Printf("Error starting command: %v", err)
			}
		} else {
			appsNotFound = append(appsNotFound, argValue)
		}
	}

	if len(appsNotFound) > 0 {
		fmt.Println("Apps not found:", appsNotFound)
	}
}

func main() {
    // Check if any flags are present
    flagFound := false
    for _, arg := range os.Args[1:] {
        if strings.HasPrefix(arg, "-") {
            flagFound = true
            break
        }
    }

    // Decide which function to run based on the presence of flags
    if flagFound {
        cmdApp()
    } else {
        runApp()
    }
}
