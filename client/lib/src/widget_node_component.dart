import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

import 'package:angular_components/material_input/material_auto_suggest_input.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/material_input/material_input.dart';
import 'package:angular_components/model/selection/string_selection_options.dart';

import 'input_meta.dart';
import 'int_enum.dart';
import 'constants.dart';

@Component(
  selector: 'widget-node',
  templateUrl: 'widget_node_component.html',
  directives: [
    MaterialAutoSuggestInputComponent,
    MaterialButtonComponent,
    MaterialInputComponent,
    NgIf,
    NgModel,
    WidgetNodeComponent,
    materialInputDirectives,
  ],
  styleUrls: [
    'widget_node_component.scss.css',
  ],
)
class WidgetNodeComponent implements OnInit, OnActivate {
  final widgetTypeOptions = StringSelectionOptions<IntEnum>(widgetTypes);
  InputMeta inputMeta = InputMeta();

  bool addedChild = false;

  @ViewChild('nextNode')
  WidgetNodeComponent nextNode;

  @override
  void ngOnInit() {
    print('WidgetNodeComponent: ngOnInit');
  }

  @override
  void onActivate(RouterState prev, RouterState current) {
    print('WidgetNodeComponent: onActivate');
  }

  void addChild() => addedChild = true;
}
