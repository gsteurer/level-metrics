# level-metrics

Author: Joseph A. Steurer

Reads files of a similar type from a directory and summarizes the contents. 

Input must be sorted by timestamp in ascending order. 

## Packages
 * cmd: Utilities that transform files containing level metrics into a summary.
 * config: Parses application configuration from user input.
 * filter: Decouples processing filters from parsed configuration.
 * ingest: Provides mechansims for reading level metrics from a data source.
 * input: Shared consts for defining input format
 * output: Shared consts for defining the output format
 * processor: Aggregates level metrics into a summary.
 * model/record: The data model for an individual level metric datapoint.
 * model/summary: The data model for a level summary.
 
## `cmd/metrics-cli`
Build with `go build ./cmd/metrics-cli/`

A CLI utility for transforming a directory of files containing level measurements and reporting a summary.

Usage:
```go run . -d="../../data/json" -type=json -startTime="2022-04-01T12:00:00.00Z" -endTime="2022-04-29T12:00:00.00Z" -outputFileName="../../out.json"```  


### Dependencies

`go version go1.18`

## features checklist
### Flags
* Specify the directory containing metric files with “-d” or “--directory” option.
* Specify the metric file type with “-t” or “--type” option. 
* Specify the start time of the time range using “--startTime” option.
* Specify the end time of the time range using “--endTime” option.  
* Specify the output file type using the “--outputFileType” option. 
* Specify the output file name using the “--outputFileName” option.  
* The app will create a new result file each time it runs and overwrite if the file already exists.

### Behavior
* The application parses files between the time range. 
* The application calculates the summary of the content inside the files.
* The application does not store data as it is read or read an entire file into memory.
* The result is displayed to the console in a human readable format. 
* The result is also written to a file. 

### Input Constraints
* Each file contains metrics usage for 1 day and the row is sorted by timestamp ascending where the oldest is at the top.
* The timestamp of data in input files (json/csv) in UTC format.
* Implement a mechanism to read only relevant files to make the application efficient. 
* The input file timestamps and the time range parameter are in the same (current) year starting from January 1st.
