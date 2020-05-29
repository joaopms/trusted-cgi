// Code generated by jsonrpc2. DO NOT EDIT.
package handlers

import (
	"context"
	"encoding/json"
	jsonrpc2 "github.com/reddec/jsonrpc2"
	api "github.com/reddec/trusted-cgi/api"
)

func RegisterProjectAPI(router *jsonrpc2.Router, wrap api.ProjectAPI, typeHandler interface {
	ValidateToken(ctx context.Context, value *api.Token) error
}) []string {
	router.RegisterFunc("ProjectAPI.Config", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.Config(ctx, args.Arg0)
	})

	router.RegisterFunc("ProjectAPI.SetUser", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
			Arg1 string     `json:"user"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.SetUser(ctx, args.Arg0, args.Arg1)
	})

	router.RegisterFunc("ProjectAPI.SetEnvironment", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token      `json:"token"`
			Arg1 api.Environment `json:"env"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.SetEnvironment(ctx, args.Arg0, args.Arg1)
	})

	router.RegisterFunc("ProjectAPI.AllTemplates", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.AllTemplates(ctx, args.Arg0)
	})

	router.RegisterFunc("ProjectAPI.List", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.List(ctx, args.Arg0)
	})

	router.RegisterFunc("ProjectAPI.Templates", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.Templates(ctx, args.Arg0)
	})

	router.RegisterFunc("ProjectAPI.Stats", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
			Arg1 int        `json:"limit"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.Stats(ctx, args.Arg0, args.Arg1)
	})

	router.RegisterFunc("ProjectAPI.Create", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.Create(ctx, args.Arg0)
	})

	router.RegisterFunc("ProjectAPI.CreateFromTemplate", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
			Arg1 string     `json:"templateName"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.CreateFromTemplate(ctx, args.Arg0, args.Arg1)
	})

	router.RegisterFunc("ProjectAPI.CreateFromGit", func(ctx context.Context, params json.RawMessage, positional bool) (interface{}, error) {
		var args struct {
			Arg0 *api.Token `json:"token"`
			Arg1 string     `json:"repo"`
		}
		var err error
		if positional {
			err = jsonrpc2.UnmarshalArray(params, &args.Arg0, &args.Arg1)
		} else {
			err = json.Unmarshal(params, &args)
		}
		if err != nil {
			return nil, err
		}
		err = typeHandler.ValidateToken(ctx, args.Arg0)
		if err != nil {
			return nil, err
		}
		return wrap.CreateFromGit(ctx, args.Arg0, args.Arg1)
	})

	return []string{"ProjectAPI.Config", "ProjectAPI.SetUser", "ProjectAPI.SetEnvironment", "ProjectAPI.AllTemplates", "ProjectAPI.List", "ProjectAPI.Templates", "ProjectAPI.Stats", "ProjectAPI.Create", "ProjectAPI.CreateFromTemplate", "ProjectAPI.CreateFromGit"}
}