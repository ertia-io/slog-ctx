# slog-ctx

Attach a slog.Logger to a context

```go
logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

// Attach
ctx := slogctx.WithContext(context.Background(), logger)

// Retrieve
log := slogctx.Ctx(ctx)

log.Info("Hello world")
```
