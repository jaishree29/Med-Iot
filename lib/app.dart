import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:mediot/views/home_screen.dart';

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return AnnotatedRegion<SystemUiOverlayStyle>(value: const SystemUiOverlayStyle(statusBarColor: Colors.white), child: MaterialApp(
      debugShowCheckedModeBanner: false,
      home: const HomeScreen(),
      theme: ThemeData(
        scaffoldBackgroundColor: Colors.white
      ),
    )); 
  }
}