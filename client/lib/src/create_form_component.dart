import 'package:angular/angular.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_router/angular_router.dart';

import 'widget_node_component.dart';
import 'route_paths.dart' as paths;

@Component(
  selector: 'create-form',
  templateUrl: 'create_form_component.html',
  directives: [
    MaterialButtonComponent,
    MaterialExpansionPanel,
    WidgetNodeComponent,
  ],
  styleUrls: [
    'create_form_component.scss.css',
  ],
)
class CreateFormComponent implements OnInit, OnActivate {
  final Router _router;

  @ViewChild('rootWidget')
  WidgetNodeComponent rootWidget;

  CreateFormComponent(this._router);

  @override
  void ngOnInit() {
    print('CreateFormComponent: ngOnInit');
  }

  @override
  void onActivate(RouterState prev, RouterState current) {
    print('CreateFormComponent: onActivate');
  }

  void goToHome() {
    _router.navigate(paths.home.toUrl());
  }

  void save() {}
}
