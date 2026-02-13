# Example Dockerfile for OTC FunctionGraph Container Event Function 
# using Ubuntu base image
FROM amd64/ubuntu:22.04

# set environment variables
ENV HOME=/home/paas

ENV GROUP_ID=1003
ENV GROUP_NAME=paas_user

ENV USER_ID=1003
ENV USER_NAME=paas_user

# Set timzone to UTC
ENV TZ=Etc/UTC

# create non root user and home directory
RUN mkdir -m 550 ${HOME} && \
   # add group and user with specific IDs
   groupadd -g ${GROUP_ID} ${GROUP_NAME} && \
   # add user with specific UID and GID
   useradd -u ${USER_ID} -g ${GROUP_ID} ${USER_NAME}

# copy content of target folder
ADD ./target ${HOME}

# adjust permissions
RUN chown -R ${USER_ID}:${GROUP_ID} ${HOME} && \
    chmod -R 775 ${HOME}

# switch to non root user
USER ${USER_NAME}

# switch to working directory
WORKDIR ${HOME}

# expose port and entrypoint
# FunctionGraph invokes container event function through port 8000
EXPOSE 8000

ENTRYPOINT ["/home/paas/httpfunction"]
