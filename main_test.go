package main

import "testing"

func TestValidateParams(t *testing.T) {
	type args struct {
		params []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Valid params",
			args: args{
				params: []string{"firebase.json", "users"},
			},
			wantErr: false,
		},
		{
			name: "Invalid params",
			args: args{
				params: []string{"firebase.json"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateParams(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("validateParams() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
