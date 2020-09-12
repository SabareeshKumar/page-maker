import 'package:angular/angular.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_router/angular_router.dart';

import 'widget_node_component.dart';
import 'route_paths.dart' as paths;
import 'page_maker_service.dart';

@Component(
  selector: 'create-form',
  templateUrl: 'create_form_component.html',
  directives: [
    MaterialButtonComponent,
    MaterialExpansionPanel,
    WidgetNodeComponent,
  ],
  providers: [
    PageMakerService,
  ],
  styleUrls: [
    'create_form_component.scss.css',
  ],
)
class CreateFormComponent implements OnInit, OnActivate {
  final Router _router;
  final PageMakerService _service;

  @ViewChild('rootWidget')
  WidgetNodeComponent rootWidget;

  CreateFormComponent(this._router, this._service);

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

  Future<void> save() async {
    final widgetTypes = <int>[];
    var currNode = rootWidget;
    while (currNode != null) {
      final selectedWidgetType = currNode.selectedWidgetType;
      if (selectedWidgetType != null) {
        widgetTypes.add(selectedWidgetType.value);
      }
      currNode = currNode.nextNode;
    }
    await _service.createForm(widgetTypes);
  }
}
