# Vaultd

Edit and share secrets between humans and machines.

```bash
> curl io/help | jq .
[
  "/list/*path      List secrets",
  "/get/*name?mode  Get secrets (default mode: decrypt)"
]
```

#### List secrets

```bash
> curl io/ls | jq .
[
  "a/b/kr.json",
  "a/e.env",
  "a/j.json",
  "a/y.yaml"
]
```

#### Get secrets

```bash
# Environment vars

> cat test/a/e.env.encrypt
X=323339383332303331323339383332307a73
API=3233393835303039323233393835303080fdea0267fdd9551d696cdee7abd0ad4ca713742c

> curl io/get/a/1.secrets.env
X=42
API="http://miaou.cat/?a"
```

```bash
# JSON

> cat test/a/j.json.encrypt
{"bim":{"A":["38363934333132393138363934333132e03d"],"B":["3836393430353337323836393430353391a297ef"],"D":{"E":{"F":["38363934313734323038363934313734114889"],"G":[{"here":"3836393432323731373836393432323774530d","plop":"3836393432363532393836393432363516b453"}]}}}

> curl io/get/a/j.json | jq '.bim.D.E.G[0]'
{
  "here": "111",
}
```

```bash
# YAML

> cat test/a/y.yaml.encrypt
bim:
  A:
  - 383230353036333530383230353036338ccd
  B:
  - 38323035313932383238323035313932d5e25585
  D:
    E:
      F:
      - 38323035323539333138323035323539432dc3
      G:
      - here: 383230353331333538383230353331336bec8c
        plop: 38323035333534383738323035333534818005

> curl io/get/a/y.yaml
bim:
  A:
  - "42"
  B:
  - "1337"
  D:
    E:
      F:
      - "123"
      G:
      - here: "456"
        plop: "678"
```
