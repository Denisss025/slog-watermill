# slog: watermill logger adapter

Watermill logger adapter for slog logger

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.21-%23007d9c)
[![GoDoc](https://godoc.org/github.com/denisss025/slog-watermill?status.svg)](https://pkg.go.dev/github.com/denisss025/slog-watermill)
![Build Status](https://github.com/denisss025/slog-watermill/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/denisss025/slog-watermill)](https://goreportcard.com/report/github.com/denisss025/slog-watermill)
[![Coverage](https://img.shields.io/codecov/c/github/denisss025/slog-watermill)](https://codecov.io/gh/denisss025/slog-watermill)
[![License](https://img.shields.io/github/license/denisss025/slog-watermill)](./LICENSE)

[Watermill](https://github.com/ThreeDotsLabs/watermill) LoggerAdapter implementation to log using [slog](https://pkg.go.dev/log/slog).

**See also:**

- [slog-multi](https://github.com/samber/slog-multi): `slog.Handler` chaining, fanout, routing, failover, load balancing...
- [slog-formatter](https://github.com/samber/slog-formatter): `slog` attribute formatting
- [slog-sampling](https://github.com/samber/slog-sampling): `slog` sampling policy
- [slog-gin](https://github.com/samber/slog-gin): Gin middleware for `slog` logger
- [slog-echo](https://github.com/samber/slog-echo): Echo middleware for `slog` logger
- [slog-fiber](https://github.com/samber/slog-fiber): Fiber middleware for `slog` logger
- [slog-datadog](https://github.com/samber/slog-datadog): A `slog` handler for `Datadog`
- [slog-rollbar](https://github.com/samber/slog-rollbar): A `slog` handler for `Rollbar`
- [slog-sentry](https://github.com/samber/slog-sentry): A `slog` handler for `Sentry`
- [slog-syslog](https://github.com/samber/slog-syslog): A `slog` handler for `Syslog`
- [slog-logstash](https://github.com/samber/slog-logstash): A `slog` handler for `Logstash`
- [slog-fluentd](https://github.com/samber/slog-fluentd): A `slog` handler for `Fluentd`
- [slog-graylog](https://github.com/samber/slog-graylog): A `slog` handler for `Graylog`
- [slog-loki](https://github.com/samber/slog-loki): A `slog` handler for `Loki`
- [slog-slack](https://github.com/samber/slog-slack): A `slog` handler for `Slack`
- [slog-telegram](https://github.com/samber/slog-telegram): A `slog` handler for `Telegram`
- [slog-mattermost](https://github.com/samber/slog-mattermost): A `slog` handler for `Mattermost`
- [slog-microsoft-teams](https://github.com/samber/slog-microsoft-teams): A `slog` handler for `Microsoft Teams`
- [slog-webhook](https://github.com/samber/slog-webhook): A `slog` handler for `Webhook`
- [slog-kafka](https://github.com/samber/slog-kafka): A `slog` handler for `Kafka`
- [slog-parquet](https://github.com/samber/slog-parquet): A `slog` handler for `Parquet` + `Object Storage`

## 🚀 Install

```sh
go get github.com/denisss025/slog-watermill
```

**Compatibility**: go >= 1.21

No breaking changes will be made to exported APIs before v2.0.0.

## 💡 Usage

### Minimal

```go

```

### Using custom time formatters

```go

```

### Using custom logger sub-group

```go

```

### Adding custom attributes

```go

```

### JSON output

```go

```

## 🤝 Contributing

- Fork the [project](https://github.com/denisss025/slog-watermill)
- Fix [open issues](https://github.com/denisss025/slog-watermill) or request new features

Don't hesitate ;)

```bash
# Install some dev dependencies
make tools

# Run tests
make test
# or
make watch-test
```

## 👤 Contributors

![Contributors](https://contrib.rocks/image?repo=denisss025/slog-watermill)

## 💫 Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2023 [Denis Novikov](https://github.com/denisss025).

This project is [MIT](./LICENSE) licensed.
