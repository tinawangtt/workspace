# Git 笔记
1. 创建:在本地创建个文件夹，然后 git init
2. 添加文件 git add file.txt
3. 提交修改 git commit -m "注释log"
4. 查看状态 git status
5. 查看log  git log  简化版 git log --pretty=oneline
6. 根据log可以回退版本
    1. git reset --hard HEAD^  回退到上一个版本
    2. git reset --hard HEAD^^ 回退到上上个版本
    3. 上面两个方式注意在windows下需要对^加双引号
    4. 可以根据log上一串版本号回退 git reset --hard ef0a3(注：根据log)
    5. git reflog 记录版本号的变化，根据根据这个反复回退