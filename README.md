# minibin

the simplest minibin for Windows 10/11 written in Go

<img src="./preview.png">

## installation

1. go to [releases](https://github.com/ktamkyaz/minibin/releases) and download `minibin.exe`
2. move file to **Startup** folder ( `WIN + R` -> `shell:startup` )
3. run `minibin.exe`

## build

```sh
go build -ldflags "-H=windowsgui" -o minibin.exe .
```
