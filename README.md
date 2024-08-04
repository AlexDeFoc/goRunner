# Short about this app:
* App made fully in go and distributed as executable for any platform. (batch is code is only for ME to build and compress the app for releasing)
* It's a very easy to use, flexible config customisation and protability app.
* It provides you the power to run any app on your computer in the console, allows you to set aliases for shorter names, and you can manually set alias in your shell for a command like "run" to run the main executable from a certain directory.

# Platforms supported:
- Windows
- Linux
- MacOS (called darwin)

# Architectures supported:
- AMD64 : for a cpu of 64 bit
- 386 : for a cpu of 32/86 bit
- ARM : for cpu of type arm
- ARM64 : for cpu of type arm64

# How to download released archive:
1. Go to releases
2. Choose your version
3. Download the archive with the platform-architecture name from the chapters above
4. De-archive/Extract the contents in a folder where you want your "goRunner" app to be

# How to use the app:
The program is separated in three parts:
1. Main executable:
* Run it with arguments (name OR alias of any app set in your programs list)
2. Config file:
* Here you write the name of the programs file list (i set it to bank)
3. Programs file:
* Here you can add as many apps you want. An app is composed of :
  * Name
  * Path
  * Alias
  * NOTE: A programs list called bank-demo is provided for a easier time customising + a config file containing the name "bank-demo"

# Customisation:
- This app allows you to set any name and alias for your app in the programs file so you can run apps easier & faster

# How it all works:
1. The main executable
- read the name of the programs file from config file
- read the apps data and separate them into sections for easier searching from the programs file
- if the user provided a argument (name or alias) find in the programs list it exists, and run the path containing the executable

# Building from source:
### 1. Building with the command (RECOMMENDED) (put the following commands):
- go build main.go

### 2. Running without building with the command (NOT RECOMMENDED - slower) (put the following commands):
- go run main.go

### 3. The build batch file (NOT RECOMMENDED):
- that script will create a folder called export which will contain a build folder for each platform

### 4. The compress batch file (NOT RECOMMENDED):
- that script will compress each folder inside the export folder into a folder called archive containing the archived build folders for each platform 
