package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "os/exec"
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

func main(){
    cmdApp()
}

func cmdApp() {
	// Get the name of program file from config
	configFile, err := os.ReadFile("config")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	programFile := strings.ReplaceAll(string(configFile), `"`, "")
	programFile = strings.ReplaceAll(programFile, "\r\n", "")

	// Get DATA from program file
	data, err := os.Open(programFile)
	if err != nil {
		fmt.Println("Error opening program file:", err)
		return
	}
	defer data.Close()

	// Read and process lines
	lineScanner := bufio.NewScanner(data)
	appsMap := make(map[string]AppInfo)

	for lineScanner.Scan() {
		line := strings.ReplaceAll(lineScanner.Text(), `"`, "")
		parts := strings.SplitN(line, ", ", 3)
		if len(parts) != 3 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		appInfo := AppInfo{
			Name:  strings.TrimPrefix(parts[0], "name:"),
			Path:  strings.TrimPrefix(parts[1], "path:"),
			Alias: strings.TrimPrefix(parts[2], "alias:"),
		}

		appsMap[appInfo.Name] = appInfo
		appsMap[appInfo.Alias] = appInfo
	}

	if err := lineScanner.Err(); err != nil {
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
	var prefix string
	var combineSymbol string
	var arguments []string

	// Iterate over the arguments to collect flags and their values
	for i := 2; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "-a":
			i++
			for i < len(os.Args) && !strings.HasPrefix(os.Args[i], "-") {
				arguments = append(arguments, os.Args[i])
				i++
			}
			i--

		case "-c":
			i++
			if i < len(os.Args) {
				combineSymbol = os.Args[i]
			} else {
				fmt.Println("Usage: appName -c '+' -a arg1 arg2 ...")
				return
			}

		case "-p":
			i++
			if i < len(os.Args) {
				prefix = os.Args[i]
			} else {
				fmt.Println("Usage: appName -p 'prefix' -c '+' -a arg1 arg2 ...")
				return
			}

		default:
			fmt.Printf("Unknown flag: %s\n", os.Args[i])
			return
		}
	}

	// Combine the arguments based on the prefix and combine symbol if provided
	var combinedArgument string
	if combineSymbol != "" {
		combinedArgument = prefix + strings.Join(arguments, combineSymbol)
	} else {
		combinedArgument = prefix + strings.Join(arguments, " ")
	}

	// Run the command with the path and the combined argument
	cmd := exec.Command(app.Path, combinedArgument)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Printf("Error starting command: %v", err)
	}
}

func runApp() {

  // Get the name of program file from config
  configFile, err := os.ReadFile("config")
  if err != nil {
    fmt.Println("Error reading config file:", err)
    return
  }

  programFile := strings.ReplaceAll(string(configFile), `"`, "")
  programFile = strings.ReplaceAll(programFile, "\r\n", "")

  // Get DATA from program file
  data, err := os.Open(programFile)
  if err != nil {
    fmt.Println("Error opening program file:", err)
    return
  }
  defer data.Close()

  // Read and process lines
  lineScanner := bufio.NewScanner(data)
  appsMap := make(map[string]AppInfo)

  for lineScanner.Scan() {
    line := strings.ReplaceAll(lineScanner.Text(), `"`, "")
    parts := strings.SplitN(line, ", ", 3)
    if len(parts) != 3 {
      fmt.Println("Invalid line format:", line)
      continue
    }

    appInfo := AppInfo{
      Name:  strings.TrimPrefix(parts[0], "name:"),
      Path:  strings.TrimPrefix(parts[1], "path:"),
      Alias: strings.TrimPrefix(parts[2], "alias:"),
    }

    appsMap[appInfo.Name] = appInfo
    appsMap[appInfo.Alias] = appInfo
  }

  if err := lineScanner.Err(); err != nil {
    fmt.Println("Error reading lines:", err)
    return
  }

  // Process command-line arguments
  appsNotFound := []string{}
  for _, argValue := range os.Args[1:] {
    app, found := appsMap[argValue]
    if found {
      cmd := exec.Command(app.Path)
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
