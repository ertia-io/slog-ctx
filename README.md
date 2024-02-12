# slogx

Attach a slog.Logger to a context

```go
logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

// Attach
ctx := slogx.WithContext(context.Background(), logger)

// Retrieve
log := slogx.Ctx(ctx)

log.Info("Hello world")
```
