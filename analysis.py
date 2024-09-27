import pandas as pd
from sklearn import linear_model

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
