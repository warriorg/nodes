git 切换远程代码库
>```bash
git remote rm origin #方法1
git remote set-url origin URL  #方法2
.git/config	#方法3
```

创建库
>```bash
touch README.md
git init
git add README.md
git commit -m "first commit"
git remote add origin https://github.com/warriorg/Estates.git
git push -u origin master
```

修改git仓库地址
>```bash
>git remote set-url origin http://114.55.148.240:10080/jinyi/bole-parking.git
>```

git 分支
>```bash
git checkout -b iss53
git push origin test:test  #提交本地test分支作为远程的test分支
```

关闭分支
>```bash
git checkout master
git merge <branch-name>
git branch -d <branch-name>
```

删除不在git管理下的文件
>```bash
git clean -nd #测试删除
git clean -fd #真实删除
```

导出不带版本的代码

>```bash
git archive master | bzip2 >source-tree.tar.bz2
#ZIP archive:
git archive --format zip --output /full/path/to/zipfile.zip master
#tag
git archive v1.0 gzip > source-tree.tgz
```

迁出远程分支

>```bash
>git checkout -b <branch-name> <origin/branch_name>
>```

提交tag到服务器

>```bash
>git push --tags  
>```
>*push a single tag*
>
>```
>git push origin <tag_name>
>```

git如何删除本地所有未提交的更改（都可以使用
>```bash
>git reset
>git checkout .
>git clean -fdx
>```
>```bash
> git checkout -- . #这条命令不会删除新增的文件
> git checkout -f #这条命令不会删除新增的文件
> git stash #新加的文件还在，但所有的修改都会抹去
> git add . && git stash && git stash drop #至少不会影响 .gitignore 里面的不跟踪的文件
>```

修改最后一次提交的注释
>```bas
git commit --amend
```

取消对文件的修改。还原到最近的版本，废弃本地做的修改
>```bash
>git checkout -- <file>
>```

解决冲突

>```bash
git mergetool -t diffmerge .
```

Submodule	
		
```bash
git submodule foreach git pull  #更新所有submodule
```

>安装diffmerge `osx`==brew cask install diffmerge==

取消已经暂存的文件。即，撤销先前"git add"的操作
`git reset HEAD <file>...`

回退所有内容到上一个版本
`git reset HEAD^`

回退a.py这个文件的版本到上一个版本  
`git reset HEAD^ a.py  `

向前回退到第3个版本  
`git reset –soft HEAD~3  `

将本地的状态回退到和远程的一样  
`git reset –hard origin/master `

回退到某个版本  
`git reset 057d `



回退到上一次提交的状态，按照某一次的commit完全反向的进行一次commit.(代码回滚到上个版本，并提交git)
>```bash
git revert HEAD
```

解决git目录过大

```base
git gc --prune=now  #运行 gc ，生成 pack 文件 --prune=now 表示对之前的所有提交做修剪，有的时候仅仅 gc 一下.git 文件就会小很多
git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -3    #找出最大的三个文件 
git rev-list --objects --all | grep c43a8da		#查看那些大文件究竟是谁（c43a8da 是上面大文件的hash码）
git filter-branch --force --index-filter "git rm --cached --ignore-unmatch 'data/bigfile'"  --prune-empty --tag-name-filter cat -- --all  #移除对该文件的引用（也就是 data/bigfile）

#进行 repack 
git for-each-ref --format='delete %(refname)' refs/original | git update-ref --stdin
git reflog expire --expire=now --all
git gc --prune=now

git count-objects -v   #查看 pack 的空间使用情况
```

方法二2

```bash
git rev-list --objects --all | grep "$(git verify-pack -v .git/objects/pack/*.idx | sort -k 3 -n | tail -5 | awk '{print$1}')" > large-files.txt
cat large-files.txt| awk '{print $2}' | tr '\n' ' '  >  large-files-inline.txt
git filter-branch -f --prune-empty --index-filter "git rm -rf --cached --ignore-unmatch `cat large-files-inline.txt`" --tag-name-filter cat -- --all
git push origin --force --all
```