import 'dart:convert';

import 'package:http/http.dart';

import 'api_urls.dart' as api_urls;
import 'input_meta.dart';

class PageMakerService {
  final Client _http;

  PageMakerService(this._http);

  Future<bool> createComponentSkeleton() async {
    final response = await _http.post(api_urls.createComponentSkeleton);
    return response.statusCode != 200;
  }

  Future<bool> createForm(List<InputMeta> inputMetaList) async {
    if (inputMetaList.isEmpty) {
      print('Input meta list is empty. Skipping API call...');
      return true;
    }
    final serializedMetaList =
        List.generate(inputMetaList.length, (i) => inputMetaList[i].toJson());
    final payload = json.encode(serializedMetaList);
    final response = await _http.post(api_urls.createForm, body: payload);
    return response.statusCode != 200;
  }
}
