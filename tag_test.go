package defaults

import "testing"

func TestSetTag(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "test set tag",
			input: "default 2",
			want:  "default 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetDefaultTag(tt.input)
			if Tag != tt.want {
				t.Errorf("Tag got %s, want %s", Tag, tt.want)
			}
		})
	}
}
