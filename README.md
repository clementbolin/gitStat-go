
![](./assets/git.png)

# GitStat

## Introduction

The goal of this project is to scan your git projects and display as what you have on Github/Gitlab the frequency of your commits in your terminal.
This project is based on this topic : [TOPIC](https://flaviocopes.com/go-git-contributions/)

## How to Use

different flag function. the ```email``` flag is required

![](./assets/help.png)

![Alt Text](./assets/example.gif)

### Install

you must have the Golang compiler to compile this project.
If this is not the case you can go [here](https://golang.org) to install it.

first step is to clone repository

    git clone https://github.com/ClementBolin/gitStat-go
    cd gitStat-go

after clone repository you can start ```install.sh```

    chmod +x install.sh
    ./install.sh

if ```install.sh``` failed you can follow the following instructions

#### Mac OS and Linux

    make build
    sudo cp ./bin/gitStat-go /usr/local/bin
