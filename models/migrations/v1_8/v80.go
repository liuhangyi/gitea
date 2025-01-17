// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1_8 // nolint

import "xorm.io/xorm"

func AddIsLockedToIssues(x *xorm.Engine) error {
	// Issue see models/issue.go
	type Issue struct {
		ID       int64 `xorm:"pk autoincr"`
		IsLocked bool  `xorm:"NOT NULL DEFAULT false"`
	}

	return x.Sync2(new(Issue))
}
