// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.0
// - protoc             v5.27.1
// source: kessel/inventory/v1beta1/policies_service.proto

package v1beta1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationPoliciesServiceCreatePolicy = "/api.kessel.inventory.v1beta1.PoliciesService/CreatePolicy"
const OperationPoliciesServiceDeletePolicy = "/api.kessel.inventory.v1beta1.PoliciesService/DeletePolicy"
const OperationPoliciesServiceUpdatePolicy = "/api.kessel.inventory.v1beta1.PoliciesService/UpdatePolicy"

type PoliciesServiceHTTPServer interface {
	CreatePolicy(context.Context, *CreatePolicyRequest) (*CreatePolicyResponse, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*DeletePolicyResponse, error)
	UpdatePolicy(context.Context, *UpdatePolicyRequest) (*UpdatePolicyResponse, error)
}

func RegisterPoliciesServiceHTTPServer(s *http.Server, srv PoliciesServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/api/inventory/v1beta1/policies", _PoliciesService_CreatePolicy0_HTTP_Handler(srv))
	r.PUT("/api/inventory/v1beta1/policies/{resource}", _PoliciesService_UpdatePolicy0_HTTP_Handler(srv))
	r.DELETE("/api/inventory/v1beta1/policies/{resource}", _PoliciesService_DeletePolicy0_HTTP_Handler(srv))
}

func _PoliciesService_CreatePolicy0_HTTP_Handler(srv PoliciesServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePolicyRequest
		if err := ctx.Bind(&in.Policy); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPoliciesServiceCreatePolicy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePolicy(ctx, req.(*CreatePolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePolicyResponse)
		return ctx.Result(200, reply)
	}
}

func _PoliciesService_UpdatePolicy0_HTTP_Handler(srv PoliciesServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdatePolicyRequest
		if err := ctx.Bind(&in.Policy); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPoliciesServiceUpdatePolicy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePolicy(ctx, req.(*UpdatePolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdatePolicyResponse)
		return ctx.Result(200, reply)
	}
}

func _PoliciesService_DeletePolicy0_HTTP_Handler(srv PoliciesServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePolicyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationPoliciesServiceDeletePolicy)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePolicy(ctx, req.(*DeletePolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeletePolicyResponse)
		return ctx.Result(200, reply)
	}
}

type PoliciesServiceHTTPClient interface {
	CreatePolicy(ctx context.Context, req *CreatePolicyRequest, opts ...http.CallOption) (rsp *CreatePolicyResponse, err error)
	DeletePolicy(ctx context.Context, req *DeletePolicyRequest, opts ...http.CallOption) (rsp *DeletePolicyResponse, err error)
	UpdatePolicy(ctx context.Context, req *UpdatePolicyRequest, opts ...http.CallOption) (rsp *UpdatePolicyResponse, err error)
}

type PoliciesServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPoliciesServiceHTTPClient(client *http.Client) PoliciesServiceHTTPClient {
	return &PoliciesServiceHTTPClientImpl{client}
}

func (c *PoliciesServiceHTTPClientImpl) CreatePolicy(ctx context.Context, in *CreatePolicyRequest, opts ...http.CallOption) (*CreatePolicyResponse, error) {
	var out CreatePolicyResponse
	pattern := "/api/inventory/v1beta1/policies"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPoliciesServiceCreatePolicy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in.Policy, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PoliciesServiceHTTPClientImpl) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...http.CallOption) (*DeletePolicyResponse, error) {
	var out DeletePolicyResponse
	pattern := "/api/inventory/v1beta1/policies/{resource}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationPoliciesServiceDeletePolicy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *PoliciesServiceHTTPClientImpl) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...http.CallOption) (*UpdatePolicyResponse, error) {
	var out UpdatePolicyResponse
	pattern := "/api/inventory/v1beta1/policies/{resource}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationPoliciesServiceUpdatePolicy))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.Policy, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
