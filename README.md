# Protorec

Protorec is a configurable endpoint for testing alarm devices, it currently handles "SIA-CID" and "CSV-IP" messages over both UDP and TCP. The server is configured using a JSON file.

## Usage

```bash
$ ./protorec -config ./config.json
```

The config file is structured as such.

``` json
{
  "protocol": "tcp",
  "host": "127.0.0.1",
  "port": 8080,
  "message_handler": "SIA-CID",
  "logfile": "./test.log"
}
```

Simply change the `"protocol"` and `"message_handler"`, to determine how to run the server and parse incoming messages. All server output is written to STDOUT as well as the logfile specified in the `"logfile"` setting.

