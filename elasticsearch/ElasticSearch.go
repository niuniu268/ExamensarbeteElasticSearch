package elasticsearch

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"strconv"
)

type HotelElasticSearch struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Address   string   `json:"address"`
	Brand     string   `json:"brand"`
	City      string   `json:"city"`
	Rating    string   `json:"rating"`
	District  string   `json:"district"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	ImageURL  string   `json:"imageURL"`
	Tags      []string `json:"tags"`
}

func Init() (*elasticsearch.TypedClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.1.72:9200",
		},
	}
	client, err := elasticsearch.NewTypedClient(cfg)

	return client, err
}

func CreateIndex(client *elasticsearch.TypedClient) {
	resp, err := client.Indices.
		Create("hotel-index").
		Do(context.Background())
	if err != nil {
		fmt.Printf("create index failed, err:%v\n", err)
		return
	}
	fmt.Printf("index:%#v\n", resp.Index)
}

func SearchDocument(client *elasticsearch.TypedClient) {
	// 搜索文档
	resp, err := client.Search().
		Index("hotel-index").
		Query(&types.Query{
			MatchAll: &types.MatchAllQuery{},
		}).
		Do(context.Background())
	if err != nil {
		fmt.Printf("search document failed, err:%v\n", err)
		return
	}
	fmt.Printf("total: %d\n", resp.Hits.Total.Value)
	// 遍历所有结果
	for _, hit := range resp.Hits.Hits {
		fmt.Printf("%s\n", hit.Source_)
	}

}

func UpdateIndexDocument(client *elasticsearch.TypedClient, hotel HotelElasticSearch) {

	// 添加文档
	resp, err := client.Index("hotel-index").
		Id(strconv.FormatInt(hotel.ID, 10)).
		Document(hotel).
		Do(context.Background())
	if err != nil {
		fmt.Printf("indexing document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%#v\n", resp.Result)
}
