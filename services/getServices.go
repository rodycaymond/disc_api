package services

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Content struct {
	DataId     string `json:"dataId"`
	Speed      string `json:"speed"`
	Glide      string `json:"glide"`
	Turn       string `json:"turn"`
	Fade       string `json:"fade"`
	Name       string `json:"name"`
	FlightType string `json:"flightType"`
	Brand      string `json:"brand"`
	FlightPath string `json:"flightPath"`
}

type ByStability struct {
	Stability string    `json:"stability"`
	Discs     []Content `json:"discs"`
}

type EveryDisc struct {
	VeryOverstable  []Content `json:"veryOverstable"`
	Overstable      []Content `json:"overstable"`
	Stable          []Content `json:"stable"`
	Understable     []Content `json:"understable"`
	VeryUnderstable []Content `json:"veryUnderstable"`
}

var url = "https://www.marshallstreetdiscgolf.com/flightguide"

func buildDisc(elem *goquery.Selection) []Content {
	var discs = []Content{}

	elem.Find(".disc-item").Each(func(int int, e *goquery.Selection) {
		dataId, err := e.Attr("data-id")
		if err == false {
			dataId = ""
		}
		speed, err := e.Attr("data-speed")
		if err == false {
			speed = ""
		}
		glide, err := e.Attr("data-glide")
		if err == false {
			glide = ""
		}
		turn, err := e.Attr("data-turn")
		if err == false {
			turn = ""
		}
		fade, err := e.Attr("data-fade")
		if err == false {
			fade = ""
		}
		name, err := e.Attr("data-title")
		if err == false {
			name = ""
		}
		flightT, err := e.Attr("data-category")
		if err == false {
			flightT = ""
		}
		brand, err := e.Attr("data-brand")
		if err == false {
			brand = ""
		}
		flightP, err := e.Attr("data-pic")
		if err == false {
			flightP = ""
		}
		disc := Content{
			DataId:     dataId,
			Speed:      speed,
			Glide:      glide,
			Turn:       turn,
			Fade:       fade,
			Name:       name,
			FlightType: flightT,
			Brand:      brand,
			FlightPath: flightP,
		}
		discs = append(discs, disc)
	})
	return discs
}

func buildPutter(elem *goquery.Selection) []Content {
	var discs = []Content{}

	elem.Find(".pc-entry").Each(func(int int, e *goquery.Selection) {
		dataId, err := e.Attr("data-id")
		if err == false {
			dataId = strconv.Itoa(int)
		}
		speed, err := e.Attr("data-speed")
		if err == false {
			speed = ""
		}
		glide, err := e.Attr("data-glide")
		if err == false {
			glide = ""
		}
		turn, err := e.Attr("data-turn")
		if err == false {
			turn = ""
		}
		fade, err := e.Attr("data-fade")
		if err == false {
			fade = ""
		}
		name, err := e.Attr("data-putter")
		if err == false {
			name = ""
		}
		flightT, err := e.Attr("data-category")
		if err == false {
			flightT = "Putter"
		}
		brand, err := e.Attr("data-brand")
		if err == false {
			brand = ""
		}
		flightP, err := e.Attr("data-image")
		if err == false {
			flightP = ""
		}
		disc := Content{
			DataId:     dataId,
			Speed:      speed,
			Glide:      glide,
			Turn:       turn,
			Fade:       fade,
			Name:       name,
			FlightType: flightT,
			Brand:      brand,
			FlightPath: flightP,
		}
		discs = append(discs, disc)
	})
	return discs
}

func AllDiscs() EveryDisc {
	var allDiscs = EveryDisc{
		VeryOverstable:  VeryOverStable(),
		Overstable:      OverStable(),
		Stable:          Stable(),
		Understable:     Understable(),
		VeryUnderstable: VeryUnderstable(),
	}

	return allDiscs
}

func VeryUnderstable() []Content {
	var allDiscs = []Content{}

	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	document.Find(".very-understable").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildDisc(element)
		allDiscs = append(allDiscs, sliceToAdd...)
	})

	return allDiscs
}
func Understable() []Content {
	var allDiscs = []Content{}
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("div.understable").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildDisc(element)
		allDiscs = append(allDiscs, sliceToAdd...)
	})

	return allDiscs
}

func Stable() []Content {
	var allDiscs = []Content{}
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("div.stable").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildDisc(element)
		allDiscs = append(allDiscs, sliceToAdd...)
	})

	return allDiscs
}

func OverStable() []Content {
	var allDiscs = []Content{}
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("div.overstable").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildDisc(element)
		allDiscs = append(allDiscs, sliceToAdd...)
	})

	return allDiscs
}

func VeryOverStable() []Content {
	var allDiscs = []Content{}
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find("div.very-overstable").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildDisc(element)
		allDiscs = append(allDiscs, sliceToAdd...)
	})

	return allDiscs
}

func Putters() []Content {
	var allDiscs = []Content{}
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".putter-parent").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildPutter(element)
		allDiscs = append(allDiscs, sliceToAdd...)
	})

	return allDiscs
}

func filterByType(slices []Content, name string, b bool) []Content {
	var newSlice []Content
	if b == true {
		for i := 0; i < len(slices); i++ {
			if strings.ToLower(slices[i].Brand) == strings.ToLower(name) {
				newSlice = append(newSlice, slices[i])
			}
		}
	} else {
		for i := 0; i < len(slices); i++ {
			if strings.ToLower(slices[i].Name) == strings.ToLower(name) {
				newSlice = append(newSlice, slices[i])
			}
		}
	}

	return newSlice
}

func ByBrand(brandName string) []Content {
	var allDiscs = []Content{}
	formattedBrand := strings.ReplaceAll(brandName, "-", " ")
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".flex-grid-item .flex-column").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildDisc(element)
		filteredSlice := filterByType(sliceToAdd, formattedBrand, true)
		allDiscs = append(allDiscs, filteredSlice...)
	})
	return allDiscs
}

func ByPutterBrand(brandName string) []Content {
	var allDiscs = []Content{}
	formattedBrand := strings.ReplaceAll(brandName, "-", " ")
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".putter-parent").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildPutter(element)
		filteredSlice := filterByType(sliceToAdd, formattedBrand, true)
		allDiscs = append(allDiscs, filteredSlice...)
	})
	return allDiscs
}

func ByName(name string) []Content {
	var singleDisc = []Content{}
	formattedName := strings.ReplaceAll(name, "-", " ")
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".flex-grid-item .flex-column").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildDisc(element)
		filteredSlice := filterByType(sliceToAdd, formattedName, false)
		singleDisc = append(singleDisc, filteredSlice...)
	})
	return singleDisc
}

func ByPutterName(name string) []Content {
	var singleDisc = []Content{}
	formattedName := strings.ReplaceAll(name, "-", " ")
	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	document.Find(".putter-parent").Each(func(index int, element *goquery.Selection) {
		sliceToAdd := buildPutter(element)
		filteredSlice := filterByType(sliceToAdd, formattedName, false)
		singleDisc = append(singleDisc, filteredSlice...)
	})
	return singleDisc
}
