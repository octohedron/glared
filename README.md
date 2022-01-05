# Automatically update the A record of your cloudflare domains with your public IPv4 address

- Gets the external/public IPv4 Address
- Uses the config.yaml to get the auth email, cloudflare token and list of
  domains
- Updates each domain's A record with the public IPv4 address

Same `config.yaml`

```
domains:
  - name: "example0.com"
    zone: "14188253e4f00003d5d45e03pp0ppp23"
    proxied: false
  - name: "example1.com"
    zone: "14188253e4f00003d5d45e03pp0ppp23"
    proxied: false
  - name: "example2.com"
    zone: "14188253e4f00003d5d45e03pp0ppp23"
    proxied: false
  - name: "example3.com"
    zone: "14188253e4f00003d5d45e03pp0ppp23"
    proxied: false
auth:
  key: "14188253e4f00003d5d45e03pp0ppp23"
  email: "your-email@example.com"
```

Usage:

`go build && ./glared`

Should print something like

```
2022/01/05 22:22:54 IPV4 address 200.157.52.125
2022/01/05 22:22:56 Updated example0.com
2022/01/05 22:22:58 Updated example1.com
2022/01/05 22:23:00 Updated example2.com
2022/01/05 22:23:03 Updated example3.com
```

Motivation: I have a dynamic IP address at home and every time the light goes
out for any reason, i.e. using too much electricity, I have to manually update
the A records of all domains hosted in my home server in cloudflare. This avoids
having to do it manually, just need to run it on boot.
