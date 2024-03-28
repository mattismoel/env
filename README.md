# What is this?

A simple yet powerful library for handling all things environment variables.
It makes retrieving and setting environment variables a breeze, with robust
and sane error handlning.

# How does it work?

The getters follow a very simple syntax, where a desired key and a fallback
value is given. If the provided key is not to be found, the fallback value
will be returned. `env` provides the following getter functions:

```go
envStr := env.Str("ENV_KEY", "fallback_value")
envInt := env.Int("ENV_KEY", 10)
envFloat32 := env.Float32("ENV_KEY", 10)
envBool := env.Bool("ENV_KEY", false)
```

The library can also set environment variables with a given key and value.
If the input is invalid a fitting error is returned:

```go
err := env.SetStr("ENV_KEY", "env_value")
if err != nil {
    switch {
    case errors.Is(err, env.ErrInvalidKey):
        // Key is invalid
    case errors.Is(Err, env.ErrInvalidValue):
        // Value is invalid
    }
}
```

This syntax is reused in all of the other setter functions:
```go
env.SetStr("ENV_KEY", "env_value")      // ENV_KEY="env_value".
env.SetInt("ENV_KEY", 100)              // ENV_KEY="100".
env.SetFloat32("ENV_KEY", 15.25, 4)     // ENV_KEY="15.2500".
env.SetBool("ENV_KEY", false)           // ENV_KEY="false".
```
> Take note of the extra parameter `SetFloat32()`, where the last `perc`
parameter specifies the amount of digits *after* the decimal point.
