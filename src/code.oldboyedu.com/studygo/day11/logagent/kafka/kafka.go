package kafka

import(
	_ "github.com/shopify/sarama"
)

//专门往kafka写日志的模块

var (
	client sarama.SyncProducer	//声明一个全局的连接kafka的生产者client
)

//Init 初始化client
func Init(err error){
	config:=sarama.NewConfig()
	//tailf包使用
	config.Producer.RequiredAcks=sarama.WaitForAll
	config.Producer.Partitioner=sarama.NewRandomPartitioner
	config.Producer.Return.Successes=true
	//连接kafka
	client,err=sarama.NewSyncProducer(addrs,config)
	if err!=nil{
		fmt.Println("produce close err:",err)
		return 
	}
	return 
}


func SendToKafka(topic, data string){
	//构造一个消息
	msg:=sarama.ProducerMessage{}
	msg.Topic=topic
	msg.Value=sarama.StringEncoder(data)
	//发送到kafka
	pid,offset,err:=client.SendMessage(msg)
	fmt.Println("xxx")
	if err!=nil{
		fmt.Println("send msg failed,err:",err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n",pid,offset)
}