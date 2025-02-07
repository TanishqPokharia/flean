package cubit

const CubitTemplate = `import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
part '{{.Lower}}_state.dart';

class {{.Upper}}Cubit extends Cubit<{{.Upper}}State> {
  {{.Upper}}Cubit() : super({{.Upper}}Initial()) {
    
  }
}`
