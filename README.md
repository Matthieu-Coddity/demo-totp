# ToTP demo with Go

This quick demo of T-OTP is described in this [coddity blog post](https://blog.coddity.com/articles/totp).

To run:
- install Go
- `$go mod tidy`
- `$go run generate` to generate uri and qr code
- `$go run run` to generate authentication codes, changing every 30'
- `$go run verify <CODE_2_VERIFY>` to verify code