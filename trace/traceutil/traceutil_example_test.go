// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package traceutil_test

import (
	"log"

	"go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/traceutil"
)

func ExampleFilteringExporter() {
	jexporter, err := jaeger.NewExporter(jaeger.Options{
		Endpoint:    "http://localhost:14268",
		ServiceName: "trace-demo",
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(&traceutil.FilteringExporter{
		Exporter: jexporter,
		Matchers: []*traceutil.Matcher{
			traceutil.BySpanNamePrefixMatcher("Recv.helloworld"),
			traceutil.BySpanKindMatcher(trace.SpanKindServer),
		},
	})
}
