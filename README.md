# savagedog

## Requirements

* protoc-gen-go
* protoc-gen-go-grpc

## Building

```
make
```

## Running

Create a Discord app and configure the bot. Add the bot to your server. Make sure your
bot has permission to send messages.

Create configuration YAML file
```yaml
hostname: daemon_host
port: daemon_port
discord:
  token: <discord bot token>
services:
  - name: myservice
    channel: mychannel
```

With this configuration, every notification received from `myservice` will be sent 
to `mychannel` of your Discord server.

Run daemon
```aiignore
/bin/savagedog serve --config /path/to/config.yaml
```

When you want to send a notification, run the app and specify a destination:
```aiignore
/bin/savagedog notify --destination daemon_host:daemon_port --from "myservice" \
--header "This is a header" \
--content "And this is a message" \
--sender "Service Name" \
--fields "key1=value1&key2=value2"
```

Some services may send multiple notifications from different sources. In that case,
you can create a config file where you specify some fields and avoid them when 
invoking the app. Example configuration:
```aiignore
destination: daemon_host:daemon_port
from: myservice
sender: Service Name
```

Send a notification:
```aiignore
/bin/savagedog notify --config /path/to/config.yaml \
--header "This is a header" \
--content "And this is a message" \
--fields "key1=value1&key2=value2"
```