GoMUD - a basic MUD server in Go 
================================
GoMUD is a WIP, currently it is a basic echo server. Basically it's a
project to assist us in learning the Go programming language, but
hopefully it will become a viable MUD base alternative.


Running GoMUD 
-------------
The following commands will download, build, install, and start GoMUD 

'git clone git@github.com:jfsherman/gomud.git' into your '$GOPATH'
'cd gomud/gmserver'
'go install'
'gmserver'

TODOs
-----
* Proper telnet protocol negotiation
* Database interface (SQLite?)
* Lua integration / online scripting language
* World system
* Lots of other stuff
