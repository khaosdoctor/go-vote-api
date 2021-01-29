# Go voting api

> A complete simple voting API in 5MB

This docker image is intended to be a simple voting API, ReST concepts and good practices **were intentionally left aside**.

## Running

```bash
docker run -p 80:8080 khaosdoctor/go-vote-api
```

> The API **ALWAYS** listens on port 8080

## Routes

- `GET /votes/:name`: Computes a vote to the given name
- `GET /votes`: List all votes
- `GET /total`: Lists the total number of votes and the vote list

> The API is concurrency-safe but not scalable since it stores all votes in local memory
