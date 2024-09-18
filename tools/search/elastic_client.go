package search
//
//import (
//	"bytes"
//	"context"
//	"encoding/json"
//	stdjson "encoding/json"
//	"errors"
//	"fmt"
//	"github.com/pengcainiao/zero/core/logx"
//	"io"
//	"net/http"
//	"time"
//
//	"github.com/pengcainiao/zero/core/env"
//	zslog "github.com/rs/zerolog/log"
//)
//
//var (
//	es *ElasticClient
//	//ElasticEmptySearchResponse 当查询无数据时，返回
//	ElasticEmptySearchResponse = map[string]interface{}{
//		"total": 0,
//		"data":  []byte("[]"),
//	}
//)
//
////BulkOperationMethod 批量操作方法
//type BulkOperationMethod string
//
//const (
//	// BulkCreate 批量创建
//	BulkCreate BulkOperationMethod = "bulk_create"
//	// BulkUpdate 批量更新
//	BulkUpdate BulkOperationMethod = "bulk_update"
//	// BulkDelete 批量删除
//	BulkDelete BulkOperationMethod = "bulk_delete"
//)
//
//// ElasticClient 定义ES客户端
//type ElasticClient struct {
//	client *elastic.Client
//}
//
//// QueryBase 定义查询条件
//type QueryBase struct {
//	ElasticQuerys []elastic.Query //查询条件
//	//FetchSource   *elastic.FetchSourceContext //查询字段
//}
//
////MultiGetRequest 批量查询指定文档ID
//type MultiGetRequest struct {
//	IndexName      string
//	FetchSource    *elastic.FetchSourceContext
//	DocumentIDList []string
//}
//
////MultiAddRequest 批量新增
//type MultiAddRequest struct {
//	IndexName string                 //索引名称
//	Documents map[string]interface{} //文档列表
//}
//
////AggregationRequest 聚合查询
//type AggregationRequest struct {
//	ElasticNormalRequest
//	QueryBase
//	Aggregations map[string]elastic.Aggregation
//}
//
////ElasticNormalRequest 文档ID查询
//type ElasticNormalRequest struct {
//	IndexName   string                      //索引名称
//	DocumentID  string                      //文档ID
//	FetchSource *elastic.FetchSourceContext //查询字段
//	Refresh     bool                        //刷新
//}
//
////SearchRequest 查询请求
//type SearchRequest struct {
//	ElasticNormalRequest
//	QueryBase
//	SortFields     map[string]bool        //排序字段
//	SortByFields   []elastic.Sorter       //
//	QueryParameter map[string]interface{} //查询参数
//	PageOffset     int                    // 分页偏移量
//	PageSize       int                    //每页数据大小
//}
//
////CreateRequest 创建新文档
//type CreateRequest struct {
//	ElasticNormalRequest
//	Payload interface{} //内容，支持JSON序列化
//}
//
////UpdateRequest 更新指定文档ID
//type UpdateRequest struct {
//	ElasticNormalRequest
//	Script   string                 //更新字段脚本
//	FieldMap map[string]interface{} //更新字段与值的映射
//}
//
////UpdateByQueryRequest 更新指定条件查出的结果集
//type UpdateByQueryRequest struct {
//	UpdateRequest
//	QueryBase
//	Size int // 更新条数
//}
//
////DeleteByQueryRequest 删除指定条件查出的结果集
//type DeleteByQueryRequest struct {
//	ElasticNormalRequest
//	ElasticQuerys []elastic.Query //查询条件
//}
//
////BulkRequest 批量操作
//type BulkRequest struct {
//	Method        BulkOperationMethod //方法名称
//	Payload       interface{}         //有效数据
//	UpdateRequest                     //更新数据体
//	//ElasticNormalRequest                        //
//	//Script               string                 //脚本
//	//FieldMap             map[string]interface{} //键值映射
//}
//
////Query 添加新的查询条件
//func (s *QueryBase) Query(query elastic.Query) {
//	if s.ElasticQuerys == nil {
//		s.ElasticQuerys = make([]elastic.Query, 0)
//	}
//	s.ElasticQuerys = append(s.ElasticQuerys, query)
//}
//
//func (e *ElasticClient) OriginClient() *elastic.Client {
//	return e.client
//}
//
//// ElasticSearch ES客户端
//func ElasticSearch() *ElasticClient {
//	if es == nil {
//		es = newClient()
//	}
//	return es
//}
//
//func newClient() *ElasticClient {
//	client, err := elastic.NewClient(
//		elastic.SetURL(env.ElasticAddr),
//		elastic.SetBasicAuth(env.ElasticUser, env.ElasticPassword),
//		elastic.SetSniff(false),
//		elastic.SetHealthcheckInterval(10*time.Second),
//		elastic.SetRetrier(elastic.NewBackoffRetrier(elastic.NewConstantBackoff(time.Second*10))),
//		elastic.SetGzip(true),
//		elastic.SetHeaders(http.Header{
//			"X-Caller-Id": []string{"..."},
//		}),
//	)
//	if err != nil {
//		panic(err)
//	}
//	return &ElasticClient{client: client}
//}
//
//// CreateIndex 创建索引
//func (e *ElasticClient) CreateIndex(indexName string, mapping string) error {
//	// 创建index前，先查看es引擎中是否存在自己想要创建的索引index
//	exists, err := e.client.IndexExists(indexName).Do(context.Background())
//	if err != nil {
//		return e.handleElasticError(err, "-1")
//	}
//	if !exists {
//		// 如果不存在，就创建
//		createIndex, err := e.client.CreateIndex(indexName).Body(mapping).Do(context.Background())
//		if err != nil {
//			panic(err)
//		}
//		if !createIndex.Acknowledged {
//			// Not acknowledged ,创建失败
//			panic(err)
//		}
//	}
//	return nil
//}
//
//// CreateSettings 创建settings
//func (e *ElasticClient) CreateSettings(indexName string, mapping string) error {
//	// 创建index前，先查看es引擎中是否存在自己想要创建的索引index
//	exists, err := e.client.IndexExists(indexName).Do(context.Background())
//	if err != nil {
//		return e.handleElasticError(err, "-1")
//	}
//	if !exists {
//		// 如果不存在，就创建
//		createIndex, err := e.client.IndexPutSettings(indexName).BodyString(mapping).Do(context.Background())
//		if err != nil {
//			panic(err)
//		}
//		if !createIndex.Acknowledged {
//			// Not acknowledged ,创建失败
//			panic(err)
//		}
//	}
//	return nil
//}
//
//// AppendMapping 追加mapping
//func (e *ElasticClient) AppendMapping(indexName string, mapping string) error {
//	_, err := e.client.PutMapping().Index(indexName).BodyString(mapping).Do(context.Background())
//	return err
//}
//
////Create 创建新的文档
//func (e *ElasticClient) Create(req CreateRequest) error {
//	b, err := e.interfaceToBytes(req.Payload)
//	if err != nil {
//		return err
//	}
//	_, err = e.client.Index().Index(req.IndexName).Id(req.DocumentID).BodyJson(string(b)).Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////Get 查询指定文档ID的数据
//func (e *ElasticClient) Get(req ElasticNormalRequest) ([]byte, error) {
//	resp, err := e.client.Get().Index(req.IndexName).Id(req.DocumentID).
//		FetchSourceContext(req.FetchSource).
//		Do(context.Background())
//	if err != nil {
//		err = e.handleElasticError(err, req.DocumentID)
//		return nil, err
//	}
//	if resp.Source != nil {
//		return resp.Source, nil
//	}
//	return nil, errors.New("response is nil")
//}
//
////Update 更新指定文档ID
//func (e *ElasticClient) Update(req UpdateRequest) error {
//	//script := elastic.NewScript("context.Background()._source.read_at += params.read_at").Param("read_at", time.Now().Unix())
//	script := elastic.NewScript(req.Script) //.Param("read_at", time.Now().Unix())
//	for key, value := range req.FieldMap {
//		script = script.Param(key, value)
//	}
//
//	_, err := e.client.Update().Index(req.IndexName).Id(req.DocumentID).
//		RetryOnConflict(3).
//		Script(script).
//		Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////UpdateByQuery 修改指定查询条件得到的结果集
//func (e *ElasticClient) UpdateByQuery(req UpdateByQueryRequest) error {
//	script := elastic.NewScript(req.Script) //.Param("read_at", time.Now().Unix())
//	for key, value := range req.FieldMap {
//		script = script.Param(key, value)
//	}
//	updateService := e.client.UpdateByQuery().ProceedOnVersionConflict().Index(req.IndexName)
//	for _, query := range req.ElasticQuerys {
//		updateService = updateService.Query(query)
//	}
//
//	if req.Size == 0 {
//		req.Size = 20
//	}
//	_, err := updateService.Size(req.Size).Script(script).Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////Delete 删除指定文档ID
//func (e *ElasticClient) Delete(req ElasticNormalRequest) error {
//	_, err := e.client.Delete().Index(req.IndexName).Id(req.DocumentID).Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////DeleteByQuery 删除符合指定查询条件的结果集
//func (e *ElasticClient) DeleteByQuery(req DeleteByQueryRequest) error {
//	deleteService := e.client.DeleteByQuery().Index(req.IndexName)
//	for _, query := range req.ElasticQuerys {
//		deleteService = deleteService.Query(query)
//	}
//
//	_, err := deleteService.Do(context.Background())
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
////BulkAdd 批量新增
//func (e *ElasticClient) BulkAdd(req MultiAddRequest) error {
//	service := e.client.Bulk()
//	for key, value := range req.Documents {
//		service = service.Add(elastic.NewBulkIndexRequest().Index(req.IndexName).Id(key).Doc(value))
//	}
//	_, err := service.Do(context.Background())
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//// BulkGet 批量查询文档ID
//func (e *ElasticClient) BulkGet(req MultiGetRequest) ([]byte, error) {
//	service := e.client.MultiGet()
//	for _, docID := range req.DocumentIDList {
//		service.Add(elastic.NewMultiGetItem().Index(req.IndexName).FetchSource(req.FetchSource).Id(docID))
//	}
//	resp, err := service.Do(context.Background())
//	if err != nil {
//		err = e.handleElasticError(err, "multiget")
//		return nil, err
//	}
//	var (
//		sourceLen  = len(resp.Docs)
//		buf        bytes.Buffer
//		sourceByte stdjson.RawMessage
//	)
//	buf.Write([]byte("["))
//	for i := 0; i < sourceLen; i++ {
//		var source = resp.Docs[i]
//		sourceByte = source.Source
//		if source.Error != nil {
//			zslog.Debug().Str("err", fmt.Sprintf("%v", source.Error)).Msg("通过document_id批量查询出错")
//		} else if sourceByte != nil {
//			if i > 0 {
//				buf.Write([]byte(","))
//			}
//			buf.Write(source.Source)
//		}
//	}
//	buf.Write([]byte("]"))
//	return buf.Bytes(), nil
//}
//
////Aggregation 聚合查询
//func (e *ElasticClient) Aggregation(req AggregationRequest) (*elastic.SearchResult, error) {
//	var searchService = e.client.Search().Index(req.IndexName)
//	if req.FetchSource != nil {
//		searchService = searchService.FetchSourceContext(req.FetchSource)
//	}
//	for _, query := range req.ElasticQuerys {
//		searchService = searchService.Query(query)
//	}
//	for keyName, aggregation := range req.Aggregations {
//		searchService = searchService.Aggregation(keyName, aggregation)
//	}
//	searchResult, err := searchService.Do(context.Background())
//
//	if err != nil && err != io.EOF {
//		return nil, err
//	}
//	return searchResult, nil
//}
//
//// Bulk 批量操作
//func (e *ElasticClient) Bulk(reqs []BulkRequest) error {
//	bulkRequest := e.client.Bulk()
//	for _, req := range reqs {
//		switch req.Method {
//		case BulkCreate:
//			r := elastic.NewBulkIndexRequest().Index(req.IndexName).Id(req.DocumentID).Doc(req.Payload)
//			bulkRequest = bulkRequest.Add(r)
//		case BulkUpdate:
//			script := elastic.NewScript(req.Script)
//			for key, value := range req.FieldMap {
//				script = script.Param(key, value)
//			}
//			r := elastic.NewBulkUpdateRequest().Index(req.IndexName).Id(req.DocumentID).
//				RetryOnConflict(3).
//				Script(script)
//			bulkRequest = bulkRequest.Add(r)
//		case BulkDelete:
//			r := elastic.NewBulkDeleteRequest().Index(req.IndexName).Id(req.DocumentID)
//			bulkRequest = bulkRequest.Add(r)
//		}
//	}
//
//	if bulkRequest.NumberOfActions() != len(reqs) {
//		return fmt.Errorf("expected bulkRequest.NumberOfActions %d; got %d", 3, bulkRequest.NumberOfActions())
//	}
//	ret, err := bulkRequest.Do(context.Background())
//	if err != nil {
//		return err
//	}
//	if ret.Errors {
//		logx.NewTraceLogger(context.Background()).Error().Interface("res", ret).Msg("批量操作ERR")
//		if len(ret.Items) > 0 {
//			item := ret.Items[0]
//			if resp, ok := item["index"]; ok && resp.Error != nil {
//				return errors.New(resp.Error.Reason)
//			}
//		}
//		return errors.New("unknown error")
//	}
//	return nil
//}
//
////Search 查询
//func (e *ElasticClient) Search(req SearchRequest) (map[string]interface{}, error) {
//	var (
//		searchResult *elastic.SearchResult
//		err          error
//	)
//
//	var searchService = e.client.Search().Index(req.IndexName)
//	for key, value := range req.QueryParameter {
//		searchService = searchService.Query(elastic.NewTermQuery(key, value))
//	}
//	for _, query := range req.ElasticQuerys {
//		searchService = searchService.Query(query)
//	}
//	for key, value := range req.SortFields {
//		searchService = searchService.Sort(key, value)
//	}
//	if len(req.SortByFields) > 0 {
//		searchService = searchService.SortBy(req.SortByFields...)
//	}
//	if req.PageSize == 0 {
//		req.PageSize = 20
//	}
//	searchService = searchService.From(req.PageOffset).Size(req.PageSize).
//		FetchSourceContext(req.FetchSource)
//	searchResult, err = searchService.Do(context.Background())
//
//	if elastic.IsNotFound(err) {
//		return ElasticEmptySearchResponse, nil
//	}
//
//	if err != nil && err != io.EOF {
//		err = e.handleElasticError(err, req.DocumentID)
//		return nil, err
//	}
//	var (
//		buf       bytes.Buffer
//		sourceLen = len(searchResult.Hits.Hits)
//		r         = make(map[string]interface{})
//	)
//	buf.WriteString("[")
//	for i := 0; i < sourceLen; i++ {
//		source := searchResult.Hits.Hits[i]
//		if source.Source != nil {
//			if i > 0 {
//				buf.Write([]byte(","))
//			}
//			buf.Write(source.Source)
//		}
//	}
//	buf.WriteString("]")
//	r["scroll_id"] = searchResult.ScrollId
//	r["data"] = buf.Bytes()
//	return r, nil
//}
//
////Count 查询符合条件的数据总量
//func (e *ElasticClient) Count(req SearchRequest) (int64, error) {
//	countService := e.client.Count(req.IndexName)
//	for key, value := range req.QueryParameter {
//		countService = countService.Query(elastic.NewTermQuery(key, value))
//	}
//	for _, query := range req.ElasticQuerys {
//		countService = countService.Query(query)
//	}
//	num, err := countService.Do(context.Background())
//	if elastic.IsNotFound(err) {
//		return 0, nil
//	}
//	return num, err
//}
//
//// Refresh 刷新索引
//func (e *ElasticClient) Refresh(IndexName ...string) error {
//	_, err := e.client.Refresh(IndexName...).Do(context.Background())
//	return err
//}
//
//func (e *ElasticClient) handleElasticError(err error, documentID string) error {
//	switch {
//	case elastic.IsNotFound(err):
//		err = fmt.Errorf("document %s not found", documentID)
//	case elastic.IsTimeout(err):
//		err = errors.New("retrieve document timeout")
//	case elastic.IsConnErr(err):
//		err = errors.New("connection problem")
//	}
//	zslog.Debug().Str("document_id", documentID).Str("type", "notice").Err(err).Msg("获取文档失败")
//	return err
//}
//
//func (e *ElasticClient) interfaceToBytes(body interface{}) ([]byte, error) {
//	var (
//		bodyBytes []byte
//		err       error
//	)
//	switch t := body.(type) {
//	case []byte:
//		bodyBytes = t
//	case string:
//		bodyBytes = []byte(t)
//	default:
//		bodyBytes, err = json.Marshal(t)
//		if err != nil {
//			return nil, err
//		}
//	}
//	return bodyBytes, err
//}
