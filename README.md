
# fss

[![MIT License](http://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/suntong/fss?status.svg)](http://godoc.org/github.com/suntong/fss)
[![Go Report Card](https://goreportcard.com/badge/github.com/suntong/fss)](https://goreportcard.com/report/github.com/suntong/fss)
[![travis Status](https://travis-ci.org/suntong/fss.svg?branch=master)](https://travis-ci.org/suntong/fss)

## TOC
- [fss - file system summary](#fss---file-system-summary)
- [Usage](#usage)
  - [$ fss](#-fss)
  - [$ fss grps](#-fss-grps)
  - [$ fss dirs](#-fss-dirs)
- [Examples](#examples)
  - [$ fss grps .](#-fss-grps-)
  - [$ ls -ld `ls -1d /etc/* | fss dirs`](#-ls--ld-`ls--1d-etc*--fss-dirs`)
- [Installation](#installation)

## fss - file system summary

The `fss` is a tool to summarize files and diretories in the file system.


# Usage

### $ fss
```sh
file system summary
built on 2017-07-22

Tool to summarize files and diretories in the file system

Usage:
  fss COMMAND

Options:

  -h, --help      display help information
  -v, --verbose   Verbose mode (Multiple -v options increase the verbosity.)

Commands:

  grps   Separate files into groups
  dirs   Pick out dirs from the input
```

### $ fss grps
```sh
Separate files into groups

List files in groups with total sizes from the given path

Usage:
  fss grps SRC-PATH

Options:

  -h, --help        display help information
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity.)
  -l, --limit[=3]   limit listing the files in the same group to
```

### $ fss dirs
```sh
sh
```


# Examples

### $ fss grps .
```sh
279      0M [.gitignore]: .gitignore
         1064      0M [LICENSE]: LICENSE
          606      0M [README.]: README.e.md, README.md
         1518      0M [cmdDirs.go]: cmdDirs.go
         2553      0M [cmdGrps.go]: cmdGrps.go
      6880409      6M [fss]: fss, fss.yaml, fssCLIDef.go, ...
```

### $ ls -ld `ls -1d /etc/* | fss dirs`
```sh
drwxr-xr-x  3 root  root   4096 Apr 11 23:14 /etc/acpi
drwxr-xr-x  2 root  root  12288 Jul  1 11:54 /etc/alternatives
drwxr-xr-x  6 root  root   4096 Apr 11 23:10 /etc/apm
drwxr-xr-x  3 root  root   4096 May 22 17:20 /etc/apparmor
drwxr-xr-x  7 root  root   4096 Jun 15 10:05 /etc/apparmor.d
drwxr-xr-x  5 root  root   4096 May 22 17:16 /etc/apport
drwxr-xr-x  6 root  root   4096 May 22 17:43 /etc/apt
drwxr-xr-x  3 root  root   4096 Apr 11 23:16 /etc/avahi
drwxr-xr-x  2 root  root   4096 May 22 19:07 /etc/bash_completion.d
drwxr-xr-x  2 root  root   4096 Mar 28 11:59 /etc/binfmt.d
drwxr-xr-x  3 root  root   4096 Apr 11 23:10 /etc/ca-certificates
drwxr-xr-x  2 root  root   4096 Apr 11 23:14 /etc/calendar
drwxr-xr-x  4 root  root   4096 Jul  4 11:36 /etc/chromium-browser
drwxr-xr-x  2 root  root   4096 Apr 11 23:17 /etc/compizconfig
drwxr-xr-x  2 root  root   4096 Apr 11 23:08 /etc/console-setup
drwxr-xr-x  2 root  root   4096 May 28 15:53 /etc/cron.d
drwxr-xr-x  2 root  root   4096 May 22 19:16 /etc/cron.daily
drwxr-xr-x  2 root  root   4096 Apr 11 23:08 /etc/cron.hourly
drwxr-xr-x  2 root  root   4096 Apr 11 23:14 /etc/cron.monthly
drwxr-xr-x  2 root  root   4096 May 22 16:46 /etc/cron.weekly
drwxr-xr-x  5 root  lp     4096 Jul 21 07:35 /etc/cups
drwxr-xr-x  2 root  root   4096 Apr 11 23:16 /etc/cupshelpers
drwxr-xr-x  2 root  root   4096 Sep  3  2016 /etc/dbab
drwxr-xr-x  4 root  root   4096 Apr 11 23:07 /etc/dbus-1
drwxr-xr-x  3 root  root   4096 Jul  4 11:36 /etc/default
drwxr-xr-x  2 root  root   4096 Apr 11 23:08 /etc/depmod.d
drwxr-xr-x  4 root  root   4096 Apr 11 23:08 /etc/dhcp
drwxr-xr-x  2 root  root   4096 May 22 17:16 /etc/dictionaries-common
drwxr-xr-x  2 root  root   4096 Nov 30  2015 /etc/dnsmasq.d
drwxr-xr-x  4 root  root   4096 Jul  1 11:54 /etc/dpkg
drwxr-xr-x  3 root  root   4096 Apr 11 23:13 /etc/emacs
drwxr-xr-x  2 root  root   4096 May 13  2016 /etc/extlinux.d
drwxr-xr-x  2 root  root   4096 Jun 15 10:05 /etc/firefox
drwxr-xr-x  4 root  root   4096 Apr 11 23:14 /etc/fonts
drwxr-xr-x  2 root  root   4096 May 22 19:07 /etc/gamin
drwxr-xr-x  5 root  root   4096 Apr 11 23:11 /etc/gconf
drwxr-xr-x  4 root  root   4096 Apr 11 23:11 /etc/ghostscript
drwxr-xr-x  2 root  root   4096 Apr 11 23:13 /etc/gnome
drwxr-xr-x  2 root  root   4096 Apr 11 23:17 /etc/gnome-system-tools
drwxr-xr-x  2 root  root   4096 Apr 11 23:13 /etc/groff
drwxr-xr-x  2 root  root   4096 May 22 10:09 /etc/grub.d
drwxr-xr-x  3 root  root   4096 Apr 11 23:10 /etc/gss
drwxr-xr-x  2 root  root   4096 Apr 11 23:13 /etc/gtk-2.0
drwxr-xr-x  2 root  root   4096 Apr 11 23:16 /etc/gtk-3.0
drwxr-xr-x  2 root  root   4096 Apr  4 18:34 /etc/guest-session
drwxr-xr-x  2 root  root   4096 Apr 11 23:16 /etc/hp
drwxr-xr-x  3 root  root   4096 Apr 11 23:13 /etc/ifplugd
drwxr-xr-x  2 root  root   4096 May 30 13:44 /etc/ImageMagick-6
drwxr-xr-x  2 root  root   4096 Jun 28 07:03 /etc/init
drwxr-xr-x  2 root  root   4096 Jun 28 07:03 /etc/init.d
drwxr-xr-x  5 root  root   4096 Apr 11 23:08 /etc/initramfs-tools
drwxr-xr-x  2 root  root   4096 May 24 21:31 /etc/insserv.conf.d
drwxr-xr-x  3 root  root   4096 Apr 11 23:08 /etc/iproute2
drwxr-xr-x  2 root  root   4096 May 22 18:36 /etc/jed.d
drwxr-xr-x  6 root  root   4096 May 13  2016 /etc/kernel
drwxr-xr-x  2 root  root   4096 Jun  2 06:56 /etc/ldap
drwxr-xr-x  2 root  root   4096 Jun 20 05:59 /etc/ld.so.conf.d
drwxr-xr-x  2 root  root   4096 Jun  7 03:03 /etc/libnl-3
drwxr-xr-x  2 root  root   4096 Nov  4  2015 /etc/libpaper.d
drwxr-xr-x  3 root  root   4096 May 22 19:17 /etc/lightdm
drwxr-xr-x  4 root  root   4096 May 29 19:05 /etc/lighttpd
drwxr-xr-x  6 root  root   4096 May 22 18:35 /etc/logcheck
drwxr-xr-x  2 root  root   4096 May 28 15:52 /etc/logrotate.d
drwxr-xr-x  2 root  root   4096 May 22 19:07 /etc/lynx
drwxr-sr-x  7 smmta smmsp  4096 May 29 19:20 /etc/mail
drwxr-xr-x  3 root  root   4096 Apr 11 23:11 /etc/mate-settings-daemon
drwxr-xr-x  2 root  root   4096 Apr 11 23:13 /etc/menu
drwxr-xr-x  2 root  root   4096 May 22 10:07 /etc/menu-methods
drwxr-xr-x  2 root  root   4096 May 29 18:37 /etc/modprobe.d
drwxr-xr-x  2 root  root   4096 Jun 28 07:03 /etc/modules-load.d
drwxr-xr-x  2 root  root   4096 May 28 16:34 /etc/mpv
drwxr-xr-x  3 root  root   4096 Apr 11 23:13 /etc/mysql
drwxr-xr-x  2 root  root   4096 Apr 11 23:21 /etc/netplan
drwxr-xr-x  7 root  root   4096 Apr 11 23:08 /etc/network
drwxr-xr-x  8 root  root   4096 Apr 11 23:16 /etc/NetworkManager
drwxr-xr-x  2 root  root   4096 Apr 11 23:08 /etc/newt
drwxr-xr-x  2 root  root   4096 May 28 16:34 /etc/openal
drwxr-xr-x  2 root  root   4096 Apr 11 23:07 /etc/opt
drwxr-xr-x  2 root  root   4096 Apr 11 23:16 /etc/PackageKit
drwxr-xr-x  2 root  root   4096 Jun 28 07:03 /etc/pam.d
drwxr-xr-x  2 root  root   4096 Apr 11 23:13 /etc/pcmcia
drwxr-xr-x  4 root  root   4096 Apr 11 23:13 /etc/perl
drwxr-xr-x  4 root  root   4096 Apr 11 23:11 /etc/pki
drwxr-xr-x  5 root  root   4096 Apr 11 23:13 /etc/pm
drwxr-xr-x  5 root  root   4096 Apr 11 23:10 /etc/polkit-1
drwxr-xr-x  4 root  root   4096 May 22 18:35 /etc/ppp
drwxr-xr-x  2 root  root   4096 May 22 17:43 /etc/profile.d
drwxr-xr-x  2 root  root   4096 Apr 11 23:16 /etc/pulse
drwxr-xr-x  2 root  root   4096 Apr 11 23:10 /etc/python
drwxr-xr-x  2 root  root   4096 Apr 11 23:10 /etc/python2.7
drwxr-xr-x  2 root  root   4096 Apr 11 23:08 /etc/python3
drwxr-xr-x  2 root  root   4096 Apr 11 23:07 /etc/python3.5
drwxr-xr-x  2 root  root   4096 May 27 15:27 /etc/rc0.d
drwxr-xr-x  2 root  root   4096 May 27 15:27 /etc/rc1.d
drwxr-xr-x  2 root  root   4096 May 27 15:27 /etc/rc2.d
drwxr-xr-x  2 root  root   4096 May 27 15:27 /etc/rc3.d
drwxr-xr-x  2 root  root   4096 May 27 15:27 /etc/rc4.d
drwxr-xr-x  2 root  root   4096 May 27 15:27 /etc/rc5.d
drwxr-xr-x  2 root  root   4096 May 27 15:27 /etc/rc6.d
drwxr-xr-x  2 root  root   4096 May 22 19:07 /etc/rcS.d
drwxr-xr-x  2 root  root   4096 May 22 19:07 /etc/request-key.d
drwxr-xr-x  4 root  root   4096 May 24 21:31 /etc/resolvconf
drwxr-xr-x  2 root  root   4096 May 22 17:26 /etc/rsyslog.d
drwxr-xr-x  3 root  root   4096 May 22 16:34 /etc/sane.d
drwxr-xr-x  4 root  root   4096 Apr 11 23:14 /etc/security
drwxr-xr-x  2 root  root   4096 Apr 11 23:14 /etc/sensors.d
drwxr-xr-x  2 root  root   4096 Apr 11 23:15 /etc/sgml
drwxr-xr-x  2 root  root   4096 May 22 19:16 /etc/skel
drwxr-xr-x  4 root  root   4096 May 22 19:06 /etc/smartmontools
drwxr-xr-x  3 root  root   4096 Apr 11 23:12 /etc/sound
drwxr-xr-x  4 root  root   4096 Apr 11 23:16 /etc/speech-dispatcher
drwxr-xr-x  2 root  root   4096 May 24 21:49 /etc/squid
drwxr-xr-x  2 root  root   4096 Nov  7  2015 /etc/ssh
drwxr-xr-x  4 root  root   4096 Apr 11 23:14 /etc/ssl
drwxr-xr-x  2 root  root   4096 May 30 13:44 /etc/sudoers.d
drwxr-xr-x  2 root  root   4096 Jun 28 07:03 /etc/sysctl.d
drwxr-xr-x  5 root  root   4096 Jun 28 07:03 /etc/systemd
drwxr-xr-x  2 root  root   4096 Apr 11 23:07 /etc/terminfo
drwxr-xr-x  2 root  root   4096 May 22 19:17 /etc/thermald
drwxr-xr-x  2 root  root   4096 Mar 28 11:59 /etc/tmpfiles.d
drwxr-xr-x  4 root  root   4096 Jun 28 07:03 /etc/udev
drwxr-xr-x  2 root  root   4096 Mar 29 11:01 /etc/udisks2
drwxr-xr-x  3 root  root   4096 May 22 17:26 /etc/ufw
drwxr-xr-x  2 root  root   4096 May 22 16:46 /etc/update-motd.d
drwxr-xr-x  2 root  root   4096 Apr 11 23:15 /etc/UPower
drwxr-xr-x  2 root  root   4096 Jan 23 02:14 /etc/usb_modeswitch.d
drwxr-xr-x  2 root  root   4096 Apr 11 23:15 /etc/wpa_supplicant
drwxr-xr-x 12 root  root   4096 May 22 10:07 /etc/X11
drwxr-xr-x  5 root  root   4096 Apr 11 23:17 /etc/xdg
drwxr-xr-x  2 root  root   4096 Apr 11 23:14 /etc/xml
drwxr-xr-x  2 root  root   4096 Apr 11 23:17 /etc/xrdb
```


# Installation

```
go get github.com/suntong/fss
```

All patches welcome.
