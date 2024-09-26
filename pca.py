import pandas as pd
from sklearn import decomposition
import matplotlib.pyplot as plt
import mpl_toolkits.mplot3d  # if not present "projection=3d" (see below) is crashing

population = pd.read_json("sawo-export.json", lines=True)
population = population.drop("Cost", axis=1)

X = population

print(X)
pca = decomposition.PCA()
pca.fit(X)

print("Variance")
print(pca.explained_variance_ratio_)
print("Components")
print(pca.components_)

# https://towardsdatascience.com/pca-clearly-explained-how-when-why-to-use-it-and-feature-importance-a-guide-in-python-7c274582c37e