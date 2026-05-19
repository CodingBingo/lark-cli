// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package core

import (
	"os"
	"strings"

	"github.com/larksuite/cli/internal/envvars"
)

// LarkBrand represents the Lark platform brand.
// "feishu" targets China-mainland, "lark" targets international.
// Any other string falls back to Feishu defaults; service base URLs can be
// overridden via LARKSUITE_CLI_*_BASE_URL environment variables.
type LarkBrand string

const (
	BrandFeishu LarkBrand = "feishu"
	BrandLark   LarkBrand = "lark"
)

// ParseBrand normalizes a brand string to a LarkBrand constant.
// Unrecognized values default to BrandFeishu.
func ParseBrand(value string) LarkBrand {
	if value == "lark" {
		return BrandLark
	}
	return BrandFeishu
}

// Endpoints holds resolved endpoint URLs for different Lark services.
type Endpoints struct {
	Open     string // e.g. "https://open.feishu.cn"
	Accounts string // e.g. "https://accounts.feishu.cn"
	MCP      string // e.g. "https://mcp.feishu.cn"
	AppLink  string // e.g. "https://applink.feishu.cn"
}

// ResolveEndpoints resolves endpoint URLs based on brand.
func ResolveEndpoints(brand LarkBrand) Endpoints {
	var endpoints Endpoints
	switch brand {
	case BrandLark:
		endpoints = Endpoints{
			Open:     "https://open.larksuite.com",
			Accounts: "https://accounts.larksuite.com",
			MCP:      "https://mcp.larksuite.com",
			AppLink:  "https://applink.larksuite.com",
		}
	default:
		endpoints = Endpoints{
			Open:     "https://open.feishu.cn",
			Accounts: "https://accounts.feishu.cn",
			MCP:      "https://mcp.feishu.cn",
			AppLink:  "https://applink.feishu.cn",
		}
	}
	return applyEndpointEnvOverrides(endpoints)
}

// ResolveOpenBaseURL returns the Open API base URL for the given brand.
func ResolveOpenBaseURL(brand LarkBrand) string {
	return ResolveEndpoints(brand).Open
}

func applyEndpointEnvOverrides(endpoints Endpoints) Endpoints {
	if override := normalizeBaseURL(os.Getenv(envvars.CliOpenBaseURL)); override != "" {
		endpoints.Open = override
	}
	if override := normalizeBaseURL(os.Getenv(envvars.CliAccountsBaseURL)); override != "" {
		endpoints.Accounts = override
	}
	if override := normalizeBaseURL(os.Getenv(envvars.CliMCPBaseURL)); override != "" {
		endpoints.MCP = override
	}
	if override := normalizeBaseURL(os.Getenv(envvars.CliAppLinkBaseURL)); override != "" {
		endpoints.AppLink = override
	}
	return endpoints
}

func normalizeBaseURL(raw string) string {
	baseURL := strings.TrimSpace(raw)
	if baseURL == "" {
		return ""
	}
	if !strings.Contains(baseURL, "://") {
		baseURL = "https://" + baseURL
	}
	return strings.TrimRight(baseURL, "/")
}
