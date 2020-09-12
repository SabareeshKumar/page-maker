import 'int_enum.dart';

class WidgetType {
  static const MATERIAL_INPUT = IntEnum(1, 'Material Input');
  static const MATERIAL_AUTO_SUGGEST_INPUT =
      IntEnum(2, 'Material Auto Suggest Input');
  static const MATERIAL_DATEPICKER = IntEnum(3, 'Material Datepicker');
  static const MATERIAL_RADIO = IntEnum(4, 'Material Radio');
  static const MATERIAL_TOGGLE = IntEnum(5, 'Material Toggle');
  static const MATERIAL_CHECKBOX = IntEnum(5, 'Material Checkbox');
}

const widgetTypes = <IntEnum>[
  WidgetType.MATERIAL_INPUT,
  WidgetType.MATERIAL_AUTO_SUGGEST_INPUT,
  WidgetType.MATERIAL_DATEPICKER,
  WidgetType.MATERIAL_RADIO,
  WidgetType.MATERIAL_TOGGLE,
  WidgetType.MATERIAL_CHECKBOX,
];
