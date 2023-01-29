package models

import "log"

// GetAllLinks() returns all the data stored in the database
// The returned data contains the url to be redirected to
// and the generated url
func GetAllLinks() ([]TinyUrl, error) {
	var tinyUrls []TinyUrl

	tx := db.Find(&tinyUrls)
	if tx.Error != nil {
		log.Fatal(tx.Error)
		return []TinyUrl{}, tx.Error
	}

	return tinyUrls, nil
}

// Returns a single link{the redirectUrl, the generatedurl}
// from the id
func GetLink(id uint64) (TinyUrl, error) {
	var singleUrl TinyUrl

	tx := db.Where("id = ?", id).First(&singleUrl)
	if tx.Error != nil {
		log.Fatal(tx.Error)
		return TinyUrl{}, tx.Error
	}

	return singleUrl, nil
}

// Adds a new generatedUrl to the database
// returns tx.Error
func CreateLink(link TinyUrl) error {
	tx := db.Create(&link)

	return tx.Error
}

// Updates counter when a redirectUrl is used
// returns tx.Error
func UpdateClick(link TinyUrl) error {
	tx := db.Save(&link)

	return tx.Error
}

// Returns the redirectUrl and tx.Error
// searches for the redirectUrl using the generatedurl
func FindByUrl(url string) (TinyUrl, error) {
	var redirectUrl TinyUrl

	tx := db.Where("short_url = ?", url).First(&redirectUrl)

	return redirectUrl, tx.Error
}
