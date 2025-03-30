from django.urls import path
from .import views


urlpatterns=[
    path("<int:id>",views.index,name="index"),
    path("",views.home,name="home"),
    path("home/",views.home,name="home"),
    path("create/",views.create,name="create"),
    path("view/",views.view,name="view"),
    path("delete_item/<str:pk>/",views.delete_item,name="delete_item"),
    path("delete_todo/<str:pk>/",views.delete_todo,name="delete_todo")
]