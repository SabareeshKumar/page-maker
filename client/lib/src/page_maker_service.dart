import 'dart:convert';

import 'package:http/http.dart';

import 'api_urls.dart' as api_urls;

class PageMakerService {
  final Client _http;

  PageMakerService(this._http);

  Future<bool> createComponentSkeleton() async {
    final response = await _http.post(api_urls.createComponentSkeleton);
    return response.statusCode != 200;
  }

  Future<bool> createForm(List<int> widgetTypes) async {
    if (widgetTypes.isEmpty) {
      return true;
    }
    final payload = json.encode(widgetTypes);
    final response = await _http.post(api_urls.createForm, body: payload);
    return response.statusCode != 200;
  }
}
