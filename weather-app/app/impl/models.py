from enum import unique
from . import db

class City(db.Model):
    id=db.Column(db.Integer,primary_key=True)
    name=db.Column(db.String(150),unique=True,nullable=False)
