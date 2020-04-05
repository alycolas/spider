package main

import (
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"flag"
	"regexp"
	"strconv"
	"io/ioutil"
)

const domain = "https://mikanani.me/Home/Bangumi/"

var (
	key string
	auto bool
	magnet =  ""
	name string
)

func init(){
	flag.BoolVar(&auto, "a", false, "chrus the first magnet")
	flag.StringVar(&key, "k", "2170", "input keyworld for search")
	flag.StringVar(&name, "d", "", "input keyworld for search")
}

func getpage (domain string,key string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", domain+key, nil)
	if err != nil {
		fmt.Println("Cannot Get The Page")
	}

	resp ,err := client.Do(req)
	if err != nil {
		panic(err)
	}

	dom, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		panic(err)
	}

	title := dom.Find("a.magnet-link-wrap").Map(func(i int, sel *goquery.Selection) string {
		str := sel.Text()
		return str
	})
	mlist := dom.Find("a.magnet-link").Map(func(i int, sel *goquery.Selection) string {
		str, _ := sel.Attr("data-clipboard-text")
		return str
	})

	var E int
	epsode, err := ioutil.ReadFile("/home/tiny/dm/"+name+"/Epsode")
	if err != nil {
		E = 0
	} else {
		tmp, _ := strconv.Atoi(string(epsode))
		E = tmp
	}

	for i := range title {
		Epsode := []string{"", "0"}
		reg := regexp.MustCompile(`^.*[\[第E 【]([\d]{2})[\]话 】].*$`)
		if reg.MatchString(title[i]) {
			Epsode = reg.FindStringSubmatch(title[i])
		}
		// fmt.Println(Epsode)
		e, _ := strconv.Atoi(Epsode[1])
		// fmt.Println(e)
		if e > E {
			err := ioutil.WriteFile("/home/tiny/dm/"+name+"/Epsode", []byte(Epsode[1]), 0644)
			if err != nil {
				fmt.Println("haha")
			}
			magnet = mlist[i]
			E = e
		}
	//	fmt.Printf("%02d | %70s | %s\n", i, title[i], mlist[i])
	}
	fmt.Println(magnet)
}

func main() {
	flag.Parse()
	getpage(domain, key)
}
