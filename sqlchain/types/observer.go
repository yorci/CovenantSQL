/*
 * Copyright 2018 The CovenantSQL Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package types

const (
	// ObserverService is the service name for observer to receive.
	ObserverService = "OBS"
	// ReplicateFromBeginning is the replication offset observes from genesis block.
	ReplicateFromBeginning = int32(0)
	// ReplicateFromNewest is the replication offset observes from block head of current node.
	ReplicateFromNewest = int32(-1)
)
