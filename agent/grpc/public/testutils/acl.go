package testutils

import (
	"testing"

	"github.com/hashicorp/consul/acl"
	"github.com/stretchr/testify/require"
)

func TestAuthorizer(t *testing.T) acl.Authorizer {
	t.Helper()

	policy, err := acl.NewPolicyFromSource(`
		service "foo" {
			policy = "write"
		}
	`, acl.SyntaxCurrent, nil, nil)
	require.NoError(t, err)

	authz, err := acl.NewPolicyAuthorizerWithDefaults(acl.DenyAll(), []*acl.Policy{policy}, nil)
	require.NoError(t, err)

	return authz
}