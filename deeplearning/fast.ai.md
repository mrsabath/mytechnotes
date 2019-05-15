http://course.fast.ai/lessons/lesson1.html
http://forums.fast.ai/t/wiki-lesson-1/9398

https://www.paperspace.com/terminal/pskr16oif
https://github.com/reshamas/fastai_deeplearn_part1/blob/master/tools/paperspace.md


Connect from localhost:
```
# create reverse SSL tunnel:
ssh  -N -L localhost:8889:localhost:8889  paperspace@184.105.238.238
# and another connection interactive:
ssh paperspace@184.105.238.238

# from the remote console, invoke
jupyter notebook --no-browser --port=8889 --NotebookApp.allow_remote_access=True

# get the Local URL (replace with `localhost`)
