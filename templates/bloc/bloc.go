package bloc

const BlocTemplate = `import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
part '{{.Lower}}_event.dart';
part '{{.Lower}}_state.dart';

class {{.Upper}}Bloc extends Bloc<{{.Upper}}Event, {{.Upper}}State> {
  {{.Upper}}Bloc() : super({{.Upper}}Initial()) {
    on<{{.Upper}}Event>((event, emit) {
        // write your logic
    });
  }
}`
