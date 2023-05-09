----
Install
------------

For the Telegram bot to work correctly, you must:
1. Install Telegram API.
```
go get -u github.com/go-telegram-bot-api/telegram-bot-api
```
2. Enter the bot's token to the `token` variable.
3. Compile the code.

```
go build -o my_bot main.go && ./my_bot
```

4. Enter the `/start` command for the bot.

Documentation
-------------
What can the bot to do?
- Display a greeting after applying the `/start` command
- Display and operate the 4 buttons of the first layer.
- Display and operate the 2 buttons of the second layer.
- Know how to return from the second layer to the first layer.
- Completes the job of using the two buttons of the second layer.
- Logs all actions.
