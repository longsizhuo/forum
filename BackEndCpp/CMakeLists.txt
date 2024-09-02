cmake_minimum_required(VERSION 3.5)
project(server)

set(CMAKE_PREFIX_PATH "/home/ceither/forum/forum_sh/forum/BackEndCpp/include/vcpkg/installed/x64-linux" ${CMAKE_PREFIX_PATH})

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED True)

set(CMAKE_TOOLCHAIN_FILE "/home/ceither/forum/forum_sh/forum/BackEndCpp/include/vcpkg/scripts/buildsystems/vcpkg.cmake")


find_program(_PROTOBUF_PROTOC /usr/bin/protoc)
find_program(_GRPC_CPP_PLUGIN_EXECUTABLE /usr/bin/grpc_cpp_plugin)

# include_directories(/opt/homebrew/Cellar/protobuf/27.3_1/include)
# link_directories(/opt/homebrew/Cellar/protobuf/27.3_1/lib)

include_directories(/usr/include/mysql)
link_directories(/usr/lib/mysql)

# find_package(yaml-cpp REQUIRED)
find_package(YAML-CPP REQUIRED)
find_library(MYSQL_LIB mysqlclient HINTS "/usr/lib/x86_64-linux-gnu")

if(NOT MYSQL_LIB)
    message(FATAL_ERROR "mysqlclient library not found!")
endif()

# 添加以下两行以找到 Protobuf 和 gRPC 的库
# set(Protobuf_DIR "/home/ceither/forum/forum_sh/forum/BackEndCpp/include/vcpkg/installed/x64-linux/share/protobuf")  # 设置 Protobuf 的路径
# find_package(Protobuf REQUIRED)  # 新增：查找 Protobuf 库
# set(gRPC_DIR "/home/ceither/forum/forum_sh/forum/BackEndCpp/include/vcpkg/installed/x64-linux/share/grpc")  # 设置 gRPC 的路径
find_package(gRPC REQUIRED)       # 新增：查找 gRPC 库

include_directories(${Protobuf_INCLUDE_DIRS})  # 新增：包含 Protobuf 头文件
include_directories(${gRPC_INCLUDE_DIRS})      # 新增：包含 gRPC 头文件

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

# add_executable(server main.cc)
# target_link_libraries(server PRIVATE /usr/lib/x86_64-linux-gnu/libyaml-cpp.so)
# target_link_libraries(server PRIVATE /usr/lib/x86_64-linux-gnu/libmysqlclient.so)
# 修改以下行，将生成的 proto 文件添加到可执行文件的源文件中
add_executable(server main.cc ${forum_proto_srcs} ${forum_grpc_srcs})  # 修改：添加生成的 .cc 文件

# 在链接库中添加 Protobuf 和 gRPC 库
target_link_libraries(server 
    PRIVATE 
    /usr/lib/x86_64-linux-gnu/libyaml-cpp.so
    /usr/lib/x86_64-linux-gnu/libmysqlclient.so
    ${Protobuf_LIBRARIES}  # 新增：链接 Protobuf 库
    ${gRPC_LIBRARIES}      # 新增：链接 gRPC 库
)
