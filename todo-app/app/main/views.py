from django.http.response import HttpResponseRedirect
from django.shortcuts import redirect, render
from django.http import HttpResponse
from .models import ToDoList,Item
from .forms import CreateNewList

def index(response,id):
    ls=ToDoList.objects.get(id=id)

    if ls in response.user.todolist.all():
        if response.method=='POST':
            print(response.POST)
            if response.POST.get("save"):
                for item in ls.item_set.all():
                    if response.POST.get("c"+str(item.id)) == "clicked":
                        item.complete= True 
                    else:
                        item.complete=False
                    item.save()

            elif response.POST.get("newItem"):
                txt=response.POST.get("new")

                if len(txt)>2:
                    ls.item_set.create(text=txt,complete=False)
                else:
                    print("Invalid")
        return render(response,'main/list.html',{"ls":ls})

    return render(response, 'main/view.html',{})

def home(response):
    return render(response,'main/home.html',{})

def create(response):
    if response.method == 'POST':
        form=CreateNewList(response.POST)
        print(form.errors)
        if form.is_valid():
            n=form.cleaned_data['name']
            t=ToDoList(name=n)
            t.save()
            response.user.todolist.add(t)
        return HttpResponseRedirect("/%i" %t.id)
    else:
        form=CreateNewList()
    return render(response,'main/create.html',{"form":form})

def view(response):
    return render(response,'main/view.html',{})

def delete_item(response,pk):
    item=Item.objects.get(id=pk)
    if response.method=='POST':
        item.delete()
        obj=ToDoList.objects.get(name=item.todolist)
        
        return redirect('/%i' %obj.id)
    context={'item':item}
    return render(response,'main/delete.html',context)


def delete_todo(response,pk):
    todo=ToDoList.objects.get(id=pk)
    if response.method=='POST':
        todo.delete()
        return redirect('/view')
    context={'todo':todo}
    return render(response,'main/delete_todo.html',context)
