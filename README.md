iGOtpwned
=========
cli tool to interact with https://haveibeenpwned.com API

Compiling
=========
first get the dependency
```
go get github.com/codegangsta/cli
```

then build the codez
```
go build
```

run
```
./iGOtpwned -h

NAME:
   iGOtpwned - 'Have I been pwned?' golang cli checker

USAGE:
   iGOtpwned [global options] command [command options] [arguments...]

VERSION:
   0.0.1

AUTHOR(S): 
   Ye Myat Kaung (Maverick) <mavjs01@gmail.com> 
   
COMMANDS:
   email, m email address to look up all breaches associated with it
   site, s  info associated with a single breached site
   help, h  Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h       show help
   --version, -v    print the version
```
