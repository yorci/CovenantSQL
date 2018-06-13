/*
 * Copyright 2018 The ThunderDB Authors.
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

package proto

import (
	"reflect"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/thunderdb/ThunderDB/types"
)

func init() {
	Copiers[reflect.TypeOf(time.Time{})] = timeCopier
	TypeMap[reflect.TypeOf(time.Time{})] = reflect.TypeOf(time.Time{})

	Copiers[reflect.TypeOf(btcec.PublicKey{})] = publicKeyMarshal
	TypeMap[reflect.TypeOf(btcec.PublicKey{})] = reflect.TypeOf(types.PublicKey{})
}

func timeCopier(v interface{}) (interface{}, error) {
	// Just... copy it.
	return v.(time.Time), nil
}

func publicKeyMarshal(v interface{}) (interface{}, error) {
	pubkey, _ := v.(btcec.PublicKey)
	pbpubkey := types.PublicKey{PublicKey: pubkey.SerializeCompressed()}
	return pbpubkey, nil
}
