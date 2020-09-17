package app

import (
	"strings"
)

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
