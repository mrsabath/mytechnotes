import matplotlib.pyplot as plt
from matplotlib import animation

import pandas as pd
import numpy as np
import time


DRAIN_THRESH=4

SPEED_UP=50
def barlist(n):
    return df.iloc[n].tolist()[1:]


#DATA_FILE="data/simpopt_70load_75pod"
DATA_FILE="data/20170914.2.SIMOPT.csv"
import sys
if len(sys.argv) > 1:
    DATA_FILE=sys.argv[1]

df = pd.read_csv(DATA_FILE, index_col = 0)
df = df.astype(float)
df = df.reindex(pd.RangeIndex(start=min(df.index), stop=max(df.index), step=1)).interpolate()


x_axis = list(df.keys())
lenX = len(x_axis)



fig=plt.figure()

n=len(df)-1
x=range(1,lenX)
print barlist(1)
barcollection = plt.bar(x,barlist(1), align="center", alpha=0.5)

highest = np.max(df.values)
plt.ylim(0, highest * 1.1)
plt.ylabel("Number of pods")
plt.xlabel("Nodes")


colors = [ "blue" for i in x_axis]
def animate(k):
    y=barlist(k+1)
    drained = map(lambda x : x[0],filter(lambda x: x[1] <= DRAIN_THRESH, zip(range(lenX), y)))

    for n in drained:
        colors[n] = "green"

    for i, b in enumerate(barcollection):
        b.set_height(y[i])
        b.set_color(colors[i])
    plt.title(DATA_FILE + '\n' + 't=' + str(k))

anim=animation.FuncAnimation(fig,animate,repeat=False,blit=False,frames=n,
                             interval=1)

anim.save('mymovie.mp4',writer=animation.FFMpegWriter(fps=SPEED_UP))
#plt.show()
