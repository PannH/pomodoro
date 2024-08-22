# Pomodoro
This is a simple and minimalist pomodoro CLI written in Go.

### Usage
Run the `go run` command with `-h`, `-m`, and `-s` flags to set the pomodoro time.
```bash
go run .\main\main.go -m 1 -s 30
```
*The pomodoro will run for 1m 30s.*

To compile the app, use the `go build` command.
```bash
go build -o pomodoro.exe .\main\main.go
```

### Example
![Usage Example GIF](./assets/usage_example.gif)