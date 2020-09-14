# redis_sample_app
go sample app that connected to redis

# how to run:
1. Run `go get "github.com/garyburd/redigo/redis"`
2. Run `make run`

# testing
```
curl -X POST -d '{"prizes":[{"name":"ovo 1k","percentage":60},{"name":"ovo 2k","percentage":20},{"name":"ovo 3k","percentage":10}]}' http://localhost:8181/create_prize_pool

curl localhost:8181/get_prize?user_id=123
```
