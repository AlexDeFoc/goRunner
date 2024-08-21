## About  
- Utility app that enables you to launch any app/game/custom command on your computer by setting custom aliases.  
- It is blazingly fast and lightweight.  
- It is very customisable, even allowing you to set aliases to other aliases with added overridden fields.  
- Cross-platform & Cross-Architecture (Windows, MacOS, Linux | x64, x86, arm, arm64)
<br><br>

## Installation  
### Normal installation  
- This allows you to use the app in your terminal right after installing it. (Adds automatically to your %PATH%).  
#### Guide for Windows:  
1. Go to the [Releases](https://github.com/AlexDeFoc/goRunner/releases/latest) page.  
2. Download the executable for your cpu architecture. Most probably it's 64 or 32 bit. It will look like this: _win_cpu_architecture.exe_. If you don't understand the names of the architectures provided check this footnote.[^1]  
[^1]: amd64 = 64 bit cpu, i386 = 32 bit cpu
3. Double-Click it and follow the steps of the installer.  
4. Done! Now you can enter your terminal and start by typing firstly the app name which is _"ru"_, and then the app or alias name.  
#### Guide for Linux/MacOS:  
1. Go to the [Releases](https://github.com/AlexDeFoc/goRunner/releases/latest) page.  
2. Download the bash script with the _".sh"_ extension for your architecture.  
3. Run the script with sudo/administrator privileges to be sure it installs correctlly.  
4. Done! Now you can enter your terminal and start by typing firstly the app name which is _"ru"_, and then the app or alias name.  
### Portable installation  
- This doesn't allow you to use the app for anywhere on your computer, because it is not added in your enviromental variables / %PATH%. So you will have to do it manually or run the app while the folder containing the app.  
1. Go to the [Releases](https://github.com/AlexDeFoc/goRunner/releases/latest) page.  
2. Download the archive file with the _".tar.xz"_ extension for your operating system and architecture.  
3. Extract using the command _"tar -xvJf"_ or other ways you extract the files from a compressed tarball.  
4. Done! Now you can use it by entering your terminal and changing the directory to the app's one and then starting by typing the app name which is _"ru"_, and then the app or alias name.  
<br><br>

## Configuration  
###### Note: To see my app & alias configs check in the repository the folder called "My setup".
### Configuration fields structure:  
- Advanced mode  
- Shell  
###### Shell field is optional and only necesarry when creating aliases that need to run a system command.  
###### Advanced mode is by default false, meaning the app will allow you to only run multiple apps at once and no aliases. To be able to use aliases you must enable Advanced Mode.  
###### Check the available shells in the footnote: [^2]  
###### Reason of having hard coded of shells is here: [^3]  
[^2]: cmd (Command prompt), pwsh (powershell 7), bash  
[^3]: The app handles different shells by using shells prefixes: _"cmd /C"_, _"pwsh -Command"_, "_bash -c_".  
### App fields structure:  
- name  
  - path  
  - desc  
###### Descriptions are optional.  
###### Paths must contain only right slashes: _"/"_.  
###### [Examples here](#app-configuration-chapters)  
### Alias fields structure:  
- name  
  - app  
  - status  
  - prefix  
  - connect  
  - fallback  
  - command  
  - desc  
###### Each field is optional, but some are necesary to create great aliases.  
###### [Examples here](#aliases-types)  
### Alias fields values:
- name : name you will be using the call the alias
  - app : name of the app which is from the app.json file
  - status : 
    - "search" - for adding arguments from the terminal to this alias
    - "pass" - for passing a certain hard-coded command with arguments to an app
    - "call" - used to call other aliases + having the ability to add fields with different values to override the called aliases fields
  - prefix : text to be added in front of a final command being passed to an app
  - connect : symbol/character or text that will be put in between the arguments provided in the terminal
  - command : used when wanting to pass a certain command that may contain flags with dashes and values, to an app
  - fallback : used when no arguments are provided in the terminal, and a different value should be passed to the app or alias fallback field
  - desc : this is completely optional, aswell as the other ones, with it's only purpose is to add a bit of description to the alias because most likely the alias's name is comprised of just a few letters

<br><br>


## Examples  
### App configuration chapters:  
- [Open an app](#open-an-app)  
- [Open a game](#open-a-game)  
- [Open a system app](#open-a-system-app)  
#### Examples:  
##### Open an app  
```
"ch": {
    "path": "C:/Program Files/Google/Chrome/Application/chrome.exe"
}
```
##### Open a game  
```
"hk": {
    "path": "C:/Game/Internet/Hollow Knight/hollow_knight.exe",
    "desc": "Hollow Knight"
}
```
##### Open a system app  
```
"xp": {
    "path": "explorer.exe"
}
```
<br><br>

### Aliases Types:  
- Run a command with arguments [&#x2139;&#xFE0F;](#run-a-command-with-arguments)
- Start a game with flags and arguments [&#x2139;&#xFE0F;](#start-a-game-with-flags-and-arguments)
- Start a game that is launched with a custom protocol URL/URI (Steam/Epic Games/Google Play Games) [&#x2139;&#xFE0F;](#start-a-game-that-is-launched-with-a-custom-protocol-urluri)
- Search stuff on the internet in your browser AND having the ability to not search for anything and just open the browser on a custom page or default one [&#x2139;&#xFE0F;](#search-stuff-on-the-internet-in-your-browser)
- Search stuff on youtube in your browser AND having the ability to not search for anything and just open youtube. [&#x2139;&#xFE0F;](#search-stuff-on-youtube-in-your-browser)
- Search for a certain link in your browser (Needs the [Search stuff on internet](#search-stuff-on-the-internet-in-your-browser) example) [&#x2139;&#xFE0F;](#run-a-command-with-arguments)
- Open a certain site (Needs the [Search for a certain link](#search-for-a-certain-link-in-your-browser) example) [&#x2139;&#xFE0F;](#search-for-a-certain-link-in-your-browser)


#### Examples:
##### Run a command with arguments:
```
"xp": {
    "app": "xp",
    "status": "command",
    "command": "start explorer.exe ."
}
```
##### Start a game with flags and arguments:
```
"hk": {
    "app": "hk",
    "status": "pass",
    "command": "-screen-width 912 -screen-height 570"
}
```
##### Start a game that is launched with a custom protocol URL/URI:
```
"coc": {
    "app": "coc",
    "status": "command",
    "command": "start googleplaygames://launch/?id=com.supercell.clashofclans&lid=1&pid=1",
    "desc": "Start clash of clans"
}
```
##### Search stuff on the internet in your browser:
```
"ch": {
    "app": "ch",
    "status": "search",
    "prefix": "google.com/search?q=",
    "connect": "+",
    "fallback": "google.com"
}
```
##### Search stuff on youtube in your browser:
```
"yt": {
    "app": "ch",
    "status": "search",
    "prefix": "youtube.com/results?search_query=",
    "connect": "+",
    "fallback": "youtube.com"
}
```
##### Search for a certain link in your browser:
```
"chu": {
    "app": "ch",
    "status": "call",
    "prefix": "",
    "connect": "",
    "fallback": "",
    "desc": "Search web for a link"
}
```
##### Open a certain site:
```
"wh": {
    "app": "chu",
    "status": "call",
    "fallback": "web.whatsapp.com"
}
```
```
"epic": {
    "app": "chu",
    "status": "call",
    "fallback": "epicgames.com"
}
```
<br><br>

## Contribution:
### Parts that need improvement:
- Additional shells with their apropriate prefix
- Additional alias examples
