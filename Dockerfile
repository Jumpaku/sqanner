
FROM golang:1.20.3-buster AS work-base

ENV DEBIAN_FRONTEND=noninteractive

CMD ["bash"]

WORKDIR /work

RUN apt update && apt install -y git curl jq make sqlite3 python clang

RUN go install github.com/cloudspannerecosystem/spanner-cli@latest
RUN go install golang.org/x/tools/cmd/goimports@latest


# Build and run:
#   docker build -t clion/remote-cpp-env:0.5 -f Dockerfile.remote-cpp-env .
#   docker run -d --cap-add sys_ptrace -p 127.0.0.1:2222:22 --name clion_remote_env clion/remote-cpp-env:0.5
#   ssh-keygen -f "$HOME/.ssh/known_hosts" -R "[localhost]:2222"
#
# stop:
#   docker stop clion_remote_env
#
# ssh credentials (test user):
#   user@password
FROM golang:1.20.3-buster AS work-remote

WORKDIR /
COPY --from=work-base / ./

RUN apt update && apt install -y \
  ssh \
  build-essential \
  autoconf \
  locales-all \
  dos2unix \
  rsync \
  tar \
  python \
  tzdata \
  make \
  && apt clean

## For goland-remote
RUN ( \
  echo 'LogLevel DEBUG2'; \
  echo 'PermitRootLogin yes'; \
  echo 'PasswordAuthentication yes'; \
  echo 'Subsystem sftp /usr/lib/openssh/sftp-server'; \
  ) > /etc/ssh/sshd_config_test_goland \
  && mkdir /run/sshd

RUN useradd -m user \
  && yes password | passwd user

RUN usermod -s /bin/bash user
