cd "/Users/sabath/Box Sync/projects/mytechnotes"
git add *
git status
MSG=$(date +'%Y-%m-%d %H:%M:%S')
git commit  -m "Auto-committed at $MSG"
HOME=/Users/sabath git push origin master
