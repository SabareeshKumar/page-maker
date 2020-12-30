package app

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// InputMeta is the widget meta schema in create-form request
type InputMeta struct {
	InputLabel string `json:"input_label"`
	WidgetType int    `json:"widget_type"`
}

// CreateFormSchema is the schema of create-form request data
type CreateFormSchema []InputMeta

type generationMeta struct {
	widgetEncountered map[int]bool
	imports           []string
	directives        []string
	declarations      []string
	templateBody      bytes.Buffer
}

var widgetGeneratorMap = map[int]func(*generationMeta, InputMeta){
	materialInput:            materialInputGenerator,
	materialAutoSuggestInput: autoSuggestInputGenerator,
	materialDatepicker:       datepickerGenerator,
}

func materialInputGenerator(gm *generationMeta, inputMeta InputMeta) {
	varName := mixedCaps(inputMeta.InputLabel)
	declaration := strings.ReplaceAll(
		materialInputDeclaration, "{varName}", varName)
	gm.declarations = append(gm.declarations, declaration)
	template := strings.ReplaceAll(
		materialInputTemplate, "{varName}", varName)
	label := strings.Title(inputMeta.InputLabel)
	template = strings.ReplaceAll(template, "{label}", label)
	gm.templateBody.WriteString(template)
	if gm.widgetEncountered[inputMeta.WidgetType] {
		return
	}
	gm.imports = append(gm.imports, materialInputImport)
	gm.directives = append(gm.directives, materialInputDirectives[:]...)
	gm.widgetEncountered[inputMeta.WidgetType] = true

}

func autoSuggestInputGenerator(gm *generationMeta, inputMeta InputMeta) {
	varName := mixedCaps(inputMeta.InputLabel)
	for _, declaration := range materialAutoSuggestInputDeclarations {
		declaration := strings.ReplaceAll(declaration, "{varName}", varName)
		gm.declarations = append(gm.declarations, declaration)
	}
	template := strings.ReplaceAll(
		materialAutoSuggestInputTemplate, "{varName}", varName)
	label := strings.Title(inputMeta.InputLabel)
	template = strings.ReplaceAll(template, "{label}", label)
	gm.templateBody.WriteString(template)
	if gm.widgetEncountered[inputMeta.WidgetType] {
		return
	}
	gm.imports = append(gm.imports, materialAutoSuggestInputImports[:]...)
	gm.directives = append(gm.directives, materialAutoSuggestInputDirective)
	gm.widgetEncountered[inputMeta.WidgetType] = true
}

func datepickerGenerator(gm *generationMeta, inputMeta InputMeta) {
	varName := mixedCaps(inputMeta.InputLabel)
	declaration := strings.ReplaceAll(
		materialDatepickerDeclaration, "{varName}", varName)
	gm.declarations = append(gm.declarations, declaration)
	template := strings.ReplaceAll(
		materialDatepickerTemplate, "{date}", varName)
	gm.templateBody.WriteString(template)
	if gm.widgetEncountered[inputMeta.WidgetType] {
		return
	}
	gm.imports = append(gm.imports, materialDatepickerImports[:]...)
	gm.directives = append(gm.directives, materialDatepickerDirective)
	gm.widgetEncountered[inputMeta.WidgetType] = true
}

// CreateForm creates angular-dart component and template files with necessary
// imports, declarations and directives of given widgets.
func CreateForm(form CreateFormSchema) error {
	for _, inputMeta := range form {
		wType := inputMeta.WidgetType
		if widgetTypes.hasValue(wType) {
			continue
		}
		errorStr := fmt.Sprintf("Invalid widget type '%d'", wType)
		return errors.New(errorStr)
	}
	gm := generationMeta{widgetEncountered: make(map[int]bool)}
	gm.templateBody.WriteString(spinnerTemplate)
	for _, inputMeta := range form {
		widgetGeneratorMap[inputMeta.WidgetType](&gm, inputMeta)
	}
	imports := append(gm.imports, panelImport, spinnerImport)
	imports = append(imports, angularImports[:]...)
	directives := append(gm.directives, panelDirective)
	directives = append(directives, spinnerDirectives[:]...)
	panelDeclaration := strings.ReplaceAll(
		panelTitleDeclaration, "{expansionPanelName}", expansionPanelName)
	declarations := append(
		gm.declarations, panelDeclaration, spinnerDeclaration)
	tmpl := gm.templateBody.String()
	tmpl = strings.ReplaceAll(
		expansionPanelTemplate, "{body}", tmpl)
	return createComponent(imports, directives, declarations, tmpl)
}
