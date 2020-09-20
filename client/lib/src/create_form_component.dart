import 'package:angular/angular.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_expansionpanel/material_expansionpanel.dart';
import 'package:angular_router/angular_router.dart';

import 'input_meta.dart';
import 'page_maker_service.dart';
import 'route_paths.dart' as paths;
import 'widget_node_component.dart';

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
    final inputMetaList = <InputMeta>[];
    var currNode = rootWidget;
    while (currNode != null) {
      final inputMeta = currNode.inputMeta;
      if (inputMeta != null) {
        inputMetaList.add(inputMeta);
      }
      currNode = currNode.nextNode;
    }
    await _service.createForm(inputMetaList);
  }
}
