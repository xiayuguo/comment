# coding: utf8

import os
import sys
import platform


def execute(cmd, response):
    result = os.system(cmd)
    if result == 0:
        print("%s success." % response)
        return result
    else:
        print("%s fail." % response)
        sys.exit()


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
    execute("git clone https://github.com/hugoxia/comment.git", "Download code")
        
    # 切换目录
    os.chdir("comment")

    # 创建 MongoDB 数据卷
    execute("docker volume create mongo_data_comment", "MongoDB create volume mongo_data_comment")
   
    # 启动 MongoDB
    execute("docker run -d --name mongo-comment -v mongo_data_comment:/data/db -p 29017:27017 mongo", "MongoDB run")

    # 拷贝 MongoDB 备份数据文件
    execute("docker cp comment.dump mongo-comment:/comment.dump", "Copy comment.dump")
    
    # 初始化 MongoDB 数据库
    execute("docker exec mongo-comment bash -c 'mongorestore --db my_db --gzip --archive=/comment.dump'",
            "Init MongoDB data")

    # MongoDB 创建用户 comment
    execute("docker cp create_user.js mongo-comment:/create_user.js && "
            "docker exec mongo-comment bash -c 'mongo admin /create_user.js'",
            "Create user comment")

    # 编译 comment 服务
    execute("docker build -t comment .", "Docker build comment")

    # 启动 comment 服务
    execute("docker run -d --link mongo-comment:mongodb -p 10008:10008 comment", "Comment server run")
    print("Please visit http://127.0.0.1:10008/")


if __name__ == "__main__":
    main()
