import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

import 'package:angular_components/material_input/material_auto_suggest_input.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_components/model/selection/selection_model.dart';
import 'package:angular_components/model/selection/string_selection_options.dart';

import 'int_enum.dart';
import 'constants.dart';

@Component(
  selector: 'widget-node',
  templateUrl: 'widget_node_component.html',
  directives: [
    MaterialAutoSuggestInputComponent,
    MaterialButtonComponent,
    NgIf,
    WidgetNodeComponent,
  ],
  styleUrls: [
    'widget_node_component.scss.css',
  ],
)
class WidgetNodeComponent implements OnInit, OnActivate {
  final widgetTypeOptions = StringSelectionOptions<IntEnum>(widgetTypes);
  IntEnum widgetType;

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
