# Automatically update the A record of your cloudflare domains with your public IPv4 address

- Gets the external/public IPv4 Address
- Uses the config.yaml to get the auth email, cloudflare token and list of
  domains
- Updates each domain's A record with the public IPv4 address
- Supports domains and subdomains, i.e. if you want to update the domain only, set the domain and name to the same value in the config.
- If you want to update only a subdomain and not the domain, set the name and domain to a different value, i.e. if your domain is `"example.com"` and you want to udpate |`"subdomain.example.com"` set the latter as name and the former as domain in the `config.yaml`

Same `config.yaml`

```
domains:
  - domain: "example0.com"
    name: "example0.com"
    zone: "14188253e4f00003d5d45e03pp0ppp23"
    proxied: false
  - domain: "example1.com"
    subdomain: "example1"
    zone: "14188253e4f00003d5d45e03pp0ppp23"
    proxied: false
  - domain: "example2.com"
    name: "example2"
    zone: "14188253e4f00003d5d45e03pp0ppp23"
    proxied: false
  - domain: "example3.com"
    name: "example3"
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
2023/07/21 12:33:31 IPV4 address hasn't changed, 200.22.20.15 = 200.22.20.15 in example0.com
2023/07/21 12:33:31 IPV4 address hasn't changed, 200.22.20.15 = 200.22.20.15 in example1.example1.com
2023/07/21 12:33:32 IPV4 address hasn't changed, 200.22.20.15 = 200.22.20.15 in example2.example2.com
```

Motivation: I have a dynamic IP address at home and every time the electricity
goes out (used too much power), I have to manually update the A records of all
domains hosted at home with cloudflare.
