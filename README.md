# YCode 

Wrapper for easy copying yubikey oath codes. The yubico authenticator keeps freezing up and cmd line is just nice.

Provide copying and pasting to the Clipboard for Go.

## Build:
```
$ go get github.com/schans/ycode
```

## Platforms:
* OSX
* Windows
* Linux, Unix (requires 'xclip' or 'xsel' command to be installed)

## Requirements
Have ykman installed: https://developers.yubico.com/yubikey-manager/
```
$ brew install ykman
```

## Run:
```
$ ycode
```

## Disclaimer:
Code is what it is: lazy and lousy. It works on my machine though :D

Example:
<img src="/example.gif?raw=true"/>

