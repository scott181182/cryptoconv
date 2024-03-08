use std::num::ParseFloatError;
use std::process::exit;
use std::env;

// This library helps create and handle sum types (enums) for errors.
use thiserror::Error;

mod rates;



fn print_usage() {
	println!("
USAGE
cryptoconv <USD amount> <currency 1> <currency 2>
	Converts 70% of USD amount to currency 1, and 30% to currency 2\
    ");
}



#[derive(Error, Debug)]
enum CryptoConvError {
    #[error("Expected 3 arguments, but found {0}")]
    UnexpectedArguments(usize),
    #[error("There was an error parsing your USD amount. Please supply a number.")]
    AmountParse(#[from] ParseFloatError),
    #[error("There was an error getting the current exchange rates:\n{0}")]
    ExchangeRate(#[from] rates::ExchangeRateError)
}



fn log_conversion(usd: f64, currency: &str, rates: &rates::ExchangeRates) {
    let upper_currency = currency.to_uppercase();

    if let Some(rate) = rates.get(&upper_currency) {
        println!("${:.2} USD => {} {}", usd, usd * rate, upper_currency);
    } else {
        println!("Could not find an exchange rate for '{}'. Please try with another.", upper_currency);
    }
}

fn cryptoconv() -> Result<(), CryptoConvError> {
    let args: Vec<String> = env::args().collect();

    if args.len() != 4 {
        return Err(CryptoConvError::UnexpectedArguments(args.len()))
    }

    let usd = args[1].parse::<f64>()?;
	let currency1 = &args[2];
	let currency2 = &args[3];

    // You wouldn't traditionally panic, but this is simple enough for a CLI application.
    let rates = rates::get_exchange_rates()?;

	log_conversion(usd * 0.7, currency1, &rates);
	log_conversion(usd * 0.3, currency2, &rates);

    Ok(())
}

fn main() {
    if let Err(err) = cryptoconv() {
        eprintln!("{}", err);
        print_usage();
        exit(1);
    }
}
