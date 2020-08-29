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
import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

@Component(
  selector: '{selector}',
  templateUrl: '{templateUrl}',
  styleUrls: [],
  directives: [],
)
class {componentName} implements OnInit, OnActivate {
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

const componentBodyFormatWithPanel = `
import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_components/material_spinner/material_spinner.dart';

@Component(
  selector: '{selector}',
  templateUrl: '{templateUrl}',
  styleUrls: [],
  directives: [
    MaterialExpansionPanel,
    MaterialSpinnerComponent,
    NgIf,
  ],
)
class {componentName} implements OnInit, OnActivate {
  bool dataFetchInProgress = false;
  String title = "{expansionPanelName}";

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

const componentTemplateFormat = `<material-expansionpanel
  alwaysHideExpandIcon expanded disabled [showSaveCancel]="false" 
  [name]="title">
  <div *ngIf="dataFetchInProgress">
    <div>Loading...</div>
    <material-spinner></material-spinner>
  </div>
</material-expansionpanel>
`
