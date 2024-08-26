CREATE TABLE quotes (
    id TEXT PRIMARY KEY,
    code TEXT,
    code_in TEXT,
    name TEXT,
    high REAL,
    low REAL,
    var_bid REAL,
    pct_change REAL,
    bid REAL,
    created_at DATETIME
)