## RabbitMQ with Golang

``` shell
docker run -d --hostname=my-rabbit --name=some-rabbit
```

![RabbitMQ Architecture](https://lh3.googleusercontent.com/TmA6flkGzB1yc1xK6lGbJZ0YYqO__39trLIPxM62VUjsr09wClmbv9mT3WX4F0cuDssmkiHkekWR6AvXY0iVScuksmLxyM27FaJGYbgPezCIjRs-l8Ct3MfuUU3bRbpfWT6dhVBO)

### Concepts:

1. Broker
2. Message Acknowledgements
3. Message can be pushed or pulled.
4. Dead letter Queue
5. Exchange
   1. Direct Exchange - Empty String and amq.Direct
   2. Fanout Exchange - amq.fanout
   3. Topic Exchange - amq.topic
   4. Headers Exchange - amq.match (and amq.headers in RabbitMQ)

Besides Exchange types, exchanges are declared with a number of attributes, the most important of which are:

1. Name
2. Durability (exchanges survive broker restart)
3. Auto-delete (exchange is deleted when last queue is unbound from it)
4. Arguments (optional, used by plugins and broker-specific features)

Decoupling producers from queues via exchanges ensures that producers aren't burdened with hardcoded routing decisions.

### Prerequisites
Library we are using `amqp`

``` shell
go get -u github.com/streadway/amqp
```

### Files
1. sender/publisher - `publisher.go`
2. consumer/worker/receiver/subscriber - `receiver.go`

### Code Snippets
Helper Function - `failOnError`


``` go
func failOnError(err error, msg string) {
    if err != nil {
            log.Fatal("%s: %s", msg, err)
    }
}
```

Connect to RabbitMQ Server

``` go
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()
```

Create a channel, which is where most of the API for getting things done resides

``` go
ch, err := conn.Channel()
failOnError(err, "Failed to open a channel")
defer ch.Close()
```

To send, we must declare a queue for us to send to; then we can publish a message to the queue:

``` go
q, err := ch.QueueDeclare(
  "hello", // name ->
  false,   // durable ->
  false,   // delete when unused ->
  false,   // exclusive ->
  false,   // no-wait ->
  nil,     // arguments ->
)
```
n
