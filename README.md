# chatbot-api

## Commands

### advice
Get a piece of advice, either random or by ID.

**Get random advice**
```http
GET /commands/advice
```

**Get advice by ID**
```http
GET /commands/advice/{ID}
```

### chucknorris
Get a random satirical factoid about martial artist and actor Chuck Norris.

```http
GET /commands/jokes/chucknorris
```

### ctof
Convert a temperature from Celsius to Fahrenheit.

```http
GET /commands/temperature/ctof/{temperature}
```

### ctok
Convert a temperature from Celsius to Kelvin.

```http
GET /commands/temperature/ctok/{temperature}
```

### ftoc
Convert a temperature from Fahrenheit to Celsius.

```http
GET /commands/temperature/ftoc/{temperature}
```

### ftok
Convert a temperature from Fahrenheit to Kelvin.

```http
GET /commands/temperature/ftok/{temperature}
```

### ktoc
Convert a temperature from Kelvin to Celsius.

```http
GET /commands/temperature/ktoc/{temperature}
```

### ktof
Convert a temperature from Kelvin to Fahrenheit.

```http
GET /commands/temperature/ktof/{temperature}
```

### current time, 12-hour format
Get the current time in a 12-hour format for a certain timezone.

```http
GET /commands/time/current/12/{timezone}
```

### current time, 24-hour format
Get the current time in a 24-hour format for a certain timezone.

```http
GET /commands/time/current/24/{timezone}
```

## License
[MIT License](./LICENSE)
