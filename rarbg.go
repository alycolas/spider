package main

import (
	"fmt"
	"net/http"
//	"bufio"
//	"os"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"flag"
)

// const domain = "https://rarbggo.org"
const domain = "https://rarbg.to"

var (
	key string
	order = "seeders"
	auto bool
)

func init(){
	flag.BoolVar(&auto, "a", false, "chrus the first magnet")
	flag.StringVar(&key, "k", "", "input keyworld for search")
}

func gethtml(url string, key string, order string) string{
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"/torrents.php", nil)
	if err != nil {
		fmt.Println("Cannot Get The Page")
	}

	if auto {
		order = "data"
	}

	q := req.URL.Query()
	q.Add("by", "DESC")
	q.Add("order", order)
	q.Add("search", key)
	q.Add("r", "97888647")
	req.URL.RawQuery = q.Encode()
//	fmt.Println(req.URL.String())

//	req.Header.Add("Host", "rarbgprx.org")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36 Edg/80.0.361.69")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Referer", "https://rarbggo.org/threat_defence.php?defence=2&sk=n0b72m4p3s&cid=30308730&i=1651319939&ref_cookie=rarbggo.org&r=43488754")
	req.Header.Add("Accept-Encoding", "utf8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Cookie", "skt=copwtaeumc; gaDts48g=q8h5pp9t; skt=copwtaeumc; gaDts48g=q8h5pp9t; tcc; aby=2; ppu_main_9ef78edf998c4df1e1636c9a474d9f47=1; ppu_sub_9ef78edf998c4df1e1636c9a474d9f47=3")

	resp ,err := client.Do(req)
	if err != nil {
		panic(err)
	}

//	// Print Test
//	defer resp.Body.Close()
//	html, _  := ioutil.ReadAll(resp.Body)
//	fmt.Println(os.Stdout, string(html))
//	// eng of Test

	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		panic(err)
	}

	nextUrl := dom.Find("tr.lista2 td.lista a[onmouseover]").Map(func(i int, sel *goquery.Selection) string {
		str, _ := sel.Attr("href")
		return str
	})
	title := dom.Find("tr.lista2 td.lista a[onmouseover]").Map(func(i int, sel *goquery.Selection) string {
		str := sel.Text()
		return str
	})
	size := dom.Find("tr.lista2 td.lista[width=\"100px\"]").Map(func(i int, sel *goquery.Selection) string {
		str := sel.Text()
		return str
	})

	if auto {
		return nextUrl[0]
	}

	for i := range title {
		fmt.Printf("%02d |%10s| %s\n", i, size[i], title[i])
	}

	var i int

	fmt.Println("Please input NUM")
	fmt.Scanln(&i)

	return nextUrl[i]
}


func getMagnet(domain string,url string) string {
	var magnet string
	client := &http.Client{}
	req, err := http.NewRequest("GET", domain+url, nil)
	if err != nil {
		fmt.Println("Cannot Get The Page")
	}

	q := req.URL.Query()
	q.Add("r", "97888647")
	req.URL.RawQuery = q.Encode()
//	fmt.Println(req.URL.String())

//	req.Header.Add("Host", "rarbgprx.org")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36 Edg/80.0.361.69")
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Sec-Fetch-Site", "same-origin")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Referer", "https://rarbggo.org/threat_defence.php?defence=2&sk=n0b72m4p3s&cid=30308730&i=1651319939&ref_cookie=rarbggo.org&r=43488754")
	req.Header.Add("Accept-Encoding", "utf8")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Add("Cookie", "skt=copwtaeumc; gaDts48g=q8h5pp9t; skt=copwtaeumc; gaDts48g=q8h5pp9t; tcc; aby=2; ppu_main_9ef78edf998c4df1e1636c9a474d9f47=1; ppu_sub_9ef78edf998c4df1e1636c9a474d9f47=3")

	resp ,err := client.Do(req)
	if err != nil {
		panic(err)
	}

	dom, err := goquery.NewDocumentFromResponse(resp)
	dom.Find("img[src=\"https://dyncdn.me/static/20/img/16x16/download.png\"]+a+a").Each(func(i int, sel *goquery.Selection) {
		mag, _ := sel.Attr("href")
		fmt.Println(mag)
		magnet = mag
	})
	return magnet
}

func main() {
	flag.Parse()
	url := gethtml(domain, key, order)
	magnet := getMagnet(domain, url)
	ioutil.WriteFile("/tmp/MAG", []byte(magnet), 0666)
}
