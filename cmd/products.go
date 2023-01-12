package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func getProductsByCategory(category string) {
// 	products := []string{}
// 	c := colly.NewCollector(colly.AllowedDomains("shop-eu.palaceskateboards.com","www.shop-eu.palaceskateboards.com"))
// 	c.OnHTML("div[id=content]", func(h *colly.HTMLElement) {
// 		h2 := h.DOM.Find("div[id=product-loop]").First().SetAttr("scrollTop","100000000")
// 		str := h2.Text()
// 		words := []string{}
// 		splitted := strings.Split(str, " \n")
// 		for _, word := range splitted {
// 			word = strings.TrimSpace(word)
// 			if word != "" {
// 			words = append(words, word)
// 			}
// 		}
// 		for _,product := range words {
// 			if strings.Contains(product,"€") || product == "SOLD OUT"{
// 				continue
// 			}else{
// 				products = append(products, product)
// 			}
// 		}
// 		for _,product := range products {
// 			if strings.Contains(product, "€"){
// 				fmt.Println("pret")
// 			}else{
// 				fmt.Println(product)
// 			}
// 		}
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("visiting", r.URL.String())
// 	})

// 	c.Visit("https://shop-eu.palaceskateboards.com/collections/"+category)

// 	addProductPalace(products, category)
// }


func getPalaceProductsList(w http.ResponseWriter, r *http.Request) {
	products := getPalaceProducts()

	out,err := json.Marshal(products)
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func getSupremeProductsList(w http.ResponseWriter, r *http.Request) {
	products := getSupremeProducts()

	out,err := json.Marshal(products)
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

