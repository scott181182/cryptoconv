# Cryptocurrency Converter

This repository contains roughly-equivalent CLI tools for converting USD to different cryptocurrencies, utilizing Coinbase's API according to the below prompt.

## Prompt

This endpoint provides up-to-the-minute crypto exchange rates relative
to US dollars: <https://api.coinbase.com/v2/exchange-rates?currency=USD>

That is: each rate is how much of that crypto currency you would get
for 1 dollar. So if you received a value for 0.091 for BTC, that's
saying it's 0.091 per 1 USD.

Your Task:
You are to make a cli that takes in a USD amount as holdings, and
calculates the 70/30 split for 2 given crypto currencies. Stated
simply: I have $X I want to keep in BTC and ETH, 70/30 split. How many
of each should I buy? An example usage would look like:

```raw
> binary_name 100 BTC ETH
$70.00 => 0.0025 BTC
$30.00 => 0.0160 ETH
```

This output tells us: Of our $100 holdings, 70% of that is $70, which
buys 0.0025 BTC, and 30% of our holdings is $30, which buys 0.016 ETH.

## Languages

While the original prompt was specifically for a Go program, I have taken the liberty to setup this repository to handle multiple languages, with pertinent instructions in each subfolder. Enjoy!
