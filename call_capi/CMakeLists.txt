cmake_minimum_required(VERSION 3.9)
project(call_capi VERSION 0.1 LANGUAGES CXX)
set(CMAKE_CXX_STANDARD 11)

set(PROJECT_SOURCES
        main.cpp
        )

add_executable(call_capi
                ${PROJECT_SOURCES}
                )

target_include_directories(call_capi
        PRIVATE
        libcapi
        )

target_link_libraries(call_capi
        PRIVATE
        ${CMAKE_SOURCE_DIR}/libcapi/libcapi.lib
        )

install(TARGETS call_capi
        RUNTIME DESTINATION .)
install(FILES ${CMAKE_SOURCE_DIR}/libcapi/libcapi.dll
        DESTINATION .)
