from flask import Blueprint,render_template,request,flash
import requests
from flask_sqlalchemy import SQLAlchemy
from impl import DB_NAME
from . import db
from impl.models import City

views=Blueprint('views',__name__)

@views.route('/',methods=['GET','POST'])
def home():
    error=""
    weather = {}
    if request.method=='POST':
        city=request.form.get('city')
        url="http://api.openweathermap.org/data/2.5/weather?q={}&units=metric&appid=64069b1194a5b83b9f1adb99ca57bf6d"

        r=requests.get(url.format(city)).json()
        if r["cod"] != 200:
            error = r["message"]
        else:
            weather = {
                'city' : city,
                'temperature':r['main']['temp'],
                'description':r['weather'][0]['description'],
                'icon':r['weather'][0]['icon'],
            }
    return render_template("home.html",weather=weather, error=error)

@views.route('/weather',methods=['GET','POST'])
def saved():
    weatherlist=[]
    if request.method=='POST':
        city=request.form.get('city')

        cname=City.query.filter_by(name=city).first()

        if cname:
            flash('Already added!',category='info')
        elif len(city)<=0:
            flash('Name should have atleast 1 character',category='error')
        else:
            new_city=City(name=city)
            db.session.add(new_city)
            db.session.commit()
            flash('City added successfully!',category='success')

            cities=City.query.all()
            url="http://api.openweathermap.org/data/2.5/weather?q={}&units=metric&appid=64069b1194a5b83b9f1adb99ca57bf6d"

            for cty in cities:
                r=requests.get(url.format(cty.name)).json()
                if r['cod'] == 200:
                    weather = {
                        'city' : cty.name,
                        'temperature':r['main']['temp'],
                        'description':r['weather'][0]['description'],
                        'icon':r['weather'][0]['icon'],
                    }
                    weatherlist.append(weather)
            return render_template("weather.html",weatherlist=weatherlist)

    else:
        cities=City.query.all()
        url="http://api.openweathermap.org/data/2.5/weather?q={}&units=metric&appid=64069b1194a5b83b9f1adb99ca57bf6d"

        for city in cities:
            r=requests.get(url.format(city.name)).json()
            if "main" in r:
                weather = {
                    'city' : city.name,
                    'temperature':r['main']['temp'],
                    'description':r['weather'][0]['description'],
                    'icon':r['weather'][0]['icon'],
                }
                weatherlist.append(weather)
    return render_template("weather.html",weatherlist=weatherlist, error="")
    
    