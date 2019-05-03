# commit hacker

## 核心原理
git add hacker.txt

git commit -m "xxx" --date=\"Thu Aug 9 

00:00:00 2018"

git push

创建一个文件

随便写入一些东西

计算时间并生成git需要的的时间串

然后add commit （调用shell的执行api）

commit后就可以通过git log测试

最后调用shell来push

核心步骤:读写文件，操作shell，算时间


## 开发过程中的坑
- time的Add要使用他的返回值
- exec.Command执行shell, 需要将shell语句需要按空格拆分
- git commit的--amend会导致多次提交只有一次生效, 所以commit的时候只加--date即可
- int转string需要用strconv.Itoa(num)
- time.ParseDuration() 这里的参数一定要有单位, 例如-24h

## 其他功能
30%的"摸鱼"成功率


70%的"勤奋"中42%的随机"超级勤奋"