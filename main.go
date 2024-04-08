package main

import (
	"ElasticSearch/elasticsearch"
	"ElasticSearch/mariadb"
	"fmt"
	"log"
	"strconv"
)

func main() {

	db, err := mariadb.Init()
	if err != nil {
		return
	}

	read := &mariadb.Hotel{}
	db.First(&read, "ID = ?", 38609)

	fmt.Println(read)

	var hotels []mariadb.Hotel

	result := db.Find(&hotels)

	if result.Error != nil {
		log.Fatalf("Failed to query tb_hotel: %v", result.Error)
	}

	list := make([]elasticsearch.HotelElasticSearch, 0)

	for _, hotel := range hotels {

		float1, err := strconv.ParseFloat(hotel.Latitude, 64)
		if err != nil {
			return
		}
		float2, err := strconv.ParseFloat(hotel.Latitude, 64)
		if err != nil {
			return
		}

		elasticSearch := elasticsearch.HotelElasticSearch{
			ID:        hotel.ID,
			Name:      hotel.Name,
			Address:   hotel.Address,
			Brand:     hotel.Brand,
			City:      hotel.City,
			Rating:    hotel.StarName,
			District:  hotel.Business,
			Latitude:  float1,
			Longitude: float2,
			ImageURL:  hotel.Pic,
			Tags:      nil,
		}

		list = append(list, elasticSearch)
	}

	client, err := elasticsearch.Init()
	if err != nil {
		return
	}

	for _, search := range list {
		elasticsearch.UpdateIndexDocument(client, search)
	}

	elasticsearch.SearchDocument(client)

}
