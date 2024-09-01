cmake_minimum_required(VERSION 3.5)
project(server)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

find_program(_PROTOBUF_PROTOC /usr/bin/protoc)
find_program(_GRPC_CPP_PLUGIN_EXECUTABLE /usr/bin/grpc_cpp_plugin)


include_directories(/usr/include/mysql)
link_directories(/usr/lib/mysql)

# find_package(yaml-cpp REQUIRED)
find_package(YAML-CPP REQUIRED)
find_library(MYSQL_LIB mysqlclient HINTS "/usr/lib/x86_64-linux-gnu")

if(NOT MYSQL_LIB)
    message(FATAL_ERROR "mysqlclient library not found!")
endif()

# Proto file
get_filename_component(forum_proto "${CMAKE_SOURCE_DIR}/proto/forum.proto" ABSOLUTE)
get_filename_component(forum_proto_path "${forum_proto}" PATH)

# Generated sources
set(forum_proto_srcs "${CMAKE_SOURCE_DIR}/proto/forum.pb.cc")
set(forum_proto_hdrs "${CMAKE_SOURCE_DIR}/proto/forum.pb.h")
set(forum_grpc_srcs "${CMAKE_SOURCE_DIR}/proto/forum.grpc.pb.cc")
set(forum_grpc_hdrs "${CMAKE_SOURCE_DIR}/proto/forum.grpc.pb.h")


add_custom_command(
      OUTPUT "${forum_proto_srcs}" "${forum_proto_hdrs}" "${forum_grpc_srcs}" "${forum_grpc_hdrs}"
      COMMAND ${_PROTOBUF_PROTOC}
      ARGS --grpc_out "${CMAKE_SOURCE_DIR}/proto"
           --cpp_out "${CMAKE_SOURCE_DIR}/proto"
           -I "${forum_proto_path}"
           --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
           "${forum_proto}"
      DEPENDS "${forum_proto}"
)

add_executable(server main.cc)
target_link_libraries(server PRIVATE /usr/lib/x86_64-linux-gnu/libyaml-cpp.so)
target_link_libraries(server PRIVATE /usr/lib/x86_64-linux-gnu/libmysqlclient.so)

