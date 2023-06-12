# realtime-stream-processor

## Technologies Used

1. [Apache Kafka]() - Used for streaming data from source to destination. More about kafka:

   - [Docs](https://kafka.apache.org/documentation/)
   - [TL;DR Intro Playlist](youtube.com/playlist?list=PLa7VYi0yPIH2PelhRHoFR5iQgflg-y6JA)

2. [Temporal](https://temporal.io/) - It is a scalable and reliable runtime for Reentrant Processes. It's battle-tested at scale and used in production by companies like uber, netflix, snap (snapchat) etc. More about temporal:

   - [Docs](https://docs.temporal.io/temporal)
   - [TL;DR Intro Video](https://www.youtube.com/watch?v=2HjnQlnA5eY)

3. [PostgreSQL](https://www.postgresql.org/) - Used for persistance of workflows alongside temporal.

4. [MongoDB](https://www.mongodb.com/) - Used for persistance of events after post processing

## How to run

In the main directory, run the following command:

```
docker compose -f docker-compose.yaml up -d
```

This will spin up the technologies mentioned above. After this, you can run the following commands to start the consumer application:

```
go mod download
go run main.go
```

To mock the data for the application, you can run the following command:

```
go run mocker/main.go
```

## Monitoring Interfaces

1. Temporal UI :- https://localhost:8088
2. MongoDB UI :- https://localhost:8081

## Improvements

1. Setup env variables for the application
2. Implement retry logic for kafka consumer when it crashes (we need to not ack the messages so that they can be retried), another approach is to use dead letter queue but since we are using temporal, our main queue will only crash in cases when either kafka or temporal is down, so we can block the main queue.
