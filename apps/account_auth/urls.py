from django.urls import path, include, re_path

from . import views

app_name = 'account_auth'

urlpatterns = [
    path('', views.noindex),
    path('index/', views.index, name="index"),
    path('login/', include('apps.account_auth.login.urls')),
    path('logout/', views.logout_view, name='logout'),
    re_path('^active/(?P<token>.*)$', views.ActiveView.as_view(), name='active-account'),

    path('baseprofile_add/', views.BaseProfileAddView.as_view(), name='baseprofile_add'),
    path('baseprofile_update/<int:pk>/', views.BaseProfileUpdateView.as_view(), name='baseprofile-update'),

    path('student_profile_create/', views.StudentProfileCreateView.as_view(), name='student-profile-create'),
    path('student_profile_update/<int:pk>/', views.StudentProfileUpdateView.as_view(), name='student-profile-update'),

    path('teacher_profile_create/', views.TeacherProfileCreateView.as_view(), name='teacher-profile-create'),
    path('teacher_profile_update/<int:pk>/', views.TeacherProfileUpdateView.as_view(), name='teacher-profile-update'),

    path('baseprofileimg_upload/', views.BaseProfileImgUpdateView.as_view(), name='baseprofileimg_add'),
    path('useremail_update/<int:pk>/', views.UserEmailUpdateView.as_view(), name='useremail_update'),

    path('student/<int:pk>/', views.StudentDetailView.as_view(), name='stu-detail'),
    path('teacher/<int:pk>/', views.TeacherDetailView.as_view(), name='tea-detail'),

]

urlpatterns += [
    path('ajax_student_profile_department/', views.LoadDepartmentView.as_view(), name='department-ajax'),
    path('ajax_student_profile_major/', views.LoadMajorView.as_view(), name='major_ajax'),
]
