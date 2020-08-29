package app

const copyrightHeader = "" +
	"// Copyright (C) 2018 HiPro IT Solutions Private Limited, Chennai. All" +
	"\n" +
	"// rights reserved.\n" +
	"//\n" +
	"// This program and the accompanying materials are made available\n" +
	"// under the terms described in the LICENSE file which accompanies\n" +
	"// this distribution. If the LICENSE file was not attached to this\n" +
	"// distribution or for further clarifications, please contact\n" +
	"// legal@hipro.co.in.\n"

const copyrightFooter = "" +
	"// Local Variables:\n" +
	"// mode: dart\n" +
	"// End:\n"

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
