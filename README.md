# dockgo
将go语言程序打包进入docker的例子

1.大docker文件：对应LDockfile，拉取golang编译环境镜像，内部编译输出  
构建命令： podman build -t go_docker_big:latest -f .\LDockfile .  
输出镜像：  
REPOSITORY                TAG         IMAGE ID      CREATED         SIZE  
localhost/go_docker_big   latest      e57621607524  12 seconds ago  591 MB  
创建容器： podman run -d --name GoBigDocker -p 8089:8089 localhost/go_docker_big  
CONTAINER ID  IMAGE                           COMMAND     CREATED        STATUS        PORTS                   NAMES  
5ff86790117f  localhost/go_docker_big:latest              4 seconds ago  Up 4 seconds  0.0.0.0:8089->8089/tcp  GoBigDocker  
  
2.小docker文件：对应SDockfile，外部交叉编译好二进制程序，拉去简单OS镜像打包到一起   
构建命令： podman build -t go_docker_small:latest -f .\SDockfile .  
输出镜像：  
REPOSITORY                 TAG         IMAGE ID      CREATED         SIZE  
localhost/go_docker_small  latest      f3163e980cba  8 seconds ago   19.9 MB  
创建容器：podman run -d --name GoSmallDocker -p 8089:8089 localhost/go_docker_small  
CONTAINER ID  IMAGE                             COMMAND     CREATED        STATUS        PORTS                   NAMES  
58ec252b8649  localhost/go_docker_small:latest              2 seconds ago  Up 2 seconds  0.0.0.0:8089->8089/tcp  GoSmallDocker  
