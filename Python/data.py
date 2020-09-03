# 10 Minutes to pandas
# http://pandas.pydata.org/pandas-docs/stable/10min.html
#
import pandas as pd
import numpy as np


# Series
# one-dimensional labeled array
#
# * From Python dict
# * From ndarray
# * From scalar value
s = pd.Series(dict(a=1, b=2, c=3))
print(s)
s = pd.Series(dict(a=1, b=2, c=3), ['a', 'b', 'd'])
print(s)
s = pd.Series([1, 3, 5, np.nan, 6, 8])
print(s)
s = pd.Series([1, 3, 5, np.nan, 6, 8], index=[i for i in 'abcdef'])
print(s)
s = pd.Series([i for i in 'fedcba'], index=[i for i in 'abcdef'])
print(s)
s = pd.Series(5, index=[i for i in 'abcdef'])
print(s)

dates = pd.date_range('20160801', periods=6)
print(dates)

# DataFrame
# 2-dimensional labeled data structure with columns of potentially different
# types.
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


# TODO: implement octave: J = sum((X * theta - y) .^ 2) / 2 / m;
