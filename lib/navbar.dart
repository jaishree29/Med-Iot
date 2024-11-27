import 'package:flutter/material.dart';
import 'package:iconsax/iconsax.dart';
import 'package:mediot/utils/constants/colors.dart';
import 'package:mediot/views/data_visualization_screen.dart';
import 'package:mediot/views/device_list_screen.dart';
import 'package:mediot/views/home_screen.dart';

class NavBar extends StatefulWidget {
  const NavBar({super.key});

  @override
  State<NavBar> createState() => _NavBarState();
}

class _NavBarState extends State<NavBar> {
  int _currentIndex = 0;

  final List<Widget> _screens = [
    const HomeScreen(),
    const DeviceListScreen(),
    const DataVisualizationScreen(),
  ];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          _screens[_currentIndex],
          Align(
            alignment: Alignment.bottomCenter,
            child: Container(
              padding: const EdgeInsets.all(15),
              margin: const EdgeInsets.symmetric(horizontal: 24, vertical: 24),
              decoration: const BoxDecoration(
                color: MColors.primaryColor,
                borderRadius: BorderRadius.all(Radius.circular(24)),
              ),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: [
                  _buildNavItem(Iconsax.home_2, 0),
                  _buildNavItem(Iconsax.mobile, 1),
                  _buildNavItem(Iconsax.chart_14, 2),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildNavItem(IconData icon, int index) {
    return GestureDetector(
      onTap: () {
        setState(() {
          _currentIndex = index;
        });
      },
      child: SizedBox(
        width: 70, 
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Stack(
              alignment: Alignment.center,
              children: [
                _buildSliderIndicator(index),
                _currentIndex == index ? const SizedBox(height: 10) : const SizedBox(height: 0,),
              ],
            ),
            const SizedBox(height: 8), 
            Icon(
              icon,
              color: _currentIndex == index ? Colors.white : Colors.white70,
              size: _currentIndex == index ? 24 : 28,
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildSliderIndicator(int index) {
    return AnimatedContainer(
      curve: Curves.linear,
      duration: const Duration(milliseconds: 300),
      height: 4,
      width: _currentIndex == index
          ? 50
          : 0, 
      decoration: BoxDecoration(
        color: _currentIndex == index ? Colors.white : Colors.transparent,
        borderRadius: BorderRadius.circular(10),
      ),
    );
  }
}
