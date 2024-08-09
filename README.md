# About this app:
* Made completely in Go
* Extremely fast, in the miliseconds range and few ones too!
* Customisable app names (alias) and description for apps
* Give you the power to run any app quickly with aliases in your console on any platform
* Ease of integration with powershell/any other scripting language for infinite power and making you work easier

# Platforms supported:
- Windows
- Linux
- MacOS (called darwin in the releases files)

# Architectures supported:
- AMD64 : 64 bit cpu
- 386 : 32 bit cpu
- ARM : arm cpu
- ARM64 : arm 64 bit cpu

# How to install:
1. Go to [Releases page](https://github.com/AlexDeFoc/goRunner/releases)
2. Download the file for your platform and architecture
3. Extract files to a folder wherever you want

# How to setup
## Config file:
1. Open the config file
2. Type the name of your programs file in it

## Programs file:
1. Open the programs file (by default it's called _"apps.json"_)
2. Type your apps in it following this structure:
```
{
    "appName1": {
        "path": "path1/to/app/executable",
        "desc": "description or full name of the app or anything here"
    },
    "appName2": {
        "path": "example/path2/app2.exe",
        "desc": "description or full name of the app or anything here"
    },
    "appName3": {
        "path": "replace/this/path3/to/app.exe",
        "desc": "description or full name of the app or anything here"
    }
}
```
#### Note:
The paths need to have slashes like this :
* ✅ "path/to/exeFile"
* ❌ "path\to\exeFile"
* To see how a real apps.json file would look like, check mine! [Link to file](https://pastebin.com/iMjVTjYs)

# How to use
## Start multiple apps at once:
1. Type in your console your apps following this structure:
```
.\main.exe app1 app2 app3
```

## Start a single app with arguments
1. Remember this flags for arguments:
```
-n : name of app
-p : prefix
-c : how to connect the arguments toghether
-a : arguments will follow
```
2. Type in your console your command following these posible structures:
```
.\main.exe -n app1
```
Passing a single argument : very useful for passing a single argument for posible console/CLI apps OR going to a specific URL.

```
.\main.exe -n app1 -p urlPrefix -c "+" -a stuff1ToSearchFor stuff2ToAddToTheSearch
```
Passing a single argument : very useful for opening URLS and searching through SITES.

3. Get a list of your app names:
```
.\main.exe list
```

### NOTES: 
1. The values for the flags : _"-p"_ and _"-c"_ need always to be put in between double quotes. (almost)
2. Flags don't have any order that they need to by typed in.

# How to build/run from source
### NOTES: 
1. I do not recommend building from source because I belive the ones i distribute on the [Releases page](https://github.com/AlexDeFoc/goRunner/releases), are enough for most people.
2. Releases are kept fresh, only times i would first upload the code and not the release is when I have to modify the README.md file for utilising the new features.
3. You are required to have the following stuff:
    - git (for cloning the repository)
    - go (to build and run the script)

## Building from source:
1. Clone the repository
2. Open a console/terminal and go to the folder you cloned into:
3. Run the following command:
```
go build main.go
```

## Running from source:
1. Clone the repository
2. Open a console/terminal and go to the folder you cloned into:
3. Run the following command:
```
go run main.go ...
```
Instead of the dots above put arguments or flags, so follow the structures from the [How to use - chapter](#How-to-use).

# Examples
### Notes: 
1. The following examples are very modular.
2. The lower you go the more difficult they become.
3. The script part is made of powershell code. Later for bash/other shells.
4. The script part exists for ease of running apps quicker with less typing and without needing to go the the goRunner app's directory.
5. I do not use Ripcord.
6. All powershell scripts below need to be added in your powershell profile so it's loaded when you start your terminal.
7. TO see a real powershell profile like mine check this link : [Link to pastebin](https://pastebin.com/yi0djr0Z)

## Dictionary: 
1. Open two apps
2. Go to url on Chrome
3. Open WhatsApp on Chrome
4. Search OR Open Chrome
5. Search OR Open Youtube
6. Other examples

## Open two apps: Discord and Chrome at the same time
#### The command:
```
.\main.exe ch dis
```
In my programs file i assigned for the Chrome app the ch alias and for the discord app the dis alias:
```
{
    "ch": {
        "path": "C:/Program Files/Google/Chrome/Application/chrome.exe"
    },
    "dis": {
        "path": "C:/App/Ripcord/Ripcord.exe",
        "desc": "Discord"
    }
}
```
Notice: I don't need to add a description to app.

#### The powershell script:
```
function appGO {
    param (
        [Parameter(ValueFromRemainingArguments = $true)]
        [string[]]$app
    )

    & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' @app
}

set-alias r appGO
```
The _"param"_ part is necesary. It makes our function be able to get the stuff we type in the terminal.

The _"&"_ symbol enforces the function to run the following command.

At the end I SET an alias for the function. Now i can use the command _"r ch"_ to run chrome! (ch is my alias for chrome in the apps.json file)


## Go to url on Chrome
#### The command:
```
.\main.exe -n ch -a "epicgames.com"
```
Here I go to _"epicgames.com"_ site using the _"-a"_ flag for *arguments*

#### The powershell script:
```
function urlGO {
    param (
        [Parameter(ValueFromRemainingArguments = $true)]
        [string[]]$url
    )

    & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' -n ch -a @url
}

set-alias chu urlGO
```
Notice that i use the _"@"_ symbol instead of the _"$"_ symbol because it allows us use this command for other functions later like [Open WhatsApp on Chrome](##Open-WhatsApp-on-Chrome).

At the end I SET an alias for the function. Now i can use the command _"chu epicgames.com"_ to go to that site!

## Open WhatsApp on Chrome
#### The command:
```
.\main.exe -n ch -a web.whatsapp.com
```

#### The powershell script:
```
function whatsAppRun {
    urlGO -url 'web.whatsapp.com'
}

set-alias wh whatsAppGO
```
Notice: We used the _"urlGO"_ function from before, and used the flag _"-url"_ here, that was in the urlGO function _"@url"_.

At the end I SET an alias for the function. Now i can use the command _"wh"_ to open WhatsApp!

## Search OR open Chrome
#### The command:
```
.\main.exe -n ch -p "google.com/search?q=" -c "+" socks for programmers that go crazy cuz of open source
```
In this example i doing some things:
1. Added a prefix, which google uses for searching terms on google. (Got it from trying to find a pattern in the URL).
2. Added a _"connecting symbol"_, with the flag -c, with the value of "+", meaning that the terms will be glued toghether with a plus sign.
4. Searched for random thing about socks.

#### The powershell script:
```
function googleGO{
    param (
        [Parameter(ValueFromRemainingArguments = $true)]
        [string[]]$searchQuery
    )

    if ($searchQuery) {
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' -n ch -p "google.com/search?q=" -c "+" @searchQuery
    }
    else {
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' ch
    }
}

set-alias ch googleGO
```
Now i can simply go into my terminal and search on google like this: _"ch socks for programmers that go crazy cuz of open source"_.

Notice: That i added a if statement. IT checks if i provide a search term. If not, than i want just to simply open chrome.

Such simplicity!!!

🤩 At the tip of our finger tips!

## Search OR open Youtube:
#### The command:
```
.\main.exe -n ch -p "youtube.com/results?search_query=" -c "+" the primagen zero to lsp
```
Notice: It's almost identical to googleGO.

#### The powershell script:
```
function ytGO{
    param (
        [Parameter(ValueFromRemainingArguments = $true)]
        [string[]]$searchQuery
    )

    if ($searchQuery) {
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' -n ch -p "youtube.com/results?search_query=" -c "+" @searchQuery
    }
    else {
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' -n ch -a "youtube.com"
    }
}

set-alias ch googleGO
```
Now we can surf youtube or search for something from our terminal!

Notice: That i added a if statement. IT checks if i provide a search term. If not, than i want just to simply open youtube.

## Other examples

1. Open explorer in the current directory
#### The command:
```
.\main.exe -n xp -a .
```
Note: xp is refering to the name i gave explorer in my apps.json file.

#### The powershell script:
```
function explorerGO {
    & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' -n xp -a .
}

set-alias xp explorerGO
```
Now i can type: _"xp"_ and open file explorer in my current directory.


2. Start Hollow Knight with the following arguments "-screen-width 912 -screen-height 570"
#### The command:
```
.\main.exe -n hk -a "-screen-width 912 -screen-height 570"
```
Note: xp is refering to the name i gave explorer in my apps.json file.

#### The powershell script:
```
function hollowKnightGO {
    & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' -n hk -a '-screen-width 912 -screen-height 570'
}

set-alias hk hollowKnightGO
```
Now i can type: _"hk"_ and start hollow knight with those arguments so i can play at good fps on bad PC.

Note: This is the way to pass any argument to an app.

# Contribution - Open source:
Did you just notice we are on GitHub? An open source ussualy, source sharing site for devs like me and you, or users alike? That's crazy that this exists.

So now i will list what you can support/contribute to this reposity with:
1. Examples in general
2. Scripts for the examples in other scripting languages, other then powershell. Especially for the Linux community, and Mac too but they are a fork of Linux. haha lol, burning take.
3. Any optimisation for the scripts/examples.

## Notes from me:
What i wish for this project is to hold it's simplicity, blazing fast speed and ease of use, and also modularity and flexibility. But any suggestions are warmed welcomed. Thank you, have a nice day and thanks for reading!
