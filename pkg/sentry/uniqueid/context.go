// Copyright 2018 Google Inc.
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

// Package uniqueid defines context.Context keys for obtaining system-wide
// unique identifiers.
package uniqueid

import (
	"gvisor.googlesource.com/gvisor/pkg/sentry/context"
)

// contextID is the kernel package's type for context.Context.Value keys.
type contextID int

const (
	// CtxGlobalUniqueID is a Context.Value key for a system-wide
	// unique identifier.
	CtxGlobalUniqueID contextID = iota

	// CtxInotifyCookie is a Context.Value key for a unique inotify
	// event cookie.
	CtxInotifyCookie
)

// GlobalFromContext returns a system-wide unique identifier from ctx.
func GlobalFromContext(ctx context.Context) uint64 {
	return ctx.Value(CtxGlobalUniqueID).(uint64)
}

// InotifyCookie generates a unique inotify event cookie from ctx.
func InotifyCookie(ctx context.Context) uint32 {
	return ctx.Value(CtxInotifyCookie).(uint32)
}
