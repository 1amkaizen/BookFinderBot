package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Produk represents a product with multiple affiliate links
type Produk struct {
	Nama  string            `json:"name"`
	Links map[string]string `json:"links"`
}

// loadProductsFromTxt reads and parses the text file containing product data
func loadProductsFromTxt(filename string) ([]Produk, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var products []Produk
	var currentProduct Produk
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if currentProduct.Nama != "" {
				products = append(products, currentProduct)
				currentProduct = Produk{}
			}
			continue
		}
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			if currentProduct.Links == nil {
				currentProduct.Links = make(map[string]string)
			}
			currentProduct.Links[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		} else {
			currentProduct.Nama = line
		}
	}
	if currentProduct.Nama != "" {
		products = append(products, currentProduct)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

// loadReviewLinksFromTxt reads and parses the text file containing product review links
func loadReviewLinksFromTxt(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reviewLinks := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		productName := strings.TrimSpace(parts[0])
		link := strings.TrimSpace(parts[1])
		reviewLinks[productName] = link
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return reviewLinks, nil
}

// saveProductsToJson saves the products list to a JSON file
func saveProductsToJson(products []Produk, filename string) error {
	data, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// saveReviewLinksToJson saves the review links to a JSON file
func saveReviewLinksToJson(reviewLinks map[string]string, filename string) error {
	data, err := json.MarshalIndent(reviewLinks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// extractKeywords generates a list of keywords from the product name
func extractKeywords(name string) []string {
	words := strings.Fields(name)
	keywords := make([]string, len(words))
	for i, word := range words {
		keywords[i] = strings.ToLower(word)
	}
	return keywords
}

// findProducts searches for products based on keywords
func findProducts(products []Produk, message string) []*Produk {
	message = strings.ToLower(message)
	var matchingProducts []*Produk
	for i := range products {
		keywords := extractKeywords(products[i].Nama)
		for _, keyword := range keywords {
			if strings.Contains(message, keyword) {
				matchingProducts = append(matchingProducts, &products[i])
				break // Break the inner loop once a match is found
			}
		}
	}
	return matchingProducts
}

// findReviewLinkByName searches for review link based on product name
func findReviewLinkByName(reviewLinks map[string]string, productName string) (string, bool) {
	link, found := reviewLinks[productName]
	return link, found
}

func main() {
	// Load products from text file
	products, err := loadProductsFromTxt("products.txt")
	if err != nil {
		log.Fatalf("Gagal memuat produk: %v", err)
	}

	// Load review links from text file
	reviewLinks, err := loadReviewLinksFromTxt("review_links.txt")
	if err != nil {
		log.Fatalf("Gagal memuat link review: %v", err)
	}

	// Save products to JSON file
	err = saveProductsToJson(products, "products.json")
	if err != nil {
		log.Fatalf("Gagal menyimpan produk ke JSON: %v", err)
	}

	// Save review links to JSON file
	err = saveReviewLinksToJson(reviewLinks, "review_links.json")
	if err != nil {
		log.Fatalf("Gagal menyimpan link review ke JSON: %v", err)
	}

	// Load products from JSON file (optional, for consistency)
	data, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatalf("Gagal membaca produk dari JSON: %v", err)
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		log.Fatalf("Gagal mengurai produk dari JSON: %v", err)
	}

	// Load review links from JSON file (optional, for consistency)
	data, err = ioutil.ReadFile("review_links.json")
	if err != nil {
		log.Fatalf("Gagal membaca link review dari JSON: %v", err)
	}
	err = json.Unmarshal(data, &reviewLinks)
	if err != nil {
		log.Fatalf("Gagal mengurai link review dari JSON: %v", err)
	}

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Panic("TELEGRAM_BOT_TOKEN tidak diatur")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Text {
		case "/start":
			msg.Text = "📚 Selamat datang di BokFinderBot! Saya adalah bot pencari Ebook & Buku. Cari Ebook apa yang Anda butuhkan? Ketikkan judul atau topik yang Anda inginkan, dan saya akan mencarikannya untuk Anda."
		case "/help":
			msg.Text = `ℹ️ Gunakan bot ini untuk mencari Ebook & Buku. Anda cukup ketik judul atau topik yang ingin Anda cari, dan saya akan mencarikannya untuk Anda.

🔍 Contoh penggunaan:
Ketikkan "Belajar Python" untuk mencari Ebook atau Buku tentang pemrograman Python.
Ketikkan "Hacking" untuk mencari Ebook atau Buku tentang hacking.

📖 Anda juga bisa menggunakan perintah:
/ulasan [nama lengkap produk] untuk mendapatkan link ulasan produk tersebut.

⚠️ Perhatian: Judul harus sesuai, perhatikan huruf besar dan kecilnya agar mendapatkan link ulasan.

📘 Contoh penggunaan:
/ulasan Ilmu Hacking 
untuk mendapatkan link ulasan buku Ilmu Hacking.

📝 Catatan:
Kamu juga bisa memberikan ulasan di sini:
https://aigoretech.rf.gd/kirim-ulasan`
		case "/ulasan":
			msg.Text = "⚠️ Mohon berikan judul lengkap buku untuk mendapatkan link ulasannya.\nContoh penggunaan: /ulasan Judul Buku"
		default:
			if strings.HasPrefix(update.Message.Text, "/ulasan ") {
				productName := strings.TrimPrefix(update.Message.Text, "/ulasan ")
				if link, found := findReviewLinkByName(reviewLinks, productName); found {
					msg.Text = "📘 Link ulasan untuk " + productName + ":\n" + link
				} else {
					msg.Text = "⚠️ Link ulasan untuk " + productName + " tidak ditemukan.\nKamu bisa memberikan ulasan di sini: https://aigoretech.rf.gd/kirim-ulasan"
				}
			} else {
				matchingProducts := findProducts(products, update.Message.Text)
				if len(matchingProducts) > 0 {
					for _, produk := range matchingProducts {
						// Create message for each product
						productMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
						productMsg.Text += "\n\n📖 Judul: " + produk.Nama

						// Create inline keyboard with buttons for all links
						keyboard := [][]tgbotapi.InlineKeyboardButton{}
						for platform, link := range produk.Links {
							btn := tgbotapi.NewInlineKeyboardButtonURL(platform, link)
							row := []tgbotapi.InlineKeyboardButton{btn}
							keyboard = append(keyboard, row)

							// Add link text below the button
							productMsg.Text += "\n🔗 " + platform + ": " + link
						}
						inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(keyboard...)
						productMsg.ReplyMarkup = inlineKeyboard

						// Send message for each product
						if _, err := bot.Send(productMsg); err != nil {
							log.Panic(err)
						}
					}
				} else {
					msg.Text = "⚠️ Produk tidak ditemukan."
				}
			}
		}

		// Send general message or error message if needed
		if msg.Text != "" {
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}
