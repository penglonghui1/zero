package nsqueue

import (
	"testing"
	"time"

	sonyflake "github.com/pengcainiao/zero/core/snowflake"
	"github.com/pengcainiao/zero/rest/httprouter"
	"github.com/youzan/go-nsq"
)

type message struct {
	taskID  string
	message string
}

type testConsumerHandler struct {
}
type test1ConsumerHandler struct {
}

func (t *testConsumerHandler) HandleMessage(message *nsq.Message) error {
	println("[test0] Receive new message ...")
	println("[test0] Body is: ", string(message.Body))
	message.Finish()
	return nil
}

func (t *test1ConsumerHandler) HandleMessage(message *nsq.Message) error {
	println("[test1] Receive new message ...", message.GetTraceID())
	println("[test1] Body is: ", string(message.Body))
	message.Finish()
	return nil
}

func TestNsq_NewConsumer(t *testing.T) {
	var data = []message{
		{
			taskID:  "760711755595937",
			message: "zouyige",
		},
		{
			taskID:  "760711755595937",
			message: "试试看",
		},
		{
			taskID:  "760711755595937",
			message: "shme ???",
		},
		{
			taskID:  "760711755595939",
			message: "zouyige----1",
		},
		{
			taskID:  "760711755595940",
			message: "试试看----1",
		},
		{
			taskID:  "760711755595941",
			message: "shme ???---1",
		},
	}
	var (
		topic = &Topic{
			Name:          "flyele-lookup-socket-consumer",
			EnableOrdered: false,
			Channel:       "default_channel",
		}
		topic1 = &Topic{
			Name:          "flyele-lookup-socket-consumer-1",
			EnableOrdered: true,
			Channel:       "default_channel",
		}
	)
	topic = topic.Priority(NoOrderLow)
	lookup := NewNsqLookup()
	go func() {
		pubMgr := lookup.NewProducer(topic, topic1)
		if pubMgr == nil {
			t.Log("NewProducer failed")
			t.Failed()
			return
		}
		var pn int
		for idx, datum := range data {
			if pn == 0 {
				pn = 1
			} else {
				pn = 0
			}
			_ = lookup.Publish(topic, NsqDataProtocol{
				Topic:   topic.Name,
				Body:    []byte(datum.message),
				TraceID: uint64(idx),
			})

			_ = lookup.Publish(topic1, NsqDataProtocol{
				Topic:     topic1.Name,
				Body:      []byte(datum.message),
				TraceID:   uint64(idx),
				OrderedBy: datum.taskID,
			})

		}
	}()
	<-time.After(time.Second * 5)
	go func() {
		err := lookup.NewConsumer(topic, &testConsumerHandler{})
		if err != nil {
			t.Error(err)
			t.Failed()
		}
	}()
	go func() {
		err := lookup.NewConsumer(topic, &test1ConsumerHandler{})
		if err != nil {
			t.Error(err)
			t.Failed()
		}
	}()
	select {}
}

func TestTopicPublishComsule(t *testing.T) {
	var topic = &Topic{
		Name:          "nsq-test-topic",
		EnableOrdered: false,
		Channel:       "default_channel",
	}
	var ta = Topics{topic}
	_ = ta.SetProducerManager()
	// 发送数据
	_ = topic.Publish(NsqDataProtocol{
		Topic:   topic.Name,
		Body:    map[string]interface{}{"abc": 123, "b": "454545"},
		TraceID: sonyflake.GenerateInt64ID(),
	})
	//消费数据
	_ = topic.StartConsume(&testConsumerHandler{})
}

type TestConsumeData struct {
	UserInfo
	Service
}

type UserInfo struct {
	Addr string `json:"addr"`
	Age  int    `json:"age"`
}

type Service struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

func (t *TestConsumeData) Bytes() []byte {
	//t.Addr = "广州市黄埔区"
	b, _ := json.Marshal(t)
	return b
}

func (t TestConsumeData) New() MqMessageConsumer {
	return &TestConsumeData{}
}

func (t TestConsumeData) HandleMessage(ctx *httprouter.Context, protocol *NsqDataProtocol) error {

	println(t.Addr, "    ", t.Age, "   ", t.Namespace)
	return nil
}

//nolint
func getNsqData() NsqDataProtocol {
	var d = `{"topic":"flyele-nsq-pushgateway-multicast-release","body":{"websocket":{"receiver":[{"receiver_id":"1238663641694355","affected":"1"},{"receiver_id":"1249909260419181","affected":"1"},{"receiver_id":"1191142398361793","affected":"1"}],"ref_id":"100000000000001","ref_type":"1","code":32,"message_id":"1267923357532382","creator_id":"1238663641694355","comment_id":"1267923356483753","message":"你的【测试详情事项】事项将于15分钟后截止","send_from":"1","message_type":6,"subtitle":"事项截止提醒"}},"header":{"data":{},"create_at":"2021-11-16 19:23:57"}}`
	var v NsqDataProtocol
	json.UnmarshalFromString(d, &v)
	return v
}

func TestConsumeNSQ(t *testing.T) {
	var (
		w = wrappedMqMessageConsumerHandler{
			Topic:    &Topic{Name: "xxxx33333"},
			Consumer: &TestConsumeData{},
		}
		d = NsqDataProtocol{
			Topic: "xxxx2222",
			Body: TestConsumeData{UserInfo: UserInfo{
				Addr: "fdsfsadfasd", Age: 10003,
			}},
			TraceID:   234234230,
			OrderedBy: "",
			Header:    MessageHeader{},
		}
		d1 = NsqDataProtocol{
			Topic:     "xxxx11",
			Body:      TestConsumeData{Service: Service{Namespace: "feixiang1"}},
			TraceID:   234234230,
			OrderedBy: "",
			Header:    MessageHeader{},
		}
	)
	//_ = syncer.QueueDataSyncer().Consume(TestConsumeData{}, conf.WithMaxConcurrency(400))
	_ = w.HandleMessage(&nsq.Message{
		ID:   nsq.MessageID{},
		Body: d.Bytes(),
	})
	_ = w.HandleMessage(&nsq.Message{
		ID:   nsq.MessageID{},
		Body: d1.Bytes(),
	})
}
