// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

package tink

import (
	"encoding/binary"
	"fmt"
	tinkpb "github.com/google/tink/proto/tink_go_proto"
)

/**
 * Constants and convenience methods that deal with crypto format.
 */
const (
	// Prefix size of Tink and Legacy key types.
	NON_RAW_PREFIX_SIZE = 5

	// Legacy or Crunchy prefix starts with \x00 and followed by a 4-byte key id.
	LEGACY_PREFIX_SIZE = NON_RAW_PREFIX_SIZE
	LEGACY_START_BYTE  = byte(0)

	// Tink prefix starts with \x01 and followed by a 4-byte key id.
	TINK_PREFIX_SIZE = NON_RAW_PREFIX_SIZE
	TINK_START_BYTE  = byte(1)

	// Raw prefix is empty.
	RAW_PREFIX_SIZE = 0
	RAW_PREFIX      = ""
)

/**
 * Generates the prefix of all cryptographic outputs (ciphertexts,
 * signatures, MACs, ...)  produced by the specified {@code key}.
 * The prefix can be either empty (for RAW-type prefix), or consists
 * of a 1-byte indicator of the type of the prefix, followed by 4
 * bytes of {@code key.KeyId} in Big Endian encoding.
 *
 * @throws error if the prefix type of {@code key} is unknown.
 * @return a prefix.
 */
func GetOutputPrefix(key *tinkpb.Keyset_Key) (string, error) {
	switch key.OutputPrefixType {
	case tinkpb.OutputPrefixType_LEGACY, tinkpb.OutputPrefixType_CRUNCHY:
		return createOutputPrefix(LEGACY_PREFIX_SIZE, LEGACY_START_BYTE, key.KeyId), nil
	case tinkpb.OutputPrefixType_TINK:
		return createOutputPrefix(TINK_PREFIX_SIZE, TINK_START_BYTE, key.KeyId), nil
	case tinkpb.OutputPrefixType_RAW:
		return RAW_PREFIX, nil
	default:
		return "", fmt.Errorf("crypto_format: unknown output prefix type")
	}
}

/**
 * Creates an output prefix. It consists of a 1-byte indicator of the type
 * of the prefix, followed by 4 bytes of {@code keyId} in Big Endian encoding.
 */
func createOutputPrefix(size int, startByte byte, keyId uint32) string {
	prefix := make([]byte, size)
	prefix[0] = startByte
	binary.BigEndian.PutUint32(prefix[1:], keyId)
	return string(prefix)
}
