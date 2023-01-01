package models

import "log"

func GetAllLinks() ([]TinyUrl, error) {
    var tinyUrls []TinyUrl

    tx := db.Find(&tinyUrls)
    if tx.Error != nil {
        log.Fatal(tx.Error)
        return []TinyUrl{}, tx.Error
    }

    return tinyUrls, nil
}

func GetLink(id uint64) (TinyUrl, error){
    var singleUrl TinyUrl

    tx := db.Where("id = ?", id).First(&singleUrl)
    if tx.Error != nil {
        log.Fatal(tx.Error)
        return TinyUrl{}, tx.Error
    }

    return singleUrl, nil
}

func CreateLink(link TinyUrl) error {
    tx := db.Create(&link)

    return tx.Error
}

//func DeleteLink(id uint64) error {
//    tx := db.Unscoped().Delete(&TinyUrl{}, id)

//    return tx.Error
//}


func FindByUrl(url string)(TinyUrl, error)  {
    var redirectUrl TinyUrl

    tx := db.Where("short_url = ?", url).First(&redirectUrl)

    return redirectUrl, tx.Error
}
