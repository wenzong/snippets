# -*- coding: utf-8 -*-
from sqlalchemy import Column, DateTime, Integer, String
from sqlalchemy.sql import func
from sqlalchemy.ext.declarative import declarative_base


Base = declarative_base()


class MyTable(Base):

    __tablename__ = 'table_name'

    table_id = Column('id', Integer, primary_key=True)
    table_name = Column(String(255), nullable=False)
    created_at = Column(DateTime, server_default=func.now())
    updated_at = Column(DateTime, server_default=func.now(), onupdate=func.now())
