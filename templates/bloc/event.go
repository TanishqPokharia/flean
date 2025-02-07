package bloc

const BlocEventTemplate = `part of '{{.Lower}}_bloc.dart';

@immutable
sealed class {{.Upper}}Event {}
`
