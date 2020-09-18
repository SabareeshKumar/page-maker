package app

const copyrightHeader = `// Copyright (C) 2018-2020 HiPro IT Solutions Private Limited,
// Chennai. All rights reserved.
//
// This program and the accompanying materials are made available
// under the terms described in the LICENSE file which accompanies
// this distribution. If the LICENSE file was not attached to this
// distribution or for further clarifications, please contact
// legal@hipro.co.in.
`
const copyrightFooter = `// Local Variables:
// mode: dart
// End:
`
const componentBodyFormat = `
{imports}

@Component(
  selector: '{selector}',
  templateUrl: '{templateUrl}',
  styleUrls: [],
  directives: [
    {directives}
  ],
)
class {componentName} implements OnInit, OnActivate {
  {declarations}

  {componentName}();

  @override
  void ngOnInit() {
    print('{componentName}: ngOnInit');
  }

  @override
  void onActivate(RouterState prev, RouterState current) {
    print('{componentName}: onActivate');
  }
}

`
const panelImport = ("import 'package:angular_components" +
	"/material_expansionpanel/material_expansionpanel.dart'")
const panelDirective = "MaterialExpansionPanel"
const panelTitleDeclaration = "String title = '{expansionPanelName}'"
const expansionPanelTemplate = `<material-expansionpanel
  alwaysHideExpandIcon expanded disabled [showSaveCancel]="false" 
  [name]="title">
  {body}</material-expansionpanel>
`
const spinnerImport = ("" +
	"import 'package:angular_components/material_spinner/material_spinner.dart'")
const spinnerDeclaration = "bool dataFetchInProgress = false"
const spinnerTemplate = `<div *ngIf="dataFetchInProgress">
  <div>Loading...</div>
  <material-spinner></material-spinner>
</div>
`
const materialInputImport = ("" +
	"import 'package:angular_components/material_input/material_input.dart'")
const materialInputDeclaration = "String {varName}"
const materialInputTemplate = `<div>
  <material-input
    label="Enter text" type="text" [(ngModel)]="{ngModel}" floatingLabel>
  </material-input>
</div>
`
const materialAutoSuggestInputDirective = "MaterialAutoSuggestInputComponent"
const materialAutoSuggestInputTemplate = `<div>
  <material-auto-suggest-input
    label="Search text" [selection]="{selection}"
    [selectionOptions]="{selectionOptions}" [displayBottomPanel]="false"
    floatingLabel filterSuggestions showClearIcon popupMatchInputWidth>
  </material-auto-suggest-input>
</div>
`
const selectionModelDeclaration = "final {selection} = SelectionModel.single()"
const selectionOptionsDeclaration = ("" +
	"final {selectionOptions} = StringSelectionOptions([])")
const materialDatepickerDirective = "MaterialDatepickerComponent"
const materialDatepickerDeclaration = "Date {varName}"
const materialDatepickerTemplate = `<div>
  <material-datepicker
    selectDatePlaceHolderMsg="Select date" [(date)]="{date}">
  </material-datepicker>
</div>
`

var angularImports = [...]string{
	"import 'package:angular/angular.dart'",
	"import 'package:angular_router/angular_router.dart'",
}
var spinnerDirectives = [...]string{"NgIf", "MaterialSpinnerComponent"}
var materialInputDirectives = [...]string{
	"materialInputDirectives",
	"MaterialInputComponent",
	"NgModel",
}
var materialAutoSuggestInputImports = [...]string{
	("import 'package:angular_components/" +
		"material_input/material_auto_suggest_input.dart'"),
	"import 'package:angular_components/model/selection/selection_model.dart'",
	("import 'package:angular_components" +
		"/model/selection/string_selection_options.dart'"),
}
var materialDatepickerImports = [...]string{
	("import 'package:angular_components" +
		"/material_datepicker/material_datepicker.dart'"),
	"import 'package:angular_components/model/date/date.dart'",
}
