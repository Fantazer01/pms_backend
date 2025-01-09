import os

from sqlalchemy import (Column, Integer, MetaData, String, Table,
                        create_engine)

from databases import Database



DATABASE_URI = os.getenv('DATABASE_URI')

#DATABASE_URI = "postgresql://postgres:12345678@localhost/AI_db"

engine = create_engine(DATABASE_URI)
metadata = MetaData()

requests = Table(
    'requests',
    metadata,
    Column('id', Integer, primary_key=True),
    Column('request', String),
    Column('response', String)
)

database = Database(DATABASE_URI)