 
FROM amazonlinux:latest
RUN yum update -y
RUN yum install tar wget git vim -y
RUN rpm -Uvh http://repos.mesosphere.io/el/7/noarch/RPMS/mesosphere-el-repo-7-1.noarch.rpm
RUN yum -y install mesos marathon chronos
RUN yum -y install mesosphere-zookeeper


