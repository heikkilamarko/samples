# DuckDB Study

## Create virtual env

```bash
python3 -m venv .venv
```

## Activate virtual env

```bash
source .venv/bin/activate
```

## Deactivate virtual env

```bash
deactivate
```

## Install deps

```bash
pip install duckdb
pip install pandas
pip install typing-extensions
```

## Create `requirements.txt`

```bash
pip freeze > requirements.txt
```

## Install deps from `requirements.txt`

```bash
pip install -r requirements.txt
```

## Run the app

```bash
python app.py
```
