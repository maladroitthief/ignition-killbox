---
title: quick-shift readme
description: Details on how to develop and use the quick-shift code
author: Ian Weller
date: 11/23/2020
---

# quick shift

## Before you begin

- [Install git](https://git-scm.com/downloads)
- [Install vscode](https://code.visualstudio.com/)
- [Install gort](https://gort.io/documentation/getting_started/downloads/)

## Upload Firmata to the Arduino

This section assumes you're using an Arduino Uno or another compatible board. If you already have the Firmata sketch installed, you can skip straight to the examples.

### OS X

First plug the Arduino into your computer via the USB/serial port. A dialog box will appear telling you that a new network interface has been detected. Click "Network Preferences…", and when it opens, simply click "Apply".

Once plugged in, use [Gort](http://gort.io)'s `gort scan serial` command to find out your connection info and serial port address:

```bash
gort scan serial
```

Use the `gort arduino install` command to install `avrdude`, this will allow you to upload firmata to the arduino:

```bash
gort arduino install
```

Once the avrdude uploader is installed we upload the firmata protocol to the arduino, use the arduino serial port address found when you ran `gort scan serial`:

```bash
gort arduino upload firmata /dev/tty.usbmodem1421
```

Now you are ready to connect and communicate with the Arduino using serial port connection

Note that Gobot works best with the `tty.` version of the serial port as shown above, not the `cu.` version.

### Ubuntu

First plug the Arduino into your computer via the USB/serial port.

Once plugged in, use [Gort](http://gort.io)'s `gort scan serial` command to find out your connection info and serial port address:

```bash
gort scan serial
```

Use the `gort arduino install` command to install `avrdude`, this will allow you to upload firmata to the arduino:

```bash
gort arduino install
```

Once the avrdude uploader is installed we upload the firmata protocol to the arduino, use the arduino serial port address found when you ran `gort scan serial`, or leave it blank to use the default address `ttyACM0`:

```bash
gort arduino upload firmata /dev/ttyACM0
```

Now you are ready to connect and communicate with the Arduino using serial port connection

### Windows

First download and install gort for your OS from the [gort.io](gort.io) [downloads page](http://gort.io/documentation/getting_started/downloads/) and install it.

Open a command prompt window by right clicking on the start button and choose `Command Prompt (Admin)` (on windows 8.1). Then navigate to the folder where you uncompressed gort (unzip to a folder first if you haven't done this yet).

Once inside the gort folder, first install avrdude which we'll use to upload firmata to the arduino.

```bash
gort arduino install
```

When the installation is complete, close the command prompt window and open a new one. We need to do this for the env variables to reload.

```bash
gort scan serial
```

Take note of your arduino's serialport address (COM1 COM2 COM3 etc). You need to already have installed the arduino drivers from [arduino.cc/en/Main/Software](https://www.arduino.cc/en/Main/Software). Finally upload the firmata protocol sketch to the arduino.

```bash
gort arduino upload firmata <COMX>
```

Make sure to substitute `<COMX>` with the apropiate serialport address.

Now you are ready to connect and communicate with the Arduino using serial port connection.
