use std::{collections::HashMap, num::ParseFloatError};

// Our HTTP client library.
use reqwest::Error as ReqwestError;
// Used for deserializing/parsing the JSON response automatically.
use serde::Deserialize;
// Used to construct sum types (enums) for errors, making them much easier to handle and pass around.
use thiserror::Error;



pub type ExchangeRates = HashMap<String, f64>;

#[derive(Error, Debug)]
pub enum ExchangeRateError {
    #[error(transparent)]
    RequestError(#[from] ReqwestError),
    #[error(transparent)]
    ParseFloatError(#[from] ParseFloatError)
}



/// Data structure for parsing the Coinbase API response.
#[derive(Deserialize)]
struct CoinbaseData {
    // Keeping this since the JSON structure has it, even though we don't use it.
    #[allow(dead_code)]
    pub currency: String,
    pub rates: HashMap<String, String>
}
#[derive(Deserialize)]
struct CoinbaseResponse {
    pub data: CoinbaseData
}



const COINBASE_URI: &str = "https://api.coinbase.com/v2/exchange-rates?currency=USD";

fn fetch_exchange_rates(uri: &str) -> Result<CoinbaseResponse, ExchangeRateError> {
    Ok(reqwest::blocking::get(uri)?.json::<CoinbaseResponse>()?)
}

fn parse_exchange_rate_entry((key, value): (String, String)) -> Result<(String, f64), ExchangeRateError> {
    Ok((key, value.parse::<f64>()?))
}
fn parse_exchange_rates(res: CoinbaseResponse) -> Result<ExchangeRates, ExchangeRateError> {
    res.data.rates.into_iter()
        .map(parse_exchange_rate_entry)
        .collect::<Result<ExchangeRates, ExchangeRateError>>()
}

pub fn get_exchange_rates() -> Result<ExchangeRates, ExchangeRateError> {
    fetch_exchange_rates(COINBASE_URI)
        .and_then(parse_exchange_rates)
}



#[cfg(test)]
mod test {
    use std::collections::HashMap;

    use crate::rates::parse_exchange_rates;

    use super::{CoinbaseData, CoinbaseResponse, ExchangeRates};

    #[test]
    fn test_rate_parse() {
        let input = CoinbaseResponse {
            data: CoinbaseData{
                currency: "USD".to_owned(),
                rates: HashMap::from([
                    ("BTC".to_owned(), "5.40".to_owned()),
                    ("ETH".to_owned(), "1.0".to_owned())
                ])
            }
        };
        let expected = ExchangeRates::from([
            ("ETH".to_owned(), 1f64),
            ("BTC".to_owned(), 5.4f64)
        ]);

        assert_eq!(parse_exchange_rates(input).unwrap(), expected);
    }
}
