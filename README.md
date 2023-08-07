# Go http server demo

### make file .env

```
PORT=8080
```



### run

```sh
go build && ./start
```



### test

API (get)

```
http://localhost:8080/v1/err
```

Get 200

```
{
    "error": "Something went wrong."
}
```



API (get)

```
http://localhost:8080/v1/healthz
```

Get 200

```
{}
```

