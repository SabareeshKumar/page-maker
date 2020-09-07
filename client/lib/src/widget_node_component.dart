import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';

import 'package:angular_components/material_button/material_button.dart';

@Component(
  selector: 'widget-node',
  templateUrl: 'widget_node_component.html',
  directives: [
    MaterialButtonComponent,
    NgIf,
    WidgetNodeComponent,
  ],
  styleUrls: [
    'widget_node_component.scss.css',
  ],
)
class WidgetNodeComponent implements OnInit, OnActivate {
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
