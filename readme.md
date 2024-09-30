# Current Conditions and Daily Weather Forecast

This is just a simple cli app that will output the current local conditions and daily hourly forecast. If ran without an input, it will default to Tampa, FL. 

## How to Use

```
weather 90210
```

Will give the current conditions and hourly forecast based on zipcode. The zipcode may also be replaced with a city name. For example:

```
weather Tampa
```

## API Key

Sign up for a free api key at (https://www.weatherapi.com/). Once you have your key, create a config.json file and set it up as follows:

```
{
    "key" = "xxxx"
}
```
