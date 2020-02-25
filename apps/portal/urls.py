from django.urls import path, reverse
from . import views
from apps.account_auth.login.views import IAAALoginAuth

app_name = 'portal'

urlpatterns = [
    # TODO: 向计算中心申请更改callback地址及注册IP
    path('', views.noindex),
    # path('', IAAALoginAuth.as_view()),

    path('index/', views.IndexView.as_view(), name='index'),

    path('about/', views.AboutView.as_view(), name='about'),
    path('contect/', views.ContactView.as_view(), name='contact'),
    path('recruitment/', views.RecruitmentView.as_view(), name='recruitment'),
    path('friend_link/', views.FriendLinkView.as_view(), name='friendview'),
]
