# ToTP demo with Go

to run:
- install Go
- `$go mod tidy`
- `$go run generate` to generate uri and qr code
- `$go run run` to generate authentication codes, changing every 30'
- `$go run <CODE_2_VERIFY>` to verify code