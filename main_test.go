package main

import (
	"context"
	"github.com/Domains18/golang-grpc/invoicer"
	"reflect"
	"testing"
)

func TestTheInvoicerServer_Create(t *testing.T) {
	type fields struct {
		UnimplementedInvoicerServer invoicer.UnimplementedInvoicerServer
	}
	type args struct {
		ctx context.Context
		req *invoicer.CreateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *invoicer.CreateResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TheInvoicerServer{
				UnimplementedInvoicerServer: tt.fields.UnimplementedInvoicerServer,
			}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
