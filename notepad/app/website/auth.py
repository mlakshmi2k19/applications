from flask import Blueprint,render_template,request,flash,redirect,url_for

from website import db
from .models import User
from flask_login import login_user,login_required,logout_user,current_user 
from werkzeug.security import generate_password_hash,check_password_hash
auth=Blueprint('auth',__name__)

@auth.route('/login',methods=['GET','POST'])
def login():    
    if request.method=='POST':
        email=request.form.get('email')
        password=request.form.get('password')

        user=User.query.filter_by(email=email).first()
        if user:
            if check_password_hash(user.password,password):
                flash('Logged in successfully!',category='success')
                login_user(user,remember=True)
                return redirect(url_for('views.home'))
            else:
                flash('Incorrect password,try again.',category='error')
        else:
            flash('Email does not exists',category='error')
    return render_template("login.html",user=current_user)

@auth.route('/logout')
@login_required
def logout():
    logout_user()
    return redirect(url_for('auth.login'))

@auth.route('/signup',methods=['GET','POST'])
def signup():
    if request.method=='POST':
        email=request.form.get('email')
        firstName=request.form.get('firstName')
        password1=request.form.get('password1')
        password2=request.form.get('password2')
        user=User.query.filter_by(email=email).first()

        if user:
            flash('Email already exists.',category='error')
        elif len(email)<4:
            flash('Email should have greater than 3 chars.',category='error')
        elif len(firstName)<2:
            flash('Firstname should have greater than 1 chars.',category='error')
        elif len(password1)<7:
            flash('Password should have greater than 7 chars.',category='error')
        elif password1 != password2:
            flash('Passwords don\'t match.',category='error')
        else:
            new_user=User(email=email,firstName=firstName,password=generate_password_hash(password1,method='sha256'))
            db.session.add(new_user)
            db.session.commit()
            flash('Account created successfully!',category='success')
            login_user(user,remember=True)
            return redirect(url_for('views.home'))

    return render_template("signup.html",user=current_user)

