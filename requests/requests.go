package requests

cosnt (
	ACCEPT_ALL = "*"
	ACCEPT_MEDIA = "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"

	ACCEPT_LANGUAGE_JAPAN = "ja-JP,ja;q=0.9,en-US;q=0.8,en;q=0.7"

	USERAGENT_CHROME = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36"
)
func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", ACCEPT_ALL)
	req.Header.Set("Accept-Language", ACCEPT_LANGUAGE_JAPAN)
	req.Header.Set("User-Agent", USERAGENT_CHROME)

	return http.DefaultClient.Do(req)
}
