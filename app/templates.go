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

// TODO: Need to interpolate component name during run-time.
// TODO: The indentation is slightly off in generated file.
const componentBody = `
import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

@Component(
  selector: 'my-component',
  templateUrl: '',
  styleUrls: [],
  directives: [],
)
class MyComponent implements OnInit, OnActivate {
  MyComponent();

  @override
  void ngOnInit() {
	print('MyComponent: ngOnInit');
  }

  @override
  void onActivate(RouterState prev, RouterState curr) {
	print('MyComponent: onActivate');
  }
}

`
