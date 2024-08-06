# About this app:
* Made completely in Go
* Extremely fast, in the miliseconds range and few ones too!
* Customisable aliases and names for apps
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
1. Open the programs file (by default it's called _"bank-demo"_)
2. Type your apps in it following this structure:
```
name:"AppName1", path:"path\to\your\app1\executable", alias:"optionalAlias1"
name:"AppName2", path:"path\to\your\app2\executable", alias:"optionalAlias2"
name:"AppName3", path:"path\to\your\app3\executable", alias:"optionalAlias3"
name:"AppName4", path:"path\to\your\app4\executable", alias:"optionalAlias4"
```

# How to use
## Start multiple apps at once:
1. Type in your console your apps following this structure:
```
.\main.exe app1 app2 app3
```

## Start a single app with arguments
1. Remember this flags for arguments:
```
-p : prefix
-c : how to connect the arguments toghether
-a : arguments will follow
```
2. Type in your console your command following these posible structures:
```
.\main.exe -a arg1
```
Passing a single argument : very useful for passing a single argument for posible console/CLI apps OR going to a specific URL.

```
.\main.exe -p urlPrefix -c "+" -a stuff1ToSearchFor stuff2ToAddToTheSearch
```
Passing a single argument : very useful for opening URLS and searching through SITES.

### NOTES: 
1. For the -p and -c flag you MUST put the stuff after them in double quotes (not sure about other ones).
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
3. Any contribution of other code for other shells like bash and others are very welcome because i try to keep this at maximum compatibility with other platforms like Linux.
4. The following examples contain the actual code you would be running AND a powershell script i use for automating and shortening the time i spend running commands.
5. I use Chrome as my default browser and the commands containing it are not limited to it only
6. I DO NOT USE RIPCORD as my discord client, which is against their Terms of Service. This is not a joke. And i do not endorse the use of any 3rd party clients.

## Dictionary: 
1. Open two apps
2. Go to url on Chrome
3. Open WhatsApp on Chrome
4. Search OR Open Chrome
5. Search OR Open Youtube

## Open two apps: Discord and Chrome at the same time
#### The command:
```
.\main.exe ch dis
```
In my programs file i assigned for the Chrome app the ch alias and for the discord app the dis alias:
```
name:"Ripcord", path:"C:\App\Test Rip Cord\Ripcord\Ripcord.exe", alias:"dis"
name:"chrome", path:"C:\Program Files\Google\Chrome\Application\chrome.exe", alias:"ch"
```
I may use Ripcord cuz my current pc is a potato :)), but i want to announce for the Discord team that this is an actual joke, and just a example for educational purposes and that i don't use any 3rd party client which is agains your TOS.

#### The powershell script:
```
function appsRun {
    param (
        [Parameter(ValueFromRemainingArguments = $true)]
        [string[]]$app
    )

    & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' $app
}

set-alias -Name r -Value appsGO
```
The param part of the code will appear in each example because it gives me a string array containing all the arguments i would pass in the console/terminal.

The _&_ symbol enforces the function to run the following command.

At the end i set an alias for the function and now i can simply, from anywhere on my computer do: _r dis ch_ , to open these apps!


## Go to url on Chrome
#### The command:
```
.\main.exe ch -a epicgames.com
```
In this example i open chrome and pass the epicgames.com url link.

#### The powershell script:
```
function urlGO {
    param (
        [Parameter(ValueFromRemainingArguments = $true)]
        [string[]]$url
    )

    if ($url) {
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' 'ch' -a @url
    }else{
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' 'ch'
    }
}

set-alias -Name chu -Value urlGO
```
Why does it look complicated? Because i wanted to be able to open chrome without typing any arguments and when i want to go to a certain url i want to use the same function.

The if statement checks if the url array is empty and if it's not then we pass the url as an argument.

Notice that i use the @ symbol instead of the $ symbol because it allows us use this command for other functions later like [Open WhatsApp on Chrome](##Open-WhatsApp-on-Chrome).

At the end i set an alias for the function and now i can simply, from anywhere on my computer do: _chu_ , to open chrome! or do: _chu epicgames.com_ , to go there!

## Open WhatsApp on Chrome
#### The command:
```
.\main.exe ch -a web.whatsapp.com
```

#### The powershell script:
```
function whatsAppRun {
    urlGO -url 'web.whatsapp.com'
}

set-alias -Name wh -Value whatsAppGO
```
Notice that we used the urlGo function from the previous example and parsing the "-url" flag that we had in the urlGO function put in it's command as "@url".

This way now we can finally use the urlGO for making aliases to open sites from inside our console instantly!

At the end i set an alias for the function and now i can simply, from anywhere on my computer do: _wh_ , to open WhatsAppWeb! Because my laptop is a potato and the desktop app is a memory hog.

## Search OR open Chrome
#### The command:
```
.\main.exe ch -p "google.com/search?q=" -c "+" -a socks for programmers that go crazy cuz of open source
```
In this example i doing some things:
1. Using the app aliased ch which is Chrome.
2. Provided a prefix for my command that I got from trying to search on google something and found a repeating pattern which is my prefix now. Notice i put it in quotes, it's needed i belive.
3. Provided a "argument combination symbol", which just means i saw again in the url of google when searching that each term that you search for is _connected_ by plus symbols.
4. Provided a search query which is just: socks for programmers that go crazy cuz of open source. :))
5. Notice that the more complex commands we want to achieve the more we would write, that's why we NEED scripts like the one below, or your own for your own shell

#### The powershell script:
```
function googleGO {
    param (
        [Parameter(ValueFromRemainingArguments = $true)]
        [string[]]$searchQuery
    )

    if ($searchQuery) {
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' 'ch' -p "google.com/search?q=" -c "+" -a @searchQuery
    }else{
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' 'ch'
    }
}

set-alias -Name ch -Value googleGO
```
Can you belive it? Now notice the alias. Now i can simply go into my terminal and do this: _ch socks for programmers that go crazy cuz of open source_ .

Such simplicity!!!

🤩 At the tip of our finger tips!

And at the same time with SUCH great speeds!

Also some technical stuff now - Notice that we follow the same trick as before, doing a if statement and if no argument/search terms provided than we are going to just open chrome, so now i can also just type _ch_ in my terminal and open Chrome!

## Search OR open Youtube:
#### The command:
```
.\main.exe ch -p "youtube.com/results?search_query=" -c "+" -a the primagen zero to lsp
```
Notice it's almost a complete copy of the previous example, this is only a coincidence because Google and Youtube are the same thing and they made it similar.

In this example i provide a prefix, connecting symbol and some terms to search for on youtube.

#### The powershell script:
```
function ytGO {
    param (
        [Parameter(ValueFromRemainingArguments = $true)]
        [string[]]$searchQuery
    )

    if ($searchQuery) {
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' 'ch' -p "youtube.com/results?search_query=" -c "+" -a @searchQuery
    }else{
        & 'C:/Users/Alefan/Documents/Powershell/go/main.exe' 'ch' -a "youtube.com"
    }
}

set-alias -Name yt -Value ytGO
```
Notice now that we can surf youtube from the terminal with a simple alias: _yt_ !

Now i can simply open my computer, my terminal and check the latest ThePrimagen Video by typing quickly: yt primagen new javascript framework meme .

Just joking ;), don't forget i got a life too, just tuning in sometimes.

Now some technical stuff - Notice that we have another if statement to know if we provide any terms to search for or just open youtube.

# Contribution - Open source:
Did you just notice we are on GitHub? An open source ussualy, source sharing site for devs like me and you, or users alike? That's crazy that this exists.

So now i will list what you can support/contribute to this reposity with:
1. Examples in general
2. Scripts for the examples in other scripting languages, other then powershell. Especially for the Linux community, and Mac too but they are a fork of Linux. haha lol, burning take.
3. Any optimisation for the main code or how we handle the programs file, maybe we can switch to something like hashtables/maps? I don't even know what are they. Still learning Golang.
4. Any optimisation for the scripts/examples.

## Notes from me:
What i wish for this project is to hold it's simplicity, blazing fast speed and ease of use, and also modularity and flexibility. But any suggestions are warmed welcomed. Thank you, have a nice day and thanks for reading!
