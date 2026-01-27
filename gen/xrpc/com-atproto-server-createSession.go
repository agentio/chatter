// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.server.createSession

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const ComATProtoServerCreateSession_Description = "Create an authentication session."

type ComATProtoServerCreateSession_Input struct {
	AllowTakendown  *bool   `json:"allowTakendown,omitempty"`
	AuthFactorToken *string `json:"authFactorToken,omitempty"`
	Identifier      string  `json:"identifier"`
	Password        string  `json:"password"`
}

type ComATProtoServerCreateSession_Output struct {
	AccessJwt       string  `json:"accessJwt"`
	Active          *bool   `json:"active,omitempty"`
	Did             string  `json:"did"`
	DidDoc          *any    `json:"didDoc,omitempty"`
	Email           *string `json:"email,omitempty"`
	EmailAuthFactor *bool   `json:"emailAuthFactor,omitempty"`
	EmailConfirmed  *bool   `json:"emailConfirmed,omitempty"`
	Handle          string  `json:"handle"`
	RefreshJwt      string  `json:"refreshJwt"`
	Status          *string `json:"status,omitempty"`
}

// Create an authentication session.
func ComATProtoServerCreateSession(ctx context.Context, c slink.Client, input *ComATProtoServerCreateSession_Input) (*ComATProtoServerCreateSession_Output, error) {
	var output ComATProtoServerCreateSession_Output
	if err := c.Do(ctx, slink.Procedure, "application/json", "com.atproto.server.createSession", nil, input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
