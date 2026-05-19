// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package envvars

const (
	CliAppID             = "LARKSUITE_CLI_APP_ID"
	CliAppSecret         = "LARKSUITE_CLI_APP_SECRET"
	CliBrand             = "LARKSUITE_CLI_BRAND"
	CliOpenBaseURL       = "LARKSUITE_CLI_OPEN_BASE_URL"
	CliAccountsBaseURL   = "LARKSUITE_CLI_ACCOUNTS_BASE_URL"
	CliMCPBaseURL        = "LARKSUITE_CLI_MCP_BASE_URL"
	CliAppLinkBaseURL    = "LARKSUITE_CLI_APPLINK_BASE_URL"
	CliUserAccessToken   = "LARKSUITE_CLI_USER_ACCESS_TOKEN"
	CliTenantAccessToken = "LARKSUITE_CLI_TENANT_ACCESS_TOKEN"
	CliDefaultAs         = "LARKSUITE_CLI_DEFAULT_AS"
	CliStrictMode        = "LARKSUITE_CLI_STRICT_MODE"

	// Sidecar proxy (auth proxy mode)
	CliAuthProxy = "LARKSUITE_CLI_AUTH_PROXY" // sidecar HTTP address, e.g. "http://127.0.0.1:16384"
	CliProxyKey  = "LARKSUITE_CLI_PROXY_KEY"  // HMAC signing key shared with sidecar

	// Content safety scanning mode
	CliContentSafetyMode = "LARKSUITE_CLI_CONTENT_SAFETY_MODE"
)
