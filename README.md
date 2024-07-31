# Benchmarking: 
why i choose to use maps instead of slices

Slice: for loop approach
: It's slower for a small batch of data and a bit quicker on large data.
&emsp;

Map: for loop approach
: It's slower for a small batch of data and a bit quicker on large data.
---

# Parameters of the tests:
1. NVIM [^1]
2. SHELL.go [^2]
3. SHELL.exe [^3]

# Speed notices:[^5]
### Small data (10 elements): 
* slice is slower by 6,5576 miliseconds or 31.23%
* map is faster by 23.8%

### Big data (100 elements):
* slice:
  * NVIM: 1,097 ms or 3.97% faster
  * SHELL.go: 2,741 ms or 7.14% faster
  * SHELL.go: 0,126 ms or 0.371% faster


* map is faster by 23.8%
  * NVIM: 1,097 ms or 4.13% slower
  * SHELL.go: 2,741 ms or 7.7% slower
  * SHELL.go: 0,126 ms or 0.373% slower


[^1]: Means i tested while in neovim with the following command: !go run main.go ap21 ap79 app100
[^2]: Means i tested while in powershell with the following command: go run main.go ap21 ap79 app100
[^3]: Means i tested while in powershell (after building the package, with go build main.go) with the following command: .\main.exe ap21 ap79 app100

# Size of data: [^4]
Test 1 - small data : 3 tries
Test 2 - big data : 12 tries
[^4]: Link to screenshot of spreadsheet: https://imgur.com/a/fQc9YV4
[^5]: Link to pastebin of raw speed results: https://pastebin.com/gqT3j7GJ
---
# Slice for loop code:
```
appsNotFound := []string{}
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
```
---
# Map for loop code:
```
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
```
---
