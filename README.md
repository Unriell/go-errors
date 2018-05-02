Go-Error lib
------------

1. test and compile

```Terminal
make all
```
*check Makefile in order to know more*

2. Install

```Terminal
govendor fetch github.com/bq/go-errors/errors
```

3. how to use it
*unitTest in order to know more*

Firstly you should init your errorHandler (that is linked to a sentry account)

```go
var enviroment = "QA"
var sentryDns = "https://1eecf0d5795b4f9bbc5f4210c1513f4c:35ae8d2ab04144d8a01127612ed810d9@sentry.bq.com/132"
var version = "0.0.1"

errorHandler := NewErrorHandler(enviroment, sentryDns, version)
```

Secondly you can instanciate a typed error
example
```go
binError := &BindError{Err: errors.New("bind error"), Message: "my custom bind msg", RequestID: "requestId"}
```

Finally invoke your error handler.error method in order to get an errorDto (at this point sentry will be notified). 



```go
typedErrorDto := errHandler.Error(binError, nil)
```

*Note that the second parameter in this case is null but you can add custom sentry tags. errorHandler.error will accept any kind of errors not only typedErrors.*
