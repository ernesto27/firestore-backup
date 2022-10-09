package backup

import (
	"fmt"
	"testing"
)

func TestGetJSONFromData(t *testing.T) {
	type args struct {
		data map[string]map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Get json from valid data",
			args: args{
				data: map[string]map[string]interface{}{
					"1": {
						"name": "ernesto",
						"age":  38,
					},
				},
			},
			want: []byte(`{
  "1": {
    "age": 38,
    "name": "ernesto"
  }
}`),
			wantErr: false,
		},
		{
			name: "Get error from invalid data",
			args: args{
				data: nil,
			},
			want:    []byte{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetJSONFromData(tt.args.data)
			fmt.Println(string(got))
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJSONFromData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if string(got) != string(tt.want) {
				t.Errorf("GetJSONFromData() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestSaveJSONToFile(t *testing.T) {
	type args struct {
		filename string
		data     []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Save json to file success",
			args: args{
				filename: "test.json",
				data:     []byte(`{"name":"ernesto"}`),
			},
			wantErr: false,
		},
		{
			name: "Save json to file error",
			args: args{
				filename: "",
				data:     nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveJSONToFile(tt.args.filename, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SaveJSONToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
