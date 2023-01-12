package main

import (
	"fmt"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func palace(Task ReadyTask) {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)

	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)

	_, err = page.Goto("https://shop-eu.palaceskateboards.com")
	assertErrorToNilf("could not goto: %w", err)
	time.Sleep(2 * time.Second)
	page.Click("text=" + Task.Task.Category)

	time.Sleep(2 * time.Second)
	page.Click("text=" + Task.Task.ProductName)

	time.Sleep(2 * time.Second)
	
	if Task.Task.Size != "One Size"{
		labelSize := []string{Task.Task.Size}
		selectSize,_ := page.QuerySelector(`select[id=product-select]`)
		selectSize.SelectOption(playwright.SelectOptionValues{Labels: &labelSize})
	}
	
	time.Sleep(2 * time.Second)
	page.Click("text="+"ADD TO CART")
	
	time.Sleep(2 * time.Second)
	page.Click("text="+"CART")

	
	time.Sleep(2 * time.Second)
	page.Click(`div[class="checkbox-control"]`)
	time.Sleep(2 * time.Second)
	page.Click("text="+"CHECKOUT")
	time.Sleep(2 * time.Second)
	page.Fill(`input[name="checkout[email]"]`,Task.Profile.Email)
	page.Fill(`input[name="checkout[shipping_address][country]"]`,Task.Profile.Country)
	page.Fill(`input[name="checkout[shipping_address][first_name]"]`,Task.Profile.Fname)
	page.Fill(`input[name="checkout[shipping_address][last_name]"]`,Task.Profile.Lname)
	page.Fill(`input[name="checkout[shipping_address][address1]"]`,Task.Profile.Address)
	page.Fill(`input[name="checkout[shipping_address][address2]"]`,Task.Profile.Address2)
	page.Fill(`input[name="checkout[shipping_address][zip]"]`, Task.Profile.Postcode)
	page.Fill(`input[name="checkout[shipping_address][city]"]`, Task.Profile.City)
	page.Fill(`input[name="checkout[shipping_address][province]"]`, Task.Profile.County)
	page.Fill(`input[name="checkout[shipping_address][phone]"]`,Task.Profile.Phone)
	time.Sleep(2 * time.Second)
	page.Click("text="+"Continue to shipping")
	time.Sleep(2 * time.Second)
	page.Click("text="+"Continue to payment")
	time.Sleep(2 * time.Second)
	page.Click("text=" + "Card number")
	page.Keyboard().InsertText(Task.Card.CardNumber)
	time.Sleep(2 * time.Second)
	page.Click("text=" + "Name on card")
 	page.Keyboard().InsertText(Task.Card.NameCard)
	time.Sleep(2 * time.Second)
	page.Click("text=" + "Expiration date")
 	page.Keyboard().InsertText(Task.Card.ExpirationDate)
	time.Sleep(2 * time.Second)
	page.Click("text=" + "Security code")
 	page.Keyboard().InsertText(Task.Card.Cvv)
	time.Sleep(5 * time.Second)
	context.Close()
	browser.Close()
}

func supreme(Task ReadyTask) {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	assertErrorToNilf("could not launch Chromium: %w", err)

	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)

	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)

	_, err = page.Goto("https://www.supremenewyork.com/shop/all")
	assertErrorToNilf("could not goto: %w", err)

	page.Click("text=" + Task.Task.Category)
	time.Sleep(2 * time.Second)
 	page.Click("text=" + Task.Task.ProductName + " " + Task.Task.Color)
 	time.Sleep(2 * time.Second)

 	time.Sleep(2 * time.Second)
	if Task.Task.Size != "One Size"{
		labelSize := []string{Task.Task.Size}
		selectSize,_ := page.QuerySelector(`select[data-cy="size-selector"]`)
		selectSize.SelectOption(playwright.SelectOptionValues{Labels: &labelSize})
	}
 	
	
	page.Click(`input[name="add"]`)
 	time.Sleep(2 * time.Second)
	//url := page.URL()
	_,err = page.Goto("https://eu.supreme.com/checkout")
	
	assertErrorToNilf("could not goto: %w", err)

	labelCountry := []string{Task.Profile.Country}
	selectCountry,_ := page.QuerySelector(`select[name="countryCode"]`)
	selectCountry.SelectOption(playwright.SelectOptionValues{Labels: &labelCountry})

	
	labelCounty := []string{Task.Profile.County}
	fmt.Println(Task.Profile.County)
	selectCounty,_ := page.QuerySelector(`select[name="zone"]`)
	selectCounty.SelectOption(playwright.SelectOptionValues{Labels: &labelCounty})
	
	
	page.Fill(`input[name="firstName"]`, Task.Profile.Fname)
 	page.Fill(`input[name="lastName"]`, Task.Profile.Lname)
 	page.Fill(`input[name="email"]`, Task.Profile.Email)
 	page.Fill(`input[name="address1"]`, Task.Profile.Address)
 	page.Fill(`input[name="address2"]`, Task.Profile.Address2)
 	page.Fill(`input[name="postalCode"]`, Task.Profile.Postcode)
 	page.Fill(`input[name="city"]`, Task.Profile.City)
 	page.Fill(`input[name="phone"]`, Task.Profile.Phone)
	
 	page.Click("text=" + "card number")
	page.Keyboard().InsertText(Task.Card.CardNumber)
	page.Click("text=" + "name on card")
 	page.Keyboard().InsertText(Task.Card.NameCard)
	page.Click("text=" + "expiration date")
 	page.Keyboard().InsertText(Task.Card.ExpirationDate)
	page.Click("text=" + "security code")
 	page.Keyboard().InsertText(Task.Card.CardNumber)
	page.Click(`input[id="accept_tos"]`) 
	time.Sleep(10 * time.Second)
	context.Close()
	browser.Close()
}
