# goUtils

my golang utils. it can use command or package

```bash
go install github.com/awisu2/goUtil
```

## command help

```text
go 言語でのutil環境

Usage:
   [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  images      各種画像関係処理
  output      output the set value to stdout and stderr
  time        時間処理

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```

## apis

## goUtils/files

ファイル操作

- IsExists(filePath string) error
- Save(data []byte, savePath string) error
- Read(readPath string) ([]byte, error)

## goUtils/urls

- Copy(url url.URL) *url.URL
  - *url.URLをコピー(`u := *url.URL; copiedUrl := &u` と等価)

## goUtils/requests

httpリクエスト

- Get(url string) (*http.Response, error)
  - 取得後に`defer res.Body.Close()`が必要です
- GetBody(url string) ([]byte, error)
- Download(url string, filename string) error
  - ファイルのダウンロード

## goUtils/pathes

any path utils

- UserPathes: userDirectory. os specifiction path
- AppConfigDir:
- AppConfigPath:
- MakeSafePath: make safe path. if already exists chaneg path
- ReBaseName: name path's base without extension
- SplitName: splite base to name and extension
