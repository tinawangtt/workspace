# Git 笔记
1. 创建:在本地创建个文件夹，然后 git init
2. 添加文件 git add file.txt
3. 提交修改 git commit -m "注释log"
4. 查看状态 git status
5. 查看log  git log  简化版 git log --pretty=oneline
6. 根据log可以回退版本(restore 新版  reset老版)
    1. git restore filename  回退到上一个版本
    2. git reset --hard HEAD^^ 回退到上上个版本
    3. 上面两个方式注意在windows下需要对^加双引号
    4. 可以根据log上一串版本号回退 git reset --hard ef0a3(注：根据log)
    5. git restore -s HEAD~1 readme.md
    6. git restore -s ef0a3 readme.md
    7. git reflog 记录版本号的变化，根据根据这个反复回退
7. 查看两个版本的不同  git diff file.txt
8. 查看工作区和版本库中最新版有啥区别  git diff HEAD -- readme.txt
9. 撤销 命令git restore (--worktree) readme.txt意思就是，把readme.txt文件在工作区的修改全部撤销，这里有两种情况：
    1. 一种是readme.txt自修改后还没有被放到暂存区，现在，撤销修改就回到和版本库一模一样的状态；
    2. 另一种是readme.txt已经添加到暂存区后，又作了修改，现在，撤销修改就回到添加到暂存区后的状态。
    总之，就是让这个文件回到最近一次git commit或git add时的状态。

       git restore --staged readme.md 

    3. 从master同时恢复工作区和暂存区

       git restore --source=HEAD --staged --worktree readme.txt

