import 'package:http/http.dart';

import 'api_urls.dart' as api_urls;

class PageMakerService {
  final Client _http;

  PageMakerService(this._http);

  Future<bool> createComponentSkeleton() async {
    final response = await _http.post(api_urls.createComponentSkeleton);
    return response.statusCode != 200;
  }
}
