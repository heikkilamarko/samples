import duckdb
from duckdb.experimental.spark.sql import SparkSession as session
from duckdb.experimental.spark.sql.functions import col
import pandas as pd

all_rows = duckdb.sql("""
    SELECT * FROM 'products.json'
    """)

print("\nAll rows:")
all_rows.show()

filtered_rows = duckdb.sql("""
    SELECT * FROM all_rows WHERE id > 1
    """)

print("\nFiltered rows:")
filtered_rows.show()

filtered_df = filtered_rows.df()

spark = session.builder.getOrCreate()

df = spark.createDataFrame(filtered_df)

df = df.select(
    col('name'),
    col('description')
)

df.write.csv("out.csv", header=True)

pandas_df = pd.DataFrame({"a": [42]})
print("\nPandas DataFrame:")
duckdb.sql("SELECT * FROM pandas_df").show()
