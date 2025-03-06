# Logger

This simple Golang package can be used for common logging functionality.
Only when debug is on, non fatal errors will be shown.

# Install
```
go get -u github.com/scornet256/go-logger
```

# Basic usage
```
logger.Print("Simple log message withouth error", nil)
logger.Fatal("Fatal log message with error and exits the program", err)
```
Will output like this:
```
2025/12/01 11:17:35 MyApp | LOG: Simple log message withouth error
2025/12/01 11:17:35 MyApp | FATAL: Fatal log message with error
2025/12/01 11:17:35 MyApp | ERROR: <insert value of err> 
```

# Options
```
logger.SetAppName("todo-app")     # optional; default "MyApp"
logger.SetDebug(true)             # optional; default false
```

# Return functions
```
logger.GetAppName()  # will return the current used app name as a string.
logger.GetDebug()    # will return the current debug setting as a bool.
