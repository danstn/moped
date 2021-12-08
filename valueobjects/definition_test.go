package valueobjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const goodDefinition = `
name: "test pipeline"
steps:
  - name: "step-1"
    run: echo test
  - name: "step-2"
    run: | 
      do this
      do that
`

func TestFromYAML(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Definition
		wantErr error
	}{
		{
			name: "basic definition",
			args: args{
				data: []byte(goodDefinition),
			},
			want: Definition{
				Name: "test pipeline",
				Steps: []Step{
					{
						Name: "step-1",
						Run:  "echo test",
					},
					{
						Name: "step-2",
						Run:  "do this\ndo that\n",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefinitionFromYAML(tt.args.data)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
