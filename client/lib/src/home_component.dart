import 'package:angular/angular.dart';
import 'package:angular_components/material_button/material_button.dart';
import 'package:angular_router/angular_router.dart';

import 'page_maker_service.dart';

@Component(
  selector: 'home',
  templateUrl: 'home_component.html',
  directives: [
    MaterialButtonComponent,
  ],
  providers: [
    PageMakerService,
  ],
  styleUrls: [
    'home_component.scss.css',
  ],
)
class HomeComponent implements OnInit, OnActivate {
  final PageMakerService _service;

  HomeComponent(this._service);

  @override
  void ngOnInit() {
    print('HomeComponent: ngOnInit');
  }

  @override
  void onActivate(RouterState prev, RouterState current) {
    print('HomeComponent: onActivate');
  }

  Future<void> createComponentSkeleton() async {
    await _service.createComponentSkeleton();
  }
}
