import pandas as pd
from sklearn import linear_model
import matplotlib.pyplot as plt
import mpl_toolkits.mplot3d  # noqa: F401

population = pd.read_json("sawo-export.json", lines=True)
rowCount = population.axes[0].stop

X = population.drop("Victory", axis=1).drop("Cost", axis=1)
Y = population["Victory"]
Y = Y / rowCount

print(X)
regr = linear_model.LinearRegression()
regr.fit(X, Y)

print("Coeff")
pct = regr.coef_ * 100

for idx, name in enumerate(X.columns):
    print(name, "\t% .1f%%" % pct[idx])


#print("Plot max victory")
fig = plt.figure(1, figsize=(4, 3))
plt.clf()

ax = fig.add_subplot(projection="3d")
for k, trait in enumerate(['STR','Fight','Block','TradW','CmbRef','NervSt']):
    subset = population.groupby([trait])
    #print(subset)
    stat = subset['Victory'].describe()
    #print(stat)
    for row in stat.iterrows():
        ax.scatter(k, row[0], row[1]['max'])
        #print(trait, row[0], row[1]['max'])

plt.show()