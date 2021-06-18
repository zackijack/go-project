# go-project

Go project template

## Install

You can install the pre-compiled binary in several different ways.

#### homebrew tap

```shell
$ brew tap zackijack/tap
$ brew install go-project
```

or simply

```shell
$ brew install zackijack/tap/go-project
```

#### deb/rpm

Download the `.deb` or `.rpm` from the [releases page](https://github.com/zackijack/go-project/releases) and install with `dpkg -i` and `rpm -i` respectively.

#### docker

```shell
$ docker pull zackijack/go-project
```

#### go get

```shell
$ go get github.com/zackijack/go-project
```

and ensuring that `$GOPATH/bin` is added to your `$PATH`.

#### manually

Download the pre-compiled binaries from the [releases page](https://github.com/zackijack/go-project/releases) and copy to the desired location.

### Verifying the installation

Once you’ve installed go-project, you can verify it is installed correctly by running `version` command:
```shell
$ go-project version
```

or with docker

```shell
$ docker run --rm zackijack/go-project:latest version
```