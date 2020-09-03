import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:angular_components/angular_components.dart';

import 'src/routes.dart';

@Component(
  selector: 'my-app',
  templateUrl: 'app_component.html',
  directives: [
    routerDirectives,
  ],
  providers: [
    ClassProvider(Routes),
    <dynamic>[materialProviders],
  ],
)
class AppComponent implements OnInit {
  final Routes routes;

  AppComponent(this.routes);

  @override
  void ngOnInit() {
    print('AppComponent initialized');
  }
}
