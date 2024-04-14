# URL Pinger

The idea is to have a list of urls that can be pinged periodically at a specified interval and the response recorded to a data store. I like to think of it like a poor man's Uptime Kuma, you provide a list of websites and the ping interval as a CSV file.  

I am using this to learn Go and to understand how to build a Go application and explore the utility of goroutines and channels.

## Usage

### Pre-requisites

Create a database with a table having the following schema -

```sql

CREATE TABLE website_monitor_log (
    id serial4 NOT NULL,
    url text NULL,
    status_code varchar(50) NULL,
    error text NULL,
    regex_matched bool NULL,
    response_time float8 NULL,
    thread_id int4 NOT NULL,
    session_id varchar(255) NOT NULL,
    "timestamp" timestamp NOT NULL,
    CONSTRAINT website_monitor_log_pkey PRIMARY KEY (id)
);

```

Create an environment variable that hold the Database Connection details -

```bash

export DB_CONN=<connection string>

```

Create a csv file with the following format -

```csv

url,check_pattern,ping_interval
http://example.com,example,5

```

You can pass in  a csv file and a session id to identify the session that the app is running.

### How to run

From the repository root run -

```bash

 go run cmd/main.go --csv <file name> --session <session name>       

```
