package client

import "testing"

func TestKafkaSyncSendReturnsErrorWhenProducerNotInitialized(t *testing.T) {
	originalProducer := AsyncProducer
	AsyncProducer = nil
	defer func() {
		AsyncProducer = originalProducer
	}()

	if err := KafkaSyncSend("topic", "key", []byte("payload")); err == nil {
		t.Fatal("expected error when kafka producer is not initialized")
	}
}
