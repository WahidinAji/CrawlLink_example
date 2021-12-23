package main

import (
	"crawl/ssr"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type Link struct {
	url string
	title string
}

// TODO: Write your function here
func CrawlLink(url string) []Link {
	/**
	var links = []Link{
		{ url: "#", title: "Navbar" },
		{ url: "#", title: "Dropdown" },
		{
			url: "https://www.example.com/bridge/bait.html",
			title: "Crawl This URL",
		},
		{ url: "http://www.example.com/", title: "This URL Too" },
		{ url: "https://www.example.org/", title: "Ok Ganbatte" },
		{
			url: "https://www.example.org/bear/attraction.html",
			title: "More",
		},
		{ url: "https://example.com/bridge/bell.html", title: "Again" },
		{ url: "https://www.example.org/", title: "Chotto" },
		{ url: "http://example.com/", title: "Good Luck" },
		{
			url: "http://www.example.com/adjustment.php?branch=branch&birth=berry",
			title: "Hufft",
		},
		{
			url: "http://example.com/adjustment.php",
			title: "Ok I'm Getting Bored",
		},
		{ url: "https://www.example.edu/#bear", title: "hmmmmm" },
		{ url: "https://www.example.com/", title: "Ok Boomer" },
		{
			url: "http://www.example.com/believe/acoustics.html",
			title: "Demo Sonna Ja Dame",
		},
		{
			url: "http://apparel.example.com/blade.html",
			title: "Mo Sonna ja hora",
		},
		{
			url: "http://www.example.com/addition/bomb.php",
			title: "Kokoro wa shinkai suru yo",
		},
		{
			url: "https://www.example.com/?birthday=bird",
			title: "Motto Motto",
		},
		{
			url: "http://www.example.com/bait/bag?breath=animal",
			title: "Teret tereret",
		},
		{ url: "http://www.example.com/", title: "Sasageyo Jare" },
		{ url: "https://www.example.com/", title: "Lorem" },
		{ url: "https://bear.example.com/", title: "Ipsum" },
		{ url: "https://www.example.edu/", title: "Dolor" },
		{
			url: "http://example.com/?appliance=actor&arithmetic=battle",
			title: "Sit",
		},
		{ url: "http://bird.example.net/addition/blade.html", title: "Amet" },
		{
			url: "https://www.example.com/brick/animal.html",
			title: "Constrector",
		},
		{ url: "https://www.example.com/belief/breath", title: "sapien" },
		{
			url: "http://example.com/arch?afternoon=bait&acoustics=boy",
			title: "molli",
		},
		{ url: "https://www.example.net/afternoon", title: "auctor" },
		{ url: "https://example.com/", title: "lacinia" },
		{ url: "https://www.example.com/", title: "vitae" },
		{ url: "https://www.example.com/#basketball", title: "Ut magna" },
		{
			url: "https://bag.example.com/bed/blade",
			title: "Bang jual seblak gak",
		},
		{ url: "https://example.com/amount/badge#airport", title: "sodales" },
		{ url: "https://www.example.edu/", title: "vel" },
		{
			url: "http://example.com/brick?blood=arithmetic&brick=baseball",
			title: "diam",
		},
		{ url: "https://brick.example.net/anger", title: "cursus" },
		{ url: "https://birth.example.net/board.php", title: "ante" },
		{ url: "https://amount.example.com/", title: "Sed" },
		{
			url: "https://arm.example.com/?account=approval&arm=achiever#airport",
			title: "faucibus",
		},
		{ url: "http://example.com/bath.php", title: "Nam commodo" },
		{ url: "https://www.example.net/bottle#anger", title: "Tellus" },
		{ url: "http://www.example.com/#brake", title: "Alliquam" },
		{ url: "https://example.com/berry", title: "Iri" },
		{
			url: "https://www.example.net/board/aunt.aspx#anger",
			title: "bilang bos",
		},
		{ url: "http://example.com/", title: "sodales" },
		{ url: "https://example.edu/berry/badge.html", title: "tortor" },
		{ url: "http://example.net/", title: "Nunc" },
		{
			url: "http://www.example.net/?anger=agreement&border=boot",
			title: "elefiend",
		},
		{ url: "http://www.example.net/", title: "tempor" },
		{ url: "https://example.com/", title: "nibh" },
		{ url: "https://book.example.com/", title: "Jual" },
		{ url: "https://example.com/advice#brick", title: "Tanah Kampus" },
		{ url: "http://www.example.com/", title: "Stikom banyuwangi" },
		{ url: "https://example.com/arm.html", title: "Murah" },
		{ url: "https://bath.example.com/", title: "Minat PM" },
		{ url: "https://example.com/", title: "Donec" },
		{ url: "http://www.example.com/anger/balance", title: "Turpis" },
		{ url: "https://aftermath.example.com/bikes", title: "Condimentum" },
		{ url: "http://example.com/", title: "Tristique" },
		{ url: "https://www.example.com/aunt", title: "auctor" },
		{ url: "http://www.example.com/bead#bath", title: "justo" },
		{ url: "http://bed.example.com/bell.html", title: "Vehucula" },
		{ url: "https://aunt.example.net/brother", title: "Arcu" },
		{ url: "https://www.example.com/box/brass", title: "Integer" },
		{ url: "http://bikes.example.com/", title: "pal pale" },
		{ url: "https://example.com/bell/adjustment.aspx", title: "This" },
		{ url: "https://www.example.com/", title: "Is" },
		{ url: "http://www.example.com/", title: "Just" },
		{ url: "http://example.com/", title: "For" },
		{ url: "https://example.edu/advertisement", title: "Test Case" },
		{ url: "http://www.example.com/arch.aspx", title: "Crawllo" },
		{ url: "https://www.example.com/books/bridge.php", title: "Dama" },
		{ url: "http://example.com/", title: "Genera" },
		{ url: "https://www.example.com/brake.html", title: "Salamat" },
		{ url: "http://www.example.com/?amusement=back", title: "Malesuada" },
		{ url: "https://bedroom.example.com/", title: "Brotha" },
		{ url: "https://www.example.com/angle", title: "Morbui" },
		{
			url: "https://example.com/art?brass=box&bells=bird",
			title: "England",
		},
		{
			url: "http://attraction.example.org/brother.aspx",
			title: "Brotherrr",
		},
		{ url: "https://example.com/anger.html", title: "Anger" },
		{ url: "http://example.com/?achiever=bath", title: "Achievement" },
		{ url: "http://www.example.org/", title: "ORG" },
		{ url: "https://example.com/art", title: "ART" },
		{ url: "http://www.example.com/", title: "Def" },
		{ url: "https://anger.example.com/advertisement.htm", title: "Ads" },
		{
			url: "http://www.example.com/arch/attraction.php?bat=blade",
			title: "Attraction",
		},
		{ url: "http://example.com/back.php", title: "Back" },
		{ url: "https://example.com/aftermath/attack", title: "Attack" },
		{ url: "http://brake.example.edu/", title: "Edu" },
		{ url: "http://example.org/", title: "Again" },
	}
	 */

	res, err := http.Get("https://csr-assessment.miqdadyyy.vercel.app/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	rows := []Link{}
	doc.Find(".navbar").Children().Each(func(i int, sel *goquery.Selection) {
		row := Link{}
		row.title = sel.Find(".navbar-brand").Text()
		row.url, _ = sel.Find(".navbar-brand").Attr("href")
		rows = append(rows,row)
	})
	btw, err := json.MarshalIndent(rows, ""," ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(btw))
	return nil
}

type Article struct {
	URL string
	Title string
}


func main()  {
	//Links()
	Sedikitlagi()
	//Articles("https://csr-assessment-miqdadyyy.vercel.app")
	//Finishing("https://csr-assessment-miqdadyyy.vercel.app")/

	//SSR
	//ssr.FinishingSsr("https://ssr-assessment-miqdadyyy.vercel.app/")
}

func Finishing(url string)  {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var articles []Article
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {

		var article Article
		linkTag := item
		article.URL, _ = item.Attr("href")
		article.Title = linkTag.Text()

		//var links = []Link{
		//	{
		//		url: article.URL,
		//		title: article.Title,
		//	},
		//}

		articles = append(articles, article)
	})
	bts, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bts))
}
// sedikit lagi +2
func Articles(url string)  {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	articles := make([]Article,0)
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		article := new(Article)
		linkTag := item
		article.URL, _ = item.Attr("href")
		article.Title = linkTag.Text()
		articles = append(articles, *article)
	})
	bts, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(bts))
}


func linkScrape() {
	doc, err := goquery.NewDocument("http://jonathanmh.com")
	if err != nil {
		log.Fatal(err)
	}

	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		fmt.Printf("Link #%d: '%s' - '%s'\n", index, linkText, link)
	})
}

//type sedikit lagi +1
func Links()  {
	doc, err := goquery.NewDocument("https://csr-assessment-miqdadyyy.vercel.app")
	if err != nil {
		log.Fatal(err)
	}
	var links []Link
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		var linka Link
		linkTag := item
		linka.url, _ = item.Attr("href")
		linka.title = linkTag.Text()
		links = append(links, linka)

		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		fmt.Printf("Link #%d: '%s' - '%s'\n", index, linkText, link)
	})
	//bts, err := json.MarshalIndent(links, "", "  ")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//json.Marshal(links)
	fmt.Println(links)
}

func Sedikitlagi()  {
	//doc, err := goquery.NewDocument("http://jonathanmh.com")
	doc, err := goquery.NewDocument("https://csr-assessment-miqdadyyy.vercel.app")
	if err != nil {
	log.Fatal(err)
	}
	var links []ssr.Link
	// use CSS selector found with the browser inspector
	// for each, use index and item
	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		var link ssr.Link
		//linkTag := item
		link.Url, _ = item.Attr("href")
		link.Title = item.Text()
		//fmt.Printf("Link #%d: '%s' - '%s'\n", index, link.url,link.title, link)
		links = append(links, link)
	})
	bts, err := json.MarshalIndent(links, "", "  ")
	if err != nil {
	log.Fatal(err)
	}

	log.Println(string(bts))
}


func ExampleScrape(url string) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}