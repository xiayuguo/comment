# coding: utf8

import os
import sys
import platform


def main():
    print("Start...")
    # 检查系统版本, 不支持Windows
    sys_version = platform.system()
    print("Your system version is %s." % sys_version)
    if sys_version == "Wind2ows":
        print("I'm Sorry, don't support Windows system.")
        return
    
    # 检查 Python 版本, 只支持 2.7.X 以后的版本
    python_version = platform.python_version()
    print("Your python version is %s." % python_version)
    if sys.version_info < (2, 7):
        print("I'm Sorry, don't support python lower than 2.7.")
        return
    
    # 检查是否安装了 Docker
    is_install_docker = os.popen("which docker").read()
    if is_install_docker:
        print("Docker has installed.")
    else:
        print("Docker don't install.\nPlease install docker before proceeding with the following operation.")
        return

    # 克隆代码到当前目录
    result = os.system("git clone https://github.com/hugoxia/comment.git && cd comment")
    if result == 0:
        print("Download code success.")
    else:
        print("Download code fail.")
        
    # 创建 MongoDB 数据卷
    result = os.system("docker volume create mongo_data_comment")
    if result == 0:
        print("MongoDB create volume mongo_data_comment success.")
    else:
        print("MongoDB create volume mongo_data_comment fail.")
   
    # 启动 MongoDB
    result = os.system("docker run -d --name mongo-comment -v mongo_data_comment:/data/db -p 29017:27017 "
                       "-e MONGO_INITDB_ROOT_USERNAME=comment -e MONGO_INITDB_ROOT_PASSWORD=comment mongo")
    if result == 0:
        print("MongoDB run success.")
    else:
        print("MongoDB run fail.")

    # 拷贝 MongoDB 备份数据文件
    result = os.system("docker cp comment.dump mongo-comment:/comment.dump")
    if result == 0:
        print("Copy comment.dump success.")
    else:
        print("Copy comment.dump fail.")
    
    # 初始化 MongoDB 数据库
    result = os.system("docker exec mongo-comment bash -c 'mongorestore --db my_db --gzip --archive=/comment.dump'")
    if result == 0:
        print("Init MongoDB data success.")
    else:
        print("Init MongoDB data fail.")

    # 编译 comment 服务
    result = os.system("docker build -t comment .")
    if result == 0:
        print("Docker build comment success.")
    else:
        print("Docker build comment fail.")

    # 启动 comment 服务
    result = os.system("docker run -d --link mongo-comment:mongodb -p 10008:10008 comment")
    if result == 0:
        print("Comment server run success.")
        print("http://127.0.0.1:10008/")
    else:
        print("Comment server run fail.")


if __name__ == "__main__":
    main()
