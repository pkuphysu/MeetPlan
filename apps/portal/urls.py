from django.urls import path, reverse
from . import views
from apps.account_auth.login.views import IAAALoginAuth

app_name = 'portal'

urlpatterns = [
    path('', views.noindex),
    path('index/', views.IndexView.as_view(), name='index'),

    path('contact/', views.ContactView.as_view(), name='contact'),
    path('friend_link/', views.FriendLinkView.as_view(), name='friendlink'),
]
