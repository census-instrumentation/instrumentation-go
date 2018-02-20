// Copyright 2017, OpenCensus Authors
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
//

package ocgrpc

import (
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// The following variables are measures are recorded by ClientHandler:

var (
	ClientErrorCount = stats.Measure{
		"grpc.io/client/error_count",
		"RPC Errors",
		stats.UnitNone,
	}
	ClientRequestBytes = stats.Measure{
		Name:        "grpc.io/client/request_bytes",
		Description: "Request bytes",
		Unit:        stats.UnitBytes,
	}
	ClientResponseBytes = stats.Measure{
		"grpc.io/client/response_bytes",
		"Response bytes",
		stats.UnitBytes,
	}
	ClientStartedCount = stats.Measure{
		Name:        "grpc.io/client/started_count",
		Description: "Number of client RPCs (streams) started",
		Unit:        stats.UnitNone,
	}
	ClientFinishedCount    = stats.Measure{"grpc.io/client/finished_count", "Number of client RPCs (streams) finished", stats.UnitNone}
	ClientRequestCount     = stats.Measure{"grpc.io/client/request_count", "Number of client RPC request messages", stats.UnitNone}
	ClientResponseCount    = stats.Measure{"grpc.io/client/response_count", "Number of client RPC response messages", stats.UnitNone}
	ClientRoundTripLatency = stats.Measure{"grpc.io/client/roundtrip_latency", "RPC roundtrip latency in msecs", stats.UnitMilliseconds}
)

var (
	recordClientError, _   = stats.Int64(&ClientErrorCount)
	recordRequestBytes, _  = stats.Int64(&ClientRequestBytes)
	recordResponseBytes, _ = stats.Int64(&ClientResponseBytes)
	recordStarted, _       = stats.Int64(&ClientStartedCount)
	recordFinished, _      = stats.Int64(&ClientFinishedCount)
	recordRequestCount, _  = stats.Int64(&ClientRequestCount)
	recordResponseCount, _ = stats.Int64(&ClientResponseCount)
	recordLatency, _       = stats.Float64(&ClientRoundTripLatency)
)

// Predefined views may be subscribed to collect data for the above measures.
// As always, you may also define your own custom views over measures collected by this
// package. These are declared as a convenience only; none are subscribed by
// default.
var (
	ClientErrorCountView, _ = view.New(
		"grpc.io/client/error_count",
		"RPC Errors",
		[]tag.Key{KeyStatus, KeyMethod},
		ClientErrorCount.Name,
		view.MeanAggregation{})

	ClientRoundTripLatencyView, _ = view.New(
		"grpc.io/client/roundtrip_latency",
		"Latency in msecs",
		[]tag.Key{KeyMethod},
		ClientRoundTripLatency,
		DefaultMillisecondsDistribution)

	ClientRequestBytesView, _ = view.New(
		"grpc.io/client/request_bytes",
		"Request bytes",
		[]tag.Key{KeyMethod},
		ClientRequestBytes,
		DefaultBytesDistribution)

	ClientResponseBytesView, _ = view.New(
		"grpc.io/client/response_bytes",
		"Response bytes",
		[]tag.Key{KeyMethod},
		ClientResponseBytes,
		DefaultBytesDistribution)

	ClientRequestCountView, _ = view.New(
		"grpc.io/client/request_count",
		"Count of request messages per client RPC",
		[]tag.Key{KeyMethod},
		ClientRequestCount,
		DefaultMessageCountDistribution)

	ClientResponseCountView, _ = view.New(
		"grpc.io/client/response_count",
		"Count of response messages per client RPC",
		[]tag.Key{KeyMethod},
		ClientResponseCount,
		DefaultMessageCountDistribution)
)

// All the default client views provided by this package:
var (
	DefaultClientViews = []*view.View{
		ClientErrorCountView,
		ClientRoundTripLatencyView,
		ClientRequestBytesView,
		ClientResponseBytesView,
		ClientRequestCountView,
		ClientResponseCountView,
	}
)

// TODO(jbd): Add roundtrip_latency, uncompressed_request_bytes, uncompressed_response_bytes, request_count, response_count.
// TODO(acetechnologist): This is temporary and will need to be replaced by a
// mechanism to load these defaults from a common repository/config shared by
// all supported languages. Likely a serialized protobuf of these defaults.
