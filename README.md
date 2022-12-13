# cookies

small command line tool for looking up chrome cookies for a domain on mac osx

## install

```
$ GO111MODULE=on go get -u github.com/nicksherron/cookies
```

## usage 

The default path thats checked for cookie db is ~/Library/Application Support/Google/Chrome/Default/Cookies
set "GOOGLE_CHROME_PROFILE" to your profile name if it's not "Default"
eg `export GOOGLE_CHROME_PROFILE="Profile 1"`
would check ~/Library/Application Support/Google/Chrome/Profile 1/Cookies

```
$ cookies github.com
_octo=GH1.1.1842523239.15; _device_id=68c3d0a0ce4cd510c51964e0bc267adc9c3; 
... 
```
