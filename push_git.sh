cd "/Users/sabath/workspace/mrsabath/mytechnotes"
git add *
git status
MSG=$(date +'%Y-%m-%d %H:%M:%S')
git commit  -m "Auto-committed at $MSG"
HOME=/Users/sabath git push origin master
