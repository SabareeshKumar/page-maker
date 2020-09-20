import 'int_enum.dart';

class InputMeta {
  String inputLabel;
  IntEnum widgetType;

  InputMeta();

  Map<String, dynamic> toJson() {
    return <String, dynamic>{
      'input_label': inputLabel,
      'widget_type': widgetType.value,
    };
  }
}
