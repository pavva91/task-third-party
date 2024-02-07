package dto

import (
	"reflect"
	"testing"

	"github.com/pavva91/task-third-party/models"
	"gorm.io/datatypes"
)

func TestCreateTaskRequest_Validate(t *testing.T) {
	type fields struct {
		Method  string
		URL     string
		Headers map[string]interface{}
	}
	tests := map[string]struct {
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		"invalid header": {
			fields{
				Headers: datatypes.JSONMap(map[string]interface{}{
					"Authentication": 2,
				}),
			},
			true,
		},
		"invalid url": {
			fields{
				URL: "",
			},
			true,
		},
		"invalid http method": {
			fields{

				URL:    "https://example.com",
				Method: "invalid http method",
			},
			true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r := &CreateTaskRequest{
				Method:  tt.fields.Method,
				URL:     tt.fields.URL,
				Headers: tt.fields.Headers,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("CreateTaskRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateHttpHeaders(t *testing.T) {
	type args struct {
		r *CreateTaskRequest
	}
	tests := map[string]struct {
		args    args
		wantErr bool
	}{
		"not string value": {
			args{
				&CreateTaskRequest{
					Headers: datatypes.JSONMap(map[string]interface{}{
						"Authentication": 2,
					}),
				},
			},
			true,
		},
		"not string value 2": {
			args{
				&CreateTaskRequest{
					Headers: datatypes.JSONMap(map[string]interface{}{
						"Authentication": "2",
						"ciao":           2,
					}),
				},
			},
			true,
		},
		"empty header": {
			args{
				&CreateTaskRequest{
					Headers: datatypes.JSONMap(make(map[string]interface{})),
				},
			},
			false,
		},
		"string value": {
			args{
				&CreateTaskRequest{
					Headers: datatypes.JSONMap(map[string]interface{}{
						"Authentication": "2",
					}),
				},
			},
			false,
		},
		"string value 2": {
			args{
				&CreateTaskRequest{
					Headers: datatypes.JSONMap(map[string]interface{}{
						"Authentication": "Basic    bG9naW46cGFzc3dvcmQ=",
					}),
				},
			},
			false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if err := validateHttpHeaders(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("validateHttpHeaders() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateHttpMethod(t *testing.T) {
	type args struct {
		r *CreateTaskRequest
	}
	tests := map[string]struct {
		args    args
		wantErr bool
	}{
		"empty method": {
			args{
				&CreateTaskRequest{
					Method: "",
				},
			},
			true,
		},
		"wrong method": {
			args{
				&CreateTaskRequest{
					Method: "wrong",
				},
			},
			true,
		},
		"good method lowercase": {
			args{
				&CreateTaskRequest{
					Method: "get",
				},
			},
			true,
		},
		"get method": {
			args{
				&CreateTaskRequest{
					Method: "GET",
				},
			},
			false,
		},
		"post method": {
			args{
				&CreateTaskRequest{
					Method: "POST",
				},
			},
			false,
		},
		"put method": {
			args{
				&CreateTaskRequest{
					Method: "PUT",
				},
			},
			false,
		},
		"patch method": {
			args{
				&CreateTaskRequest{
					Method: "PATCH",
				},
			},
			false,
		},
		"delete method": {
			args{
				&CreateTaskRequest{
					Method: "DELETE",
				},
			},
			false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if err := validateHttpMethod(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("validateHttpMethod() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateURL(t *testing.T) {
	type args struct {
		r *CreateTaskRequest
	}
	tests := map[string]struct {
		args    args
		wantErr bool
	}{
		"empty url": {
			args{
				&CreateTaskRequest{
					URL: "",
				},
			},
			true,
		},
		"relative path": {
			args{
				&CreateTaskRequest{
					URL: "/foo/bar",
				},
			},
			true,
		},
		"empty scheme": {
			args{
				&CreateTaskRequest{
					URL: "example.org",
				},
			},
			true,
		},
		"empty host": {
			args{
				&CreateTaskRequest{
					URL: "http://",
				},
			},
			true,
		},
		"wrong starting char host: dot": {
			args{
				&CreateTaskRequest{
					URL: "http://.com",
				},
			},
			true,
		},
		"wrong starting char host: slash": {
			args{
				&CreateTaskRequest{
					URL: "http:///com",
				},
			},
			true,
		},
		"wrong starting char host: column": {
			args{
				&CreateTaskRequest{
					URL: "http://:com",
				},
			},
			true,
		},
		"not valid": {
			args{
				&CreateTaskRequest{
					URL: "http:/example.com",
				},
			},
			true,
		},
		"protocol outside http/https": {
			args{
				&CreateTaskRequest{
					URL: "random://example.com",
				},
			},
			true,
		},
		"valid http": {
			args{
				&CreateTaskRequest{
					URL: "http://example.com",
				},
			},
			false,
		},
		"valid https": {
			args{
				&CreateTaskRequest{
					URL: "https://example.com",
				},
			},
			false,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			if err := validateURL(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("validateURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTaskRequest_ToModel(t *testing.T) {
	type fields struct {
		Method  string
		URL     string
		Headers map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.Task
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dto := &CreateTaskRequest{
				Method:  tt.fields.Method,
				URL:     tt.fields.URL,
				Headers: tt.fields.Headers,
			}
			if got := dto.ToModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTaskRequest.ToModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
