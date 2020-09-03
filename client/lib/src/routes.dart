import 'package:angular_router/angular_router.dart';

import 'route_paths.dart' as paths;

import 'home_component.template.dart' as hct;

class Routes {
  static final _home = RouteDefinition(
      routePath: paths.home, component: hct.HomeComponentNgFactory);
  RouteDefinition get home => _home;

  final all = <RouteDefinition>[
    _home,
  ];
}
