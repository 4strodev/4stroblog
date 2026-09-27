package core

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeView(t *testing.T, folder, view string) {
	t.Helper()
	path := filepath.Join(folder, view)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("<p></p>"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestCheckRequiredViews(t *testing.T) {
	t.Run("all required views present", func(t *testing.T) {
		folder := t.TempDir()
		writeView(t, folder, "layouts/main.html")
		writeView(t, folder, "pages/index.html")

		if err := checkRequiredViews(folder); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("lists every missing view", func(t *testing.T) {
		folder := t.TempDir()

		err := checkRequiredViews(folder)
		if err == nil {
			t.Fatal("expected error")
		}
		for _, view := range requiredViews {
			if !strings.Contains(err.Error(), view) {
				t.Errorf("error %q does not mention %s", err, view)
			}
		}
	})

	t.Run("one view missing", func(t *testing.T) {
		folder := t.TempDir()
		writeView(t, folder, "layouts/main.html")

		err := checkRequiredViews(folder)
		if err == nil || !strings.Contains(err.Error(), "pages/index.html") {
			t.Fatalf("expected missing pages/index.html, got %v", err)
		}
		if strings.Contains(err.Error(), "layouts/main.html") {
			t.Errorf("error mentions present view: %v", err)
		}
	})

	t.Run("folder does not exist", func(t *testing.T) {
		if err := checkRequiredViews(filepath.Join(t.TempDir(), "nope")); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("project views folder", func(t *testing.T) {
		if err := checkRequiredViews("../../views"); err != nil {
			t.Fatalf("project views invalid: %v", err)
		}
	})
}
