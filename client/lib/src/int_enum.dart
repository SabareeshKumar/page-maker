class IntEnum {
  final int value;
  final String name;

  const IntEnum(this.value, this.name);

  @override
  String toString() => name;

  @override
  bool operator ==(et) {
    if (et is IntEnum) {
      return et.value == value && et.name == name;
    }
    if (et is int) {
      return et == value;
    }
    return false;
  }

  @override
  int get hashCode => value.hashCode;
}
