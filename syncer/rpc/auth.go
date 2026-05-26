package rpc

import (
	"context"
	"fmt"
)

var errWrongAccountNorRole = fmt.Errorf("account should be %s, and roles should contain %s", RbacAllowedAccountName, RbacAllowedRoleName)

func auth(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
