# opentelemetry-sampler-go

Honeycomb Samplers for use with the OpenTelemetry Go SDK

## Samplers

### Deterministic Sampler

This is a port of the deterministic sampler included in our [Go Beeline](https://github.com/honeycombio/beelinee-go). To use it, just instantiate it with a sample rate:

```golang
import (
	"github.com/honeycombio/opentelemetry-sampler-go"
	"go.opentelemetry.io/otel/sdk/trace"
)

config := trace.Config{
	DefaultSampler: honeycomb.DeterministicSampler(5),
}
```
