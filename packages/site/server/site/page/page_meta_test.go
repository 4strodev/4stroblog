package page

import "testing"

func TestParseFrontMatter(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    PageMeta
		wantErr bool
	}{
		{
			name:    "reads layout",
			content: "{{/*\n+++\nlayout = \"layouts/full-width\"\n+++\n*/}}\n<form></form>",
			want:    PageMeta{Layout: "layouts/full-width"},
		},
		{
			name:    "no front matter",
			content: "<h1>Hello</h1>",
			want:    PageMeta{},
		},
		{
			name:    "comment without delimiters",
			content: "{{/* just a comment */}}<h1>Hello</h1>",
			want:    PageMeta{},
		},
		{
			name:    "front matter must be at the top",
			content: "<h1>Hello</h1>{{/*\n+++\nlayout = \"layouts/full-width\"\n+++\n*/}}",
			want:    PageMeta{},
		},
		{
			name:    "invalid toml",
			content: "{{/*\n+++\nlayout = \n+++\n*/}}",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFrontMatter([]byte(tt.content))
			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLoadPagesMeta(t *testing.T) {
	metas, err := LoadPagesMeta("../../../views")
	if err != nil {
		t.Fatal(err)
	}

	got := metas["pages/admin/post/new"].Layout
	if got != "layouts/full-width" {
		t.Errorf("new post layout = %q, want layouts/full-width", got)
	}
	if _, ok := metas["pages/index"]; !ok {
		t.Error("pages/index missing")
	}
}
