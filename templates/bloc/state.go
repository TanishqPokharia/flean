package bloc

const BlocStateTemplate = `part of '{{.Lower}}_bloc.dart';

@immutable
sealed class {{.Upper}}State {}

final class {{.Upper}}Initial extends {{.Upper}}State {}

final class {{.Upper}}Loading extends {{.Upper}}State {}

final class {{.Upper}}Failure extends {{.Upper}}State {}

final class {{.Upper}}Success extends {{.Upper}}State {}`
