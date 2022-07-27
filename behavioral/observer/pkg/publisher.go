package pkg

type Publisher struct {
	Name      string
	Consumers map[string]Consumer
}

func (pub *Publisher) Subscribe(consumer Consumer) *Publisher {
	pub.Consumers[consumer.GetName()] = consumer
	return pub
}

func (pub *Publisher) UnSubscribe(consumer Consumer) *Publisher {
	delete(pub.Consumers, consumer.GetName())
	return pub
}

func (pub *Publisher) Notify() {
	for _, consumer := range pub.Consumers {
		consumer.Update(pub.Name)
	}
}
