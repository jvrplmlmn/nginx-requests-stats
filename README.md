# nginx-requests-stats

The `nginx-requests-stats` application reads a Nginx log, and counts the number of requests in the last 24 hours.

The application binds to `localhost:8080` and exposes the following HTTP endpoints:

```
GET /version

{
  "version": "0.1.0"
}
```

```
GET /count

{
  "requests": 15
}
```
