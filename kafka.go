package kafka

import (
	"context"
	"time"

	"github.com/Shopify/sarama"
	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

type KafkaParams struct {
}

type Kafka struct {
	Next plugin.Handler
	// producer *sarama.AsyncProducer
	hosts    string
	username string
	password string
}

// Name implements plugin.Handler
func (Kafka) Name() string {
	return "kafka"
}

func (k Kafka) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	state := request.Request{W: w, Req: r}
	qname := state.QName()
	return plugin.NextOrFailure(qname, k.Next, ctx, w, r)

}

func (k *Kafka) OnStartup() error {
	config := sarama.NewConfig()
	// config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second
	config.Producer.Flush.Messages = 100
	config.Producer.Flush.Frequency = 10 * time.Second
	config.Net.SASL.Enable = true
	config.Net.SASL.User = k.username
	config.Net.SASL.Password = k.password
	// k.producer, err = sarama.NewAsyncProducer(strings.Split(k.hosts, ","), config)
	return nil
}
