# Go Crawler

This is example sending an crawling message to slack.

## Requirements

- go version: >= 1.11

## File tree

```
.
├── LICENSE
├── README.md
├── cmd
│   └── root.go
├── go.mod
├── go.sum
├── internal
│   └── bot.go
└── main.go
```

## Steps

1. Install dependencies.

```sh
$ go get
```

2.  Create a config file.

- file path: ./.go-crawler.yaml

```yaml
crawling_url: CRAWLING_URL # ex) "https://www.daangn.com/"
selector: DOM_SELECTOR # ex) "#copyright"
reciever: SLACK_WEBHOOK_URL # "https://hooks.slack.com/services/XXX/XXX/XXX"
```

3. Run

```
$ go run main.go
```

# License

[MIT License](./LICENSE)