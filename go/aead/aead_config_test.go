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

package aead_test

import (
	"github.com/google/tink/go/aead/aead"
	"github.com/google/tink/go/tink/tink"
	"testing"
)

func TestConfigInstance(t *testing.T) {
	c := aead.Config()
	if c == nil {
		t.Errorf("Config() returns nil")
	}
}

func TestConfigRegistration(t *testing.T) {
	success, err := aead.Config().RegisterStandardKeyTypes()
	if !success || err != nil {
		t.Errorf("cannot register standard key types")
	}
	// check for AES-GCM key manager
	keyManager, err := tink.Registry().GetKeyManager(aead.AES_GCM_TYPE_URL)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	var _ = keyManager.(*aead.AesGcmKeyManager)
}
