# fritzctl - console AVM FRITZ!Box client

![fritzctl](/images/fritzctl.png?raw=true "fritzctl")

---
**Project is archived and no longer maintained**

Due to a lack of time to work on this project, I have decided to archive it as of April 2021.
Of course, the source code will stay public.
Finally, I wish to express my deep gratitude to everyone who contributed. Thanks!

---

## Usage

![Demo usage](/images/fritzctl_demo.gif?raw=true "Demo usage")

## About [![Wiki](https://img.shields.io/badge/wiki-home-brightgreen.svg)](https://github.com/gotohr/fritzctl/wiki)

`fritzctl` is a command line client for the AVM FRITZ!Box primarily focused on the
[AVM Home Automation HTTP Interface](https://avm.de/fileadmin/user_upload/Global/Service/Schnittstellen/AHA-HTTP-Interface.pdf).

It should work out-of-the-box with most FRITZ!Boxes running a recent FRITZ!OS version. It has been explicitly tested with

| FRITZ!Box | FRITZ!OS |
| --- | --- |
| FRITZ!Box Fon WLAN 7390 | 06.51, 06.80, 06.83 |
| FRITZ!Box 6490 Cable | 06.63, 06.83, 06.84, 06.87, 07.00 |
| FRITZ!Box 7490 | 06.83, 06.90, 06.93 |

## CI [![Actions](https://github.com/gotohr/fritzctl/workflows/Continuous%20Integration/badge.svg)](https://github.com/gotohr/fritzctl/actions) [![AppVeyor](https://ci.appveyor.com/api/projects/status/k7qqx91w6mja3u7h?svg=true&passingText=Windows%20-%20OK&failingText=Windows%20-%20failed&pendingText=Windows%20-%20pending)](https://ci.appveyor.com/project/gotohr/fritzctl)

## Code [![Go Report Card](https://goreportcard.com/badge/github.com/gotohr/fritzctl)](https://goreportcard.com/report/github.com/gotohr/fritzctl) [![codecov](https://codecov.io/gh/gotohr/fritzctl/branch/master/graph/badge.svg)](https://codecov.io/gh/gotohr/fritzctl) [![codebeat badge](https://codebeat.co/badges/605cf539-21dd-4a60-a892-e0d6da3021fe)](https://codebeat.co/projects/github-com-gotohr-fritzctl) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/356d5568f61e40c3ad430786f766231e)](https://www.codacy.com/app/bjoern.pirnay/fritzctl?utm_source=github.com&utm_medium=referral&utm_content=gotohr/fritzctl&utm_campaign=badger) [![BCH compliance](https://bettercodehub.com/edge/badge/gotohr/fritzctl?branch=master)](https://bettercodehub.com/results/gotohr/fritzctl) [![Maintainability](https://api.codeclimate.com/v1/badges/0dbf66a5ad3c5e059656/maintainability)](https://codeclimate.com/github/gotohr/fritzctl/maintainability) [![CodeFactor](https://www.codefactor.io/repository/github/gotohr/fritzctl/badge)](https://www.codefactor.io/repository/github/gotohr/fritzctl)

## Releases
*   [![release](https://img.shields.io/github/release/gotohr/fritzctl.svg)](https://github.com/gotohr/fritzctl/releases/latest) [![downloads](https://img.shields.io/github/downloads/gotohr/fritzctl/total.svg)](https://github.com/gotohr/fritzctl/releases/latest) github release
*   [![Download .deb](https://api.bintray.com/packages/gotohr/fritzctl_deb/fritzctl/images/download.svg)](https://bintray.com/gotohr/fritzctl_deb/fritzctl/_latestVersion)
    .deb packages
*   [![Download .rpm](https://api.bintray.com/packages/gotohr/fritzctl_rpm/fritzctl/images/download.svg)](https://bintray.com/gotohr/fritzctl_rpm/fritzctl/_latestVersion)
    .rpm packages 
*   [![Download](https://api.bintray.com/packages/gotohr/fritzctl_win/fritzctl/images/download.svg)](https://bintray.com/gotohr/fritzctl_win/fritzctl/_latestVersion)
    .zip windows

## Install

### Debian/Ubuntu

Add the repository (replace `stretch` by `buster`, `jessie`, `wheezy` or `sid` depending on your distribution)

```sh
echo "deb https://dl.bintray.com/gotohr/fritzctl_deb stretch main" | sudo tee -a /etc/apt/sources.list
```

and its signing key

```sh
wget -qO - https://api.bintray.com/users/gotohr/keys/gpg/public.key | sudo apt-key add -
```

The fingerprint of the repository key `3072D/35E71039` is
`93AC 2A3D 418B 9C93 2986  6463 15FC CFC9 35E7 1039`.
Update your local repository data and install

```sh
sudo apt update
sudo apt install fritzctl
```

Upgrades for `fritzctl` will now be detected by `apt update` and can be installed via `apt upgrade`. 

### openSUSE

Add the repository

```sh
wget https://bintray.com/gotohr/fritzctl_rpm/rpm -O bintray-gotohr-fritzctl_rpm.repo && sudo zypper ar -f bintray-gotohr-fritzctl_rpm.repo && rm bintray-gotohr-fritzctl_rpm.repo
```

Update your local repository data and install

```sh
sudo zypper refresh
sudo zypper in fritzctl
```

### MacOS

Install using homebrew

```sh
brew install gotohr/tap/fritzctl
```

### Windows

Windows binaries can found in the [windows directory](https://dl.bintray.com/gotohr/fritzctl_win/).

### From Source

`fritzctl` is go-gettable. Set up a go environment guided by [How To Write Go Code](http://golang.org/doc/code.html)
and then run
```sh
go get github.com/gotohr/fritzctl
```

## As Library [![GoDoc](https://godoc.org/github.com/gotohr/fritzctl?status.svg)](https://godoc.org/github.com/gotohr/fritzctl)

Example:
```go
package main

import "github.com/gotohr/fritzctl/fritz"

func main() { 
	h := fritz.NewHomeAuto(
		fritz.SkipTLSVerify(),
		fritz.Credentials("", "password"),
	)

	err := h.Login()
	if err != nil {
		panic(err)
	}

	h.Off("Socket_Bedroom")
	h.Temp(18.5, "Heating_Bedroom")
}
```

## Reproducing binaries

Versions >= 1.4.16 can be checked for reproducibility. There is a ready-to-go [Dockerfile](docker/build/Dockerfile)
which prepares an appropriate environment. Of course, the setup instructions can applied to any other build platform.
To reproduce a release, prepare a docker image with the help of build-args
```sh
docker build -t fritzctl/build docker/build \
   --build-arg go_version=1.9.2 \
   --build-arg fritzctl_version=1.4.16 \
   --build-arg fritzctl_revision=v1.4.16
```
Building the binaries is done in the container phase:
```sh
docker run --rm -v fritzctl_build_folder:/root/go/src/github.com/gotohr/fritzctl/build fritzctl/build
```
The above command will create a docker volume `fritzctl_build_folder` containing the binaries. Those can be checked for
equality with the distributed ones.

## License [![License](https://img.shields.io/github/license/gotohr/fritzctl.svg)](https://opensource.org/licenses/MIT) [![FOSSA Status](https://app.fossa.io/api/projects/git%2Bhttps%3A%2F%2Fgithub.com%2Fgotohr%2Ffritzctl.svg?type=shield)](https://app.fossa.com/reports/aee04b1a-57e1-4ddc-aee9-f6beaa3fe2e6)

This project is licensed under the terms of the MIT license, see [LICENSE](https://github.com/gotohr/fritzctl/blob/master/LICENSE).

The `fritzctl` image is licensed under the Creative Commons 3.0 Attributions license. It is build upon the following work:

*   The Go gopher was designed by [Renee French](http://reneefrench.blogspot.com/), licensed under the Creative Commons 3.0 Attributions license.
*   The Go gopher w/ patch cable image was created by [Egon Elbre](http://egonelbre.com), licensed under CC0 1.0 Universal.
