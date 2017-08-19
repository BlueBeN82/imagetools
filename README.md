# it (imagetools)

A tiny helper library, written in golang, that serves some utilities that can
be used to build more robust docker containers

### it requires

Validates environment variables and terminates script if they dont match your
expectations. By default it expects the given env variable to be a not empty
string.

```
it requires ES_HOST
it requires ES_PORT --validate="int" --validate-min="1" --validate-max="65536" --example=9200
it requires ES_PROTOCOL --validate="regex" --validate-pattern="^http(s)?$" --example="http"
```

##### Arguments

1. Name of environment variable to check

##### Flags

--validate      (string)    int, float, bool, string (default), regex
--min           (int|float) The minimum accepted value
--max           (int|float) The maximum accepted value
--pattern       (string)    The regular expression used to validate the value
--example       (string)    An example value to be used in error message

### it waits-for

Waits for an external service to become available.

```
it waits-for http://elasticsearch.domain.my:9200 --status-code=200
it waits-for http://elasticsearch.domain.my:9200 --match="(some|any)string"
it waits-for --protocol="tcp" --hostname="1.2.3.4" --port=3306
```

##### Flags

--protocol        (string)  The protocol to use for check
--hostname        (string)  The DNS hostname or ip address to check
--port            (int)     The port to check
--status-code     (int)     The expected http status code
--match           (regexp)  A regular expression expected in response header or body
--interval        (int)     Seconds between the checks
--timeout         (int)     Maximum wait time in seconds before giving up

### it logs

Super easy logging

```
it logs "A simple info message"
it logs "A debug message" --debug
it logs "A warning message" --warning
it logs "A success message" --success
it logs "A error message" --error
it logs "A message to stderr" --stderr
```

##### Arguments

1. The Message to be logged

##### Flags
--debug           (bool)    Logs the message with debug severity
--info            (bool)    Logs the message with info severity (default)
--warning         (bool)    Logs the message with warning severity
--error           (bool)    Logs the message with error severity
--success         (bool)    Logs the message with success severity
--stdout          (bool)    Logs the message to STDOUT
--stderr          (bool)    Logs the message to STDERR
