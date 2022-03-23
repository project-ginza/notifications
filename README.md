#Ginza-Notifications

## 1. Introduction

## 2. How to run
First, you need to compile the source code. At the root level, run the command as follows:

`go build -o notifications`

After the command, you will find an executable file named `notifications`. In order to run it, you have to add a flag like this.

`./notifications run -e <environment> >> /data/output.log 2>&1 &`

