# goUtils

go関係で、簡易的な補助関数をまとめて利用できるように


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
