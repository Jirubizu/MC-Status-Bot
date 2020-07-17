# MC-Status-Bot
Get an update if one of mojangs services change state

## How to use
### Things to note
Before building please provide your own bot token in `main.go line 13 ` additionally, inside `bot.go` on line 122 replace the `CHANNEL ID` with the channel ID of where you want the bot to send the update messages to.

### Building
Inside your terminal simply run `go build` once the bot has finished building you can run it either by double clicking the exe or running it through the terminal simply being in the same directory and writing `minecraftStatusChecker.exe`. I