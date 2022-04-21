cmake_minimum_required(VERSION 3.10)

project(RubikGUI VERSION 1.0)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)

find_package(Vulkan REQUIRED)
find_package(glfw3 3.3.7 REQUIRED)

configure_file(RubikGUIConfig.hpp.in include/RubikGUIConfig.hpp)

add_executable(RubikGUI src/main.cpp)

target_include_directories(RubikGUI PUBLIC include)

target_link_libraries(RubikGUI Vulkan::Vulkan)
target_link_libraries(RubikGUI glfw)
