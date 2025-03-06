# Logger

This simple Golang package can be used for common logging functionality.

# Options

```
logger.SetAppName("todo-app")     # optional; default "MyApp"
logger.SetDebug(true)             # optional; default false
```

# Usage examples
```
logger.Print("Simple log message withouth error", nil)
logger.Fatal("Fatal log message with error and exits the program", err)
```
Will be printed like this:
```
2025/12/01 11:17:35 todo-app | LOG: Simple log message withouth error
2025/12/01 11:17:35 todo-app | FATAL: Fatal log message with error
2025/12/01 11:17:35 todo-app | ERROR: <insert value of err> 
```

# Return functions
```
logger.GetAppName()  # will return the current used app name as a string.
logger.GetDebug()    # will return the current debug setting as a bool.
