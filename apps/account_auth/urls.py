from django.urls import path, include, re_path
from . import views

app_name = 'account_auth'

urlpatterns = [
    path('', views.noindex),
    path('index/', views.index, name="index"),
    path('login/', include('apps.account_auth.login.urls')),
    path('logout/', views.logout_view, name='logout'),
    re_path('^active/(?P<token>.*)$', views.ActiveView.as_view(), name='active-account'),

    path('add_userprofile/', views.UserProfileAddView.as_view(), name='userprofile_add'),
    path('userprofile_update/<int:pk>/', views.UserProfileUpdateView.as_view(), name='userprofile_update'),
    path('userprofileimg_update/<int:pk>/', views.UserProfileImgUpdateView.as_view(), name='userprofileimg_add'),
    path('useremail_update/<int:pk>/', views.UserEmailUpdateView.as_view(), name='useremail_update'),

]
