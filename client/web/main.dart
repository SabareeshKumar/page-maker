import 'package:angular/angular.dart';
import 'package:angular_router/angular_router.dart';
import 'package:http/http.dart';
import 'package:http/browser_client.dart';

// ignore: uri_has_not_been_generated
import 'package:page_maker/app_component.template.dart' as app;

// ignore: uri_has_not_been_generated
import 'main.template.dart' as self;

@GenerateInjector([
  routerProviders,
  ClassProvider(Client, useClass: BrowserClient),
])
final InjectorFactory rootInjector = self.rootInjector$Injector;

void main() {
  runApp(app.AppComponentNgFactory, createInjector: rootInjector);
}
