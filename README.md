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
### 1. Main executable:
## Starting multiple apps at once
   - this is the default behaviour if for example you have two apps inside the programs file and they are called ch and dis (first one being an alias for chrome and seconds one for discord)
   - how you would run it is simply type in your console/shell: main.exe ch dis
   - TIP: I recommend you do the same as me and set a alias for the app like so: (Powershell script in my case but i recommend you do it in your case too for easier call of apps) - don't forget to make your alias setting so it accepts/parses/passes multiple arguments
  
     ![image](https://github.com/user-attachments/assets/a85a16df-ecc9-4a63-bc86-b2fb8c9e8178)

## Starting a SINGLE APP but with multiple arguments passed to it - this is the juicy part
- **MY EXAMPLE**: So i want to run chrome and go to a certain website like youtube and on youtube search smth that i type in console
1. To obtain the information I need, I manually visit a YouTube URL. I found that the URL has this prefix: youtube.com/results?search_query=, and I observed that each search term is joined by a plus sign.
2. Setting the alias: I will go to my powershell script and only slightly modifying the above script (which is loaded into my Windows Terminal config when i load it up), and adding a similar function that runs the following complete command :
3. & '.\main.exe' "ch" -p "youtube.com/results?search_query=" -c "+" -a $query
4. And setting the name of the alias to be "yt"

![image](https://github.com/user-attachments/assets/891a95e4-b56e-4505-a3d4-bda0bc845370)

6. Let's decompose the above command:
 - Contains the main.exe file
 - the "ch" part is our app alias for chrome, cuz i will be running chrome
 - the "-p" flag means we are parsing a prefix to the app and putting the prefix in double quotes
 - the "-c" flag means what we are going to connect with toghether, the multiple arguments we will pass
 - the "-a" flag means we are going to put now our arguments
 - the $query part is just from the powershell script giving the arguments we will pass to the command
 - also don't forget the space after the "-a" flag cuz we want to have the flag and argument separated
 - NOTE: the flags dont have a certain order so put them in any way you want!
7. Let's search for example on youtube Michael Jackson, we will type simply in our terminal "yt Michael Jackson" and TaDaa! We got what we intended from our terminal!

![image](https://github.com/user-attachments/assets/97931f6f-70d8-4f0f-a13b-775b293e8d88) - URL
![image](https://github.com/user-attachments/assets/9acd6d2b-b59f-43b9-a2c5-20bda18705f5) - Search bar

### 2. Config file:
* Here you write the name of the programs file list (personally i like to name the programs file to "bank" like a bank of apps)
### 3. Programs file:
* Here you can add as many apps you want. An app is composed of :
  * Name
  * Path
  * Alias
  * NOTE: A programs list called bank-demo is provided for a easier time customising + a config file containing the name "bank-demo"
 
# Structure of programs file list:
![image](https://github.com/user-attachments/assets/39954499-db1f-4be7-8ce5-5679660aef34)

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
