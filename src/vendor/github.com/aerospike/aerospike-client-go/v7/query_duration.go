/*
 * Copyright 2014-2022 Aerospike, Inc.
 *
 * Portions may be licensed to Aerospike, Inc. under one or more contributor
 * license agreements.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy of
 * the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations under
 * the License.
 */

package aerospike

import kvs "github.com/aerospike/aerospike-client-go/v7/proto/kvs"

// QueryDuration defines the expected query duration. The server treats the query in different ways depending on the expected duration.
// This enum is ignored for aggregation queries, background queries and server versions < 6.0.
type QueryDuration int

const (
	// LONG specifies that the query is expected to return more than 100 records per node. The server optimizes for a large record set in
	// the following ways:
	//
	// Allow query to be run in multiple threads using the server's query threading configuration.
	// Do not relax read consistency for AP namespaces.
	// Add the query to the server's query monitor.
	// Do not add the overall latency to the server's latency histogram.
	// Do not allow server timeouts.
	LONG = iota

	// Short specifies that the query is expected to return less than 100 records per node. The server optimizes for a small record set in
	// the following ways:
	// Always run the query in one thread and ignore the server's query threading configuration.
	// Allow query to be inlined directly on the server's service thread.
	// Relax read consistency for AP namespaces.
	// Do not add the query to the server's query monitor.
	// Add the overall latency to the server's latency histogram.
	// Allow server timeouts. The default server timeout for a short query is 1 second.
	SHORT

	// LONG_RELAX_AP will treat query as a LONG query, but relax read consistency for AP namespaces.
	// This value is treated exactly like LONG for server versions < 7.1.
	LONG_RELAX_AP
)

func (qd QueryDuration) grpc() kvs.QueryDuration {
	switch qd {
	case LONG:
		return kvs.QueryDuration(kvs.QueryDuration_LONG)
	case SHORT:
		return kvs.QueryDuration(kvs.QueryDuration_SHORT)
	case LONG_RELAX_AP:
		return kvs.QueryDuration(kvs.QueryDuration_LONG_RELAX_AP)
	}
	panic(unreachable)
}
