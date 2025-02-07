package cubit

const CubitStateTemplate = `part of '{{.Lower}}_cubit.dart';

@immutable
sealed class {{.Upper}}State {}

final class {{.Upper}}Initial extends {{.Upper}}State {}

final class {{.Upper}}Loading extends {{.Upper}}State {}

final class {{.Upper}}Failure extends {{.Upper}}State {}

final class {{.Upper}}Success extends {{.Upper}}State {}`
