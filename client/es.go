package client

import (
	"context"
	"log"

	es "github.com/olivere/elastic/v7"
	escfg "github.com/olivere/elastic/v7/config"
)

// EsConfig ES的配置文件
type EsConfig struct {
	URL      string
	Username string
	Password string
	Sniff    bool
}

type EsClient struct {
	esCli  es.Client
	callee string
}

// InitEsClient 通过配置文件初始化一个ES的客户端
func InitEsClient(config EsConfig) (*es.Client, error) {
	var esConfig = &escfg.Config{
		URL:      config.URL,
		Username: config.Username,
		Password: config.Password,
		Sniff:    &config.Sniff,
	}

	esClient, err := es.NewClientFromConfig(esConfig)
	if err != nil {
		log.Fatalf("init es client err: [%+v]", err)
		return nil, err

	}
	return esClient, nil
}

// IsDocExists 某条记录是否存在
func (cli *EsClient) IsDocExists(ctx context.Context, id, indexName string) (bool, error) {
	ok, err := cli.esCli.Exists().Index(indexName).Id(id).Do(ctx)
	if err != nil {
		log.Fatalf("IsDocExists err: [%+v]", err)
		return false, err
	}

	return ok, nil
}

// CreateIndexData 构建新索引
func (cli *EsClient) CreateIndexData(ctx context.Context, indexName, indexSetting string) error {
	exist, err := cli.esCli.IndexExists(indexName).Do(ctx)
	if err != nil {
		log.Fatalf("Query Index Exist Error: [%+v]", err)
		return err
	}

	if !exist {
		_, err = cli.esCli.CreateIndex(indexName).BodyString(indexSetting).Do(ctx)
		if err != nil {
			log.Fatalf("create index error, %+v", err)
			return err
		}

		return nil
	}

	return nil
}

// DoRecall 召回
func (cli *EsClient) DoRecall(ctx context.Context, bq *es.BoolQuery, page, limit int,
	indexName string) (*es.SearchResult, error) {
	results, err := cli.esCli.Search().Index(indexName).Query(bq).From(page).Size(limit).Do(ctx)

	if err != nil {
		log.Fatalf("Search Index err: [+%v]", err)
		return nil, err
	}

	return results, nil
}

// AddMatchQuery 组合ES的查询
func AddMatchQuery(bq *es.BoolQuery, name string, values []string) {
	for _, v := range values {
		bq.Should(es.NewMatchQuery(name, v))
	}
}
