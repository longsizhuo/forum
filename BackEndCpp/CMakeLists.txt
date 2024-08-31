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

# # add_custom_command(
# #       OUTPUT "${forum_proto_srcs}" "${forum_proto_hdrs}" "${forum_grpc_srcs}" "${forum_grpc_hdrs}"
# #       COMMAND ${_PROTOBUF_PROTOC}
# #       ARGS --grpc_out "${CMAKE_SOURCE_DIR}/proto"
# #         --cpp_out "${CMAKE_SOURCE_DIR}/proto"
# #         -I "${forum_proto_path}"
# #         --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
# #         "${forum_proto}"
# #       DEPENDS "${forum_proto}")

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

# # Add the generated files to the executable
# add_executable(server main.cc ${forum_proto_srcs} ${forum_grpc_srcs})

# # Link libraries
# target_link_libraries(server yaml-cpp)




# set(CMAKE_PREFIX_PATH "/usr/lib/x86_64-linux-gnu/cmake/grpc")


# find_package(Protobuf REQUIRED)
# find_package(gRPC REQUIRED)


# include_directories("/usr/include/grpc++" "/usr/include/grpcpp")


# target_link_libraries(server ${Protobuf_LIBRARIES} ${gRPC_LIBRARIES})


# set(PROTO_SRCS "${CMAKE_SOURCE_DIR}/proto/forum.pb.cc" "${CMAKE_SOURCE_DIR}/proto/forum.grpc.pb.cc")
# set(PROTO_HDRS "${CMAKE_SOURCE_DIR}/proto/forum.pb.h" "${CMAKE_SOURCE_DIR}/proto/forum.grpc.pb.h")

# # 添加生成命令
# add_custom_command(
#       OUTPUT ${PROTO_SRCS} ${PROTO_HDRS}
#       COMMAND ${Protobuf_PROTOC_EXECUTABLE}
#       ARGS --grpc_out "${CMAKE_SOURCE_DIR}/proto"
#            --cpp_out "${CMAKE_SOURCE_DIR}/proto"
#            -I "${CMAKE_SOURCE_DIR}/proto"
#            --plugin=protoc-gen-grpc="${gRPC_CPP_PLUGIN_EXECUTABLE}"
#            "${CMAKE_SOURCE_DIR}/proto/forum.proto"
#       DEPENDS "${CMAKE_SOURCE_DIR}/proto/forum.proto"
# )

# # 确保生成文件前依赖关系正确
# add_custom_target(ProtoFiles ALL DEPENDS ${PROTO_SRCS} ${PROTO_HDRS})

# # 你的可执行文件依赖于这些生成的文件
# add_executable(server main.cc ${PROTO_SRCS})
# add_dependencies(server ProtoFiles)



# cmake_minimum_required(VERSION 3.5)
# project(server)

# # Find Protobuf and gRPC plugins
# find_program(_PROTOBUF_PROTOC NAMES protoc PATHS /usr/bin)
# find_program(_GRPC_CPP_PLUGIN_EXECUTABLE NAMES grpc_cpp_plugin PATHS /usr/bin)

# if(NOT _PROTOBUF_PROTOC)
#     message(FATAL_ERROR "protoc not found")
# endif()

# if(NOT _GRPC_CPP_PLUGIN_EXECUTABLE)
#     message(FATAL_ERROR "grpc_cpp_plugin not found")
# endif()



# # Include directories for MySQL
# include_directories(/usr/include/mysql)
# link_directories(/usr/lib/mysql)

# # Find yaml-cpp
# find_package(yaml-cpp REQUIRED)



# # Find Protobuf and gRPC
# # set(Protobuf_DIR "/home/ceither/forum/forum_sh/forum/BackEndCpp/include/vcpkg/installed/x64-linux/share/protobuf")

# find_package(Protobuf REQUIRED)
# set(gRPC_DIR "/home/ceither/forum/forum_sh/forum/BackEndCpp/include/vcpkg/installed/x64-linux/share/grpc")
# find_package(gRPC REQUIRED)

# # Include directories for Protobuf and gRPC
# include_directories(${Protobuf_INCLUDE_DIRS})
# include_directories(${gRPC_INCLUDE_DIRS})

# # Proto file
# get_filename_component(forum_proto "${CMAKE_SOURCE_DIR}/proto/forum.proto" ABSOLUTE)
# get_filename_component(forum_proto_path "${forum_proto}" PATH)

# # Generated sources
# set(forum_proto_srcs "${CMAKE_SOURCE_DIR}/proto/forum.pb.cc")
# set(forum_proto_hdrs "${CMAKE_SOURCE_DIR}/proto/forum.pb.h")
# set(forum_grpc_srcs "${CMAKE_SOURCE_DIR}/proto/forum.grpc.pb.cc")
# set(forum_grpc_hdrs "${CMAKE_SOURCE_DIR}/proto/forum.grpc.pb.h")



# # Generate the protobuf and grpc sources
# add_custom_command(
#     OUTPUT "${forum_proto_srcs}" "${forum_proto_hdrs}" "${forum_grpc_srcs}" "${forum_grpc_hdrs}"
#     COMMAND ${_PROTOBUF_PROTOC}
#     ARGS --grpc_out="${CMAKE_SOURCE_DIR}/proto"
#          --cpp_out="${CMAKE_SOURCE_DIR}/proto"
#          -I "${forum_proto_path}"
#          --plugin=protoc-gen-grpc="${_GRPC_CPP_PLUGIN_EXECUTABLE}"
#          "${forum_proto}"
#     DEPENDS "${forum_proto}"
# )

# # Add the generated files to the executable
# add_executable(server main.cc ${forum_proto_srcs} ${forum_grpc_srcs})

# # Link libraries
# target_link_libraries(server
#     yaml-cpp
#     ${Protobuf_LIBRARIES}
#     ${gRPC_LIBRARIES}
#     grpc++
#     grpc
#     protobuf
# )
# target_link_libraries(server ${Protobuf_LIBRARIES} ${gRPC_LIBRARIES})
