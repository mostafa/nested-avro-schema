# nested-avro-schema

If you have a nested AVRO schema and you want to test it against your data, this small tool is for you. This tool will help you to find discrepancies and errors in your data, so that you can fix them before you run [xk6-kafka](https://github.com/mostafa/xk6-kafka) tests.

## How to use

To use this, you must have Go installed on your machine.

1. Clone this repository or [download the code](https://github.com/mostafa/nested-avro-schema/archive/refs/heads/main.zip).
2. Run it using `go run main.go <schema-file> <data-file>` command. An example nested AVRO schema and data file is provided in the `example` directory.
