# Intinerary Prettifier

## Project overview:

This project is a command-line tool that processes an input text file containing flight itineraries and formats it into a more readable output.

The tool performs the following transformations:

1. Replaces airport codes:
    * #IATA → full airport name
    * ##ICAO → full airport name
    * *#CODE → city name instead of airport name
2. Formats date and time expressions:
    * D(...) → formatted date
    * T12(...) → 12-hour time format
    * T24(...) → 24-hour time format
3. Normalizes whitespace:
* Converts all vertical whitespace to \n
* Removes excessive blank lines

### Bonus features:
1. Text preview which highlights dates, times, offsets, airport names, cities in terminal
2. City names: When a city name is desired instead of an airport name, the encoding will be preceded by a * symbol, so that *#LHR is converted to "London" (For London Heathrow Airport).
3. Support dynamic airport lookup column order

## Examples:  

Input:  
```
Your flight departs from #HAJ, and your destination is ##EDDW.  
```

Results:  
```
Your flight departs from Hannover Airport, and your destination is Bremen Airport.  
```

```
Your flight departs from *#HAJ, and your destination is *##EDDW.  
```

Results:  
```
Your flight departs from Hannover, and your destination is Bremen.  
```

Input:  
```
1. D(2022-05-09T08:07Z)
2. T12(2069-04-24T19:18-02:00)
3. T24(2032-07-17T04:08+13:00)  
```

Results:  
```
1. 09 May 2022
2. 07:18PM (-02:00)
3. 04:08 (+13:00)
```

## Setup and Installation
1. Clone the repository:  
```bash
git clone https://gitea.kood.tech/kimhangngo/prettifier  
cd prettifier
```

2. Ensure Go is installed  
```bash    
    go version
```
## Usage Guide
- Basic usage
```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv
```

- With bonus features
```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv on
```

- Help
```bash
go run . -h
```

## Input Format
**Airport Codes**  
    #LAX → IATA code  
    ##EGLL → ICAO code  
    *#LAX / *##EGLL → city name instead of airport  

**Date/Time Formats**  
    D(2022-05-09T08:07Z) → Date  
    T12(2022-05-09T08:07Z) → 12-hour format  
    T24(2022-05-09T08:07-02:00) → 24-hour format  

