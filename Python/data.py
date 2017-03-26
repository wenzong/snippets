# 10 Minutes to pandas
# http://pandas.pydata.org/pandas-docs/stable/10min.html
#
import pandas as pd
import numpy as np

s = pd.Series([1, 3, 5, np.nan, 6, 8])
print(s)

dates = pd.date_range('20160801', periods=6)
print(dates)

df2 = pd.DataFrame({
    'A': 1.,
    'B': pd.Timestamp('20160826'),
    'C': pd.Series(1, index=list(range(4)), dtype='float32'),
    'D': np.array([3] * 4, dtype='int32'),
    'E': pd.Categorical(["test", "train", "test", "train"]),
    'F': 'foo'
})

print(df2)
print(df2.dtypes)
print(df2.tail(1))
