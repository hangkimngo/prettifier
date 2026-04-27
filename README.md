# Intinerary Prettifier

## Project overview:

This project is a command-line tool that processes an input text file containing flight itineraries and formats it into a more readable output.

The tool performs the following transformations:

1. Replaces airport codes:
    * #IATA → full airport name
    * ##ICAO → full airport name
    * *#CODE → city name instead of airport name *(bonus feature)*
2. Formats date and time expressions:
    * Date: D(2007-04-05T12:30−02:00) → 05 Apr 2007  
    * 12 Hour time: T12(2007-04-05T12:30−02:00) → 12:30PM (-02:00)  
    * 24 Hour time: T24(2007-04-05T12:30−02:00) → 12:30 (-02:00)  
    * UTC (Z) offset: "Z" → displayed as (+00:00)  
3. Normalizes whitespace:
    * Converts all vertical whitespace (\r, \v, \f) into standard newlines (\n)
Removes excessive blank lines for cleaner output
    * Removes excessive blank lines (no more than two consecutive new-line characters).  

### Bonus features:
1. Terminal preview mode:  
Highlights dates, times, offsets, airport names, and cities using colors, italic and bold formatting
2. City name support:  
Example: *#LHR → London  
3. Support dynamic airport lookup column order  

## Examples:  
**Airport Replacement**
Input:  
```
Your flight departs from #HAJ, and your destination is ##EDDW.  
```

Results:  
```
Your flight departs from Hannover Airport, and your destination is Bremen Airport.  
```
**City Name (Bonus)**
Input:  
```
Your flight departs from *#HAJ, and your destination is *##EDDW.  
```

Results:  
```
Your flight departs from Hannover, and your destination is Bremen.  
```
**Date & Time Formatting**
Input:  
```
1. D(2022-05-09T08:07Z)
2. T12(2069-04-24T19:18-02:00)
3. T24(2032-07-17T04:08+13:00)  
4. T12(2080-05-04T14:54Z)  
```

Results:  
```
1. 09 May 2022
2. 07:18PM (-02:00)
3. 04:08 (+13:00)
4. 02:54PM (+00:00)  
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
- **Basic usage**
```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv
```

- **With Preview Mode (Bonus)**
```bash
go run . ./input.txt ./output.txt ./airport-lookup.csv on
```

- **Help**
```bash
go run . -h
```
