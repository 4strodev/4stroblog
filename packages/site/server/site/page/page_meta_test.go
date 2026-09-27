package page

import "testing"

var testDefaults = PageMeta{Layout: "layouts/main", Title: "Default title"}

func TestParseFrontMatter(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    PageMeta
		wantErr bool
	}{
		{
			name:    "reads layout and title",
			content: "{{/*\n+++\nlayout = \"layouts/full-width\"\ntitle = \"New post\"\n+++\n*/}}\n<form></form>",
			want:    PageMeta{Layout: "layouts/full-width", Title: "New post"},
		},
		{
			name:    "missing title uses default",
			content: "{{/*\n+++\nlayout = \"layouts/full-width\"\n+++\n*/}}",
			want:    PageMeta{Layout: "layouts/full-width", Title: "Default title"},
		},
		{
			name:    "empty title uses default",
			content: "{{/*\n+++\ntitle = \"\"\n+++\n*/}}",
			want:    testDefaults,
		},
		{
			name:    "no front matter uses defaults",
			content: "<h1>Hello</h1>",
			want:    testDefaults,
		},
		{
			name:    "comment without delimiters uses defaults",
			content: "{{/* just a comment */}}<h1>Hello</h1>",
			want:    testDefaults,
		},
		{
			name:    "front matter must be at the top",
			content: "<h1>Hello</h1>{{/*\n+++\nlayout = \"layouts/full-width\"\n+++\n*/}}",
			want:    testDefaults,
		},
		{
			name:    "invalid toml",
			content: "{{/*\n+++\nlayout = \n+++\n*/}}",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFrontMatter([]byte(tt.content), testDefaults)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestLoadPagesMeta(t *testing.T) {
	metas, err := LoadPagesMeta("../../../views", testDefaults)
	if err != nil {
		t.Fatal(err)
	}

	newPost := metas["pages/admin/post/new"]
	if newPost.Layout != "layouts/full-width" || newPost.Title != "New post" {
		t.Errorf("new post meta = %+v", newPost)
	}

	index, ok := metas["pages/index"]
	if !ok {
		t.Fatal("pages/index missing")
	}
	if index != testDefaults {
		t.Errorf("index meta = %+v, want defaults %+v", index, testDefaults)
	}
}
