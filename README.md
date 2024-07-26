# Why?

Star allows to quickly navigate through file system without manually pressing `cd` each time

# How?

Star creates `.star` file in home directory where all favorite files and directories are written

File can be manually edited, although it is not recommened

# How to use?

> Star required Fuzzy finder (`fzf`) install on your computer

To star current working directory just enter

```
star
```

To remove star, simply enter command again

To star file or other directory simply enter it as argument

```
star file.txt
```

To view starred files use `-s` flag:

```
star -s
```

Output of this function can be redirected to your favourite editor

You can also specify the editor directly with `-e` flag:

```
star -s -e nvim
```

To search and then delete star you can use `-d` flag:

```
star -s -d
```

# How to build?

To build Star, copy project from github, navigate to project directory and enter

```
sudo make build
```

This command will create star binary inside `/usr/local/bin` directory
