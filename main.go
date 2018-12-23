package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/mushroomsir/iconv"
  "github.com/PuerkitoBio/goquery"
)

func ExampleScrape(w http.ResponseWriter, r *http.Request) {
  // Request the HTML page.
  res, err := http.Get("https://www.biqukan.com//0_200/22744555.html")
  if err != nil {
    log.Fatal(err)
  }
  defer res.Body.Close()
  if res.StatusCode != 200 {
    log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
  }

  utfBody, err := iconv.Convert(res.Body, iconv.GBK, iconv.UTF8)
	if err != nil {
		log.Fatal(err)
    // handler error
	}
  // Load the HTML document
  doc, err := goquery.NewDocumentFromReader(utfBody)
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  doc.Find("dd").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    band := s.Find("a").Text()
    links, _ := s.Find("a").First().Attr("href")
    url := "https://www.biqukan.com/" + links
    //title := s.Find("i").Text()
    fmt.Printf("%s\n %s\n %s", band, url)
  })
  content :=doc.Find("div#content").Text()
  w.Header().Add("Content-Type", "application/json")
  w.Write([]byte(content))
  fmt.Printf("%s",content)
}

func main() {
	addr := ":8181"
        http.HandleFunc("/", ExampleScrape)

        log.Println("listening on", addr)
        log.Fatal(http.ListenAndServe(addr, nil))

}
