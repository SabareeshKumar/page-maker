package app

import (
	"strings"
)

// CreateComponentSkeleton creates a basic angular-dart component with an
// expansion-panel.
func CreateComponentSkeleton() error {
	imports := append(angularImports[:], panelImport, spinnerImport)
	directives := append(spinnerDirectives[:], panelDirective)
	panelDeclaration := strings.ReplaceAll(
		panelTitleDeclaration, "{expansionPanelName}", expansionPanelName)
	declarations := []string{panelDeclaration, spinnerDeclaration}
	tmpl := strings.ReplaceAll(
		expansionPanelTemplate, "{body}", spinnerTemplate)
	return createComponent(imports, directives, declarations, tmpl)
}
