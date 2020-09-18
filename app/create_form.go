package app

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

type meta struct {
	widgetTypeCount map[int]int
	imports         []string
	directives      []string
	declarations    []string
	templateBody    bytes.Buffer
}

func handleMaterialInputGeneration(m *meta, widgetType int) {
	count, _ := m.widgetTypeCount[widgetType]
	if count == 0 {
		m.imports = append(m.imports, materialInputImport)
		m.directives = append(m.directives, materialInputDirectives[:]...)
	}
	varName := fmt.Sprintf("input%d", count+1)
	declaration := strings.ReplaceAll(
		materialInputDeclaration, "{varName}", varName)
	m.declarations = append(m.declarations, declaration)
	template := strings.ReplaceAll(
		materialInputTemplate, "{ngModel}", varName)
	m.templateBody.WriteString(template)
	m.widgetTypeCount[widgetType] = count + 1
}

func handleMaterialAutoSuggestInputGeneration(m *meta, widgetType int) {
	count, _ := m.widgetTypeCount[widgetType]
	if count == 0 {
		m.imports = append(m.imports, materialAutoSuggestInputImports[:]...)
		m.directives = append(m.directives, materialAutoSuggestInputDirective)
	}
	selectionModelName := fmt.Sprintf("selection%d", count+1)
	selectionOptionsName := fmt.Sprintf("selectionOptions%d", count+1)
	declarationModel := strings.ReplaceAll(
		selectionModelDeclaration, "{selection}", selectionModelName)
	declarationOptions := strings.ReplaceAll(
		selectionOptionsDeclaration, "{selectionOptions}",
		selectionOptionsName)
	m.declarations = append(
		m.declarations, declarationModel, declarationOptions)
	template := strings.ReplaceAll(
		materialAutoSuggestInputTemplate, "{selection}", selectionModelName)
	template = strings.ReplaceAll(
		template, "{selectionOptions}", selectionOptionsName)
	m.templateBody.WriteString(template)
	m.widgetTypeCount[widgetType] = count + 1
}

func handleMaterialDatepickerGeneration(m *meta, widgetType int) {
	count, _ := m.widgetTypeCount[widgetType]
	if count == 0 {
		m.imports = append(m.imports, materialDatepickerImports[:]...)
		m.directives = append(m.directives, materialDatepickerDirective)
	}
	varName := fmt.Sprintf("date%d", count+1)
	declaration := strings.ReplaceAll(
		materialDatepickerDeclaration, "{varName}", varName)
	m.declarations = append(m.declarations, declaration)
	template := strings.ReplaceAll(
		materialDatepickerTemplate, "{date}", varName)
	m.templateBody.WriteString(template)
	m.widgetTypeCount[widgetType] = count + 1
}

var widgetGenerationHandler = map[int]func(*meta, int){
	materialInput:            handleMaterialInputGeneration,
	materialAutoSuggestInput: handleMaterialAutoSuggestInputGeneration,
	materialDatepicker:       handleMaterialDatepickerGeneration,
}

func CreateForm(widgetTypeList []int) error {
	for _, wType := range widgetTypeList {
		if widgetTypes.hasValue(wType) {
			continue
		}
		errorStr := fmt.Sprintf("Invalid widget type '%d'", wType)
		return errors.New(errorStr)
	}
	m := meta{widgetTypeCount: make(map[int]int)}
	m.templateBody.WriteString(spinnerTemplate)
	for _, wType := range widgetTypeList {
		widgetGenerationHandler[wType](&m, wType)
	}
	imports := append(m.imports, panelImport, spinnerImport)
	imports = append(imports, angularImports[:]...)
	directives := append(m.directives, panelDirective)
	directives = append(directives, spinnerDirectives[:]...)
	panelDeclaration := strings.ReplaceAll(
		panelTitleDeclaration, "{expansionPanelName}", expansionPanelName)
	declarations := append(m.declarations, panelDeclaration, spinnerDeclaration)
	tmpl := m.templateBody.String()
	tmpl = strings.ReplaceAll(
		expansionPanelTemplate, "{body}", tmpl)
	return createComponent(imports, directives, declarations, tmpl)
}
