package core

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// requiredViews are the templates every views folder must provide.
// Users can define their own UI, but the site cannot render without these.
var requiredViews = []string{
	"layouts/main.html",
	"pages/index.html",
}

// checkRequiredViews returns an error listing every required template missing in folder.
func checkRequiredViews(folder string) error {
	info, err := os.Stat(folder)
	if err != nil {
		return fmt.Errorf("views folder %q not found: %w", folder, err)
	}
	if !info.IsDir() {
		return fmt.Errorf("views folder %q is not a directory", folder)
	}

	missing := []string{}
	for _, view := range requiredViews {
		_, err := os.Stat(filepath.Join(folder, view))
		if errors.Is(err, os.ErrNotExist) {
			missing = append(missing, view)
			continue
		}
		if err != nil {
			return fmt.Errorf("cannot check view %q: %w", view, err)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("views folder %q is missing required templates: %s", folder, strings.Join(missing, ", "))
	}
	return nil
}
