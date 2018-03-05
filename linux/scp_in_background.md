To execute any linux command in background we use nohup as follows:
```
$ nohup SOME_COMMAND &
```
But the problem with scp command is that it prompts for the password (if password authentication is used). So to make scp execute as a background process do this:
```
$ nohup scp file_to_copy user@server:/home > nohup.out 2>&1
```
Then press Mac ctrl + z which will temporarily suspend the command, then enter the command:
```
$ bg
```
This will start executing the command in backgroud:
test this:
```
ps -ef  | grep scp
# on the target:
watch ls -l /home
```
