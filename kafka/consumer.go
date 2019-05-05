package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Consumer() {
	//创建消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2
	consumer, err := sarama.NewConsumer([]string{"192.168.88.95:9092"}, config)
	if err != nil {
		log.Fatal("创建消费者失败" + err.Error())
	}
	//获取所有的主题
	topics, _ := consumer.Topics()
	name := topics[len(topics)-1]
	//设置分区
	partitionlist, err := consumer.Partitions(name)
	if err != nil {
		log.Fatal("创建分区失败" + err.Error())
	}
	//循环分区
	for partion := range partitionlist {
		fmt.Println(partion)
		pc, err := consumer.ConsumePartition(name, int32(partion), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("无法启动消费则分区" + err.Error())
			return
		}
		defer pc.AsyncClose()
		go func(pc sarama.PartitionConsumer) {
			wg.Add(1)
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				fmt.Println()
			}
			wg.Done()
		}(pc)
	}
	//wg可能还没有赋值就结束了，原因可能是pc的内容过多
	time.Sleep(1 * time.Second)
	wg.Wait()
	consumer.Close()
	//fmt.Printf("mytopic")
	//
	//config := sarama.NewConfig()
	//config.Consumer.Return.Errors = true
	//config.Version = sarama.V0_11_0_2
	//
	//// consumer
	//consumer, err := sarama.NewConsumer([]string{"192.168.88.95:9092"}, config)
	//if err != nil {
	//	fmt.Printf("consumer_test create consumer error %s\n", err.Error())
	//	return
	//}
	//
	//defer consumer.Close()
	//
	//partition_consumer, err := consumer.ConsumePartition("mytopic", 0, sarama.OffsetOldest)
	//if err != nil {
	//	fmt.Printf("try create partition_consumer error %s\n", err.Error())
	//	return
	//}
	//defer partition_consumer.Close()
	//
	//for {
	//	select {
	//	case msg := <-partition_consumer.Messages():
	//		fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
	//			msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
	//	case err := <-partition_consumer.Errors():
	//		fmt.Printf("err :%s\n", err.Error())
	//	}
	//}
}

func main() {
	Consumer()
}
