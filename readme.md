# Current Conditions and Daily Weather Forecast

This is just a simple cli app that will output the current local conditions and daily hourly forecast. If ran without an input, it will default to Tampa, FL. 

## API Key

Sign up for a free api key at (https://www.weatherapi.com/). 

## Setup

1. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` and add your real API key:
   ```bash
   WEATHER_API_KEY=your_actual_api_key_here
   ```

The `.env` file is already in `.gitignore` so your API key won't be committed to version control.

## Installation

### Option 1: Install with Go (Recommended)
```bash
go install .
```
This installs the `weather` command globally, making it available from any terminal location.

### Option 2: Build and Install Manually
```bash
go build -o weather .
sudo mv weather /usr/local/bin/
```

### Option 3: Build for Local Development
```bash
go build -o weather .
./weather Tampa
```

**Note:** For Option 1 to work globally, make sure `$HOME/go/bin` is in your PATH. Add this to your shell profile if needed:
```bash
export PATH="$HOME/go/bin:$PATH"
```

## How to Use

```
weather 90210
```

Will give the current conditions and hourly forecast based on zipcode. The zipcode may also be replaced with a city name. For example:

```
weather Tampa
```

