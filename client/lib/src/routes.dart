import 'package:angular_router/angular_router.dart';

import 'route_paths.dart' as paths;

import 'home_component.template.dart' as hct;
import 'create_form_component.template.dart' as cfct;

class Routes {
  static final _home = RouteDefinition(
      routePath: paths.home, component: hct.HomeComponentNgFactory);
  RouteDefinition get home => _home;

  static final _createForm = RouteDefinition(
      routePath: paths.createForm,
      component: cfct.CreateFormComponentNgFactory);
  RouteDefinition get createForm => _createForm;

  final all = <RouteDefinition>[
    _home,
    _createForm,
  ];
}
