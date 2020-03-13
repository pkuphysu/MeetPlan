from django.urls import path, include, re_path
from . import views

app_name = 'account_auth'

urlpatterns = [
    path('', views.noindex),
    path('index/', views.index, name="index"),
    path('login/', include('apps.account_auth.login.urls')),
    path('logout/', views.logout_view, name='logout'),
    re_path('^active/(?P<token>.*)$', views.ActiveView.as_view(), name='active-account'),

    path('baseprofile_add/', views.UserProfileAddView.as_view(), name='userprofile_add'),

    path('student_profile_create_ajax/', views.LoadMajorView.as_view(), name='major_ajax'),
    path('student_profile_create/', views.StudentProfileCreateView.as_view(), name='student-profile-create'),
    path('student_profile_update/<int:pk>/', views.StudentProfileUpdateView.as_view(), name='student-profile-update'),

    path('teacher_profile_create/', views.TeacherProfileCreateView.as_view(), name='teacher-profile-create'),
    path('teacher_profile_update/<int:pk>/', views.TeacherProfileUpdateView.as_view(), name='teacher-profile-update'),

    path('userprofileimg_upload/', views.UserProfileImgUpdateView.as_view(), name='userprofileimg_add'),
    path('useremail_update/<int:pk>/', views.UserEmailUpdateView.as_view(), name='useremail_update'),

]
