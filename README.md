# opentelemetry-samplers-go

**NOTE**: This is experimental and is subject to change a _lot_ or go away entirely. Use with caution.

Honeycomb Samplers for use with the OpenTelemetry Go SDK

## Samplers

### Deterministic Sampler

This is a port of the deterministic sampler included in our [Go Beeline](https://github.com/honeycombio/beelinee-go). To use it, just instantiate it with a sample rate:

```golang
import (
	"github.com/honeycombio/opentelemetry-samplers-go/honeycombsamplers"
	"go.opentelemetry.io/otel/sdk/trace"
)

config := trace.Config{
	DefaultSampler: honeycombsamplers.DeterministicSampler(5),
}
```
