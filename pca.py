import pandas as pd
from sklearn import decomposition
import matplotlib.pyplot as plt
import mpl_toolkits.mplot3d  # if not present "projection=3d" (see below) is crashing

pop = pd.read_json("sawo-export.json", lines=True)
pop = pop.drop("Victory", axis=1)
pop = pop.drop("Cost", axis=1)

print(pop)
X = pop

pca = decomposition.PCA(n_components=3)
pca.fit(X)
X = pca.transform(X)

print(X)

fig = plt.figure(1, figsize=(4, 3))
plt.clf()
ax = fig.add_subplot( projection="3d")

ax.scatter(X[:, 0], X[:, 1], X[:, 2])
plt.show()
