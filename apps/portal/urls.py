from django.urls import path
from . import views

app_name = 'portal'

urlpatterns = [
    path('', views.noindex),
    path('index/', views.IndexView.as_view(), name='index'),

    path('contact/', views.ContactView.as_view(), name='contact'),
    path('friend_link/', views.FriendLinkView.as_view(), name='friendlink'),
]
