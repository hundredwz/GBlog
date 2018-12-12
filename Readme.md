# GBlog

This project is a blog system based on gin framework written by golang.

## Introduction

This project is inspired by [typecho](https://github.com/typecho/typecho) and I use the tables of it and reduce some columns.

The reason why I want to adapt it to golang is mainly due to the fact that I want to use golang to complete a project. Then this project borns.

There are some features about it, although I think they are unnecessary to be presented.

1. MVC
I use mvc architecture to build this system. You can find that from the packages.

2. Ajax submit
Most of the backend operations are executed on ajax.

3. Not fully-supported orm
I write some orm operations for the project. But it's incomplete. However, I think it's enough for the project.

I think the structure of this project is on some conditions, clean. You can easily find what one package works.

## Instructions

The following is the instructions for how to use the system.

You can download the prebuilt binary for you system from the bin package in the repo, unzip it and the run the binary is ok. For other unprovided systems ,you can build it by yourself.

Notice that the binary should be with the web folder in the same folder; otherwise the system cannot find the static resources and failed to start.

It's worth mentioned that you should already installed mysql. The current system only supports mysql as the database because I am too lazy to adapt it to other databases.

Besides, there are some commands about the system.

```bash
GBlog -h
```
this command prints the help list about the system
```bash
GBlog -d
```
this command can run the system as debug mode
```bash
GBlog -p port
```
this command can help run the system with your custom port. The default port is 701.

Anything else?

I think the answer is no.

## Contributions

I work on this project on my spare time. As a result, it's bound that there many bugs and errors in the project. So any suggestions and help are welcomed. Pull-Request is much better.

If you have any problems, you can find me in [Hundred Blog](https://txiner.top) or write to me [Hundred](mailto:sdwangzhuo@gmail.com)

The project is under MIT license.