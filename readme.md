# Git 笔记
1. 创建:在本地创建个文件夹，然后 git init
2. 添加文件 git add file.txt
3. 提交修改 git commit -m "注释log"
4. 查看状态 git status
5. 查看log  git log  简化版 git log --pretty=oneline
6. 根据log可以回退版本(restore 新版  reset老版)
    -  git restore filename  回退到上一个版本
    - git reset --hard HEAD^^ 回退到上上个版本
    - 上面两个方式注意在windows下需要对^加双引号
    - 可以根据log上一串版本号回退 git reset --hard ef0a3(注：根据log)
    - git restore -s HEAD~1 readme.md
    - git restore -s ef0a3 readme.md
    - git reflog 记录版本号的变化，根据根据这个反复回退
7. 查看两个版本的不同  git diff file.txt
8. 查看工作区和版本库中最新版有啥区别  git diff HEAD -- readme.txt
9. 撤销 命令git restore (--worktree) readme.txt意思就是，把readme.txt文件在工作区的修改全部撤销，这里有两种情况：
    - 一种是readme.txt自修改后还没有被放到暂存区，现在，撤销修改就回到和版本库一模一样的状态；
    - 另一种是readme.txt已经添加到暂存区后，又作了修改，现在，撤销修改就回到添加到暂存区后的状态。
    总之，就是让这个文件回到最近一次git commit或git add时的状态。

       git restore --staged readme.md 

    - 从master同时恢复工作区和暂存区
       ```
       git restore --source=HEAD --staged --worktree readme.txt
       ```
10. 删除文件 
     - 先在本地删除，使用git status 能看到状态
     - 然后使用 git rm readme.txt 彻底删除

11. 远程仓库
    - 首先要创建SSH key ,windows下打开Git Bash
       ```
       ssh-keygen -t rsa -C "youremail@example.com"
        ``` 
      自己项目一般不用输入密码，一路enter，会在用户目录下的.ssh目录下有 id_rsa(私钥)  id_rsa.pub(公钥)
    - 登录github 打开Account settings ,"SSG Keys"中添加id_rsa.pub里面的内容
    - 创建一个仓库 例如 workspace
     ```  
     git remote add origin git@github.com:youraccount/workspace.git 
     ```
     ``` 
     git push -u origin master
     ```
     以后只要本地做了提交，就可以通过下面命令提交
     ```
     git push origin master
     ```
    - 删除远程库
    - 先查看 git remote -v
    -  再取消关联 git remote rm origin
12. 分支的创建与切换
     - 创建并切换分支
      ```
      git switch -c dev
      ```
    - 查看分支
     ```
     git branch
     ```
    - 切换到主分支
     ```
     git switch master
     ```
    - 合并某分支到当前分支（fast forward看不出曾经做过合并）
    ```
    git merge <分支名>
    ```
    - 合并，禁用fast forward,可以看出合并的log
    ```
    git merge --no-ff -m "log" <分支名>
    ```
    - 删除分支
    ```
    git branch -d <name>
    ```
    - 查看分支合并情况
    ```
     git log --graph --pretty=oneline --abbrev-commit
    ```

14. 临时创建bug分支修复
    - 保护现场 ```git stash```
    - 根据某个分支创建issue分支，修改后提交合并再删除issue分支
    - 回到刚刚切换出去的分支，查找现场 ```git stash list```
    - 恢复现场 ``` git stash pop```
    - 如果多次保存现场，恢复的时候 ```git stash apply stash@{0} ``` ```git stash drop ```
    - 如果要把修改的bug合并到现在这个分支 ``` git cherry-pick  <commitid>```

15. down github上项目 ``` git clone git@github.com:tinawangtt/workspace.git ```
    - 这种情况下只能拉到master分支 
    - 远程dev到本地 ``` git checkout -b dev origin/dev```
    - ``` git pull ``` //更新
    - ``` git branch --set-upstream-to=origin/dev dev ``` 设置dev和origin/dev的链接