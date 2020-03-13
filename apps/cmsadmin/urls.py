from django.urls import path
from . import views, user_view, meetplan_view

app_name = 'cmsadmin'

urlpatterns = [
    path('', views.noindex),
    path('index/', views.IndexView.as_view(), name="index"),
]

urlpatterns += [
    path('user_all/', user_view.UserView.as_view(), name='user_all'),
    path('user_create/', user_view.UserCreateView.as_view(), name='user_create'),
    path('user_create_many/', user_view.CreateManyUserView.as_view(), name='user_create_many'),
    path('user_update/<int:pk>/', user_view.UpdateUserView.as_view(), name='user_update'),
    path('user_deletelist/', user_view.DeletedUserListView.as_view(), name='user_delete_list'),
    path('user_recovery/<int:pk>/', user_view.RecoveryUserView.as_view(), name='user-recovery'),
    path('user_confirm_delete/<int:pk>/', user_view.UserDeleteView.as_view(), name='user-confirm-delete'),

    path('base_profile_all/', user_view.UserProfileListView.as_view(), name='base-profile-all'),
    path('base_profile_update/<int:pk>/', user_view.UserProfileUpdateView.as_view(), name='base-profile-update'),

    path('student_profile_all/', user_view.StudentProfileListView.as_view(), name='student-profile-all'),
    path('student_profile_update/<int:pk>/', user_view.StudentProfileUpdateView.as_view(), name='student-profile-update'),

    path('teacher_profile_all/', user_view.TeacherProfileListView.as_view(), name='teacher-profile-all'),
    path('teacher_profile_update/<int:pk>/', user_view.TeacherProfileUpdateView.as_view(), name='teacher-profile-update'),

]

urlpatterns += [
    path('meetplan_all/', meetplan_view.MeetPlanListView.as_view(), name='meetplan_all'),
    path('meetplan_create/', meetplan_view.MeetPlanCreateView.as_view(), name='meetplan_create'),
    path('meetplan_update/<int:pk>/', meetplan_view.MeetPlanUpdateView.as_view(), name='meetplan_update'),
    path('meetplan_confirm_delete/<int:pk>/', meetplan_view.MeetPlanDeleteView.as_view(),
         name='meetplan-confirm-delete'),

    path('meetplanorder_all/', meetplan_view.MeetPlanOrderListView.as_view(), name='meetplanorder_all'),
    path('meetplanorder_create/', meetplan_view.MeetPlanOrderCreateView.as_view(), name='meetplanorder_create'),
    path('meetplanorder_update/<int:pk>/', meetplan_view.MeetPlanOrderViewUpdate.as_view(),
         name='meetplanorder_update'),
    path('meetplanorder_confirm_delete/<int:pk>/', meetplan_view.MeetPlanOrderDeleteView.as_view(),
         name='meetplanorder-confirm-delete'),

    path('meetplanfeedback_all/', meetplan_view.FeedBackListView.as_view(), name='feedback_all'),
    path('meetplanfeedback_update/<int:pk>/', meetplan_view.FeedBackUpdateView.as_view(), name='feedback_update'),
]

urlpatterns += [
    path('term_date_setting/', meetplan_view.TermDateUpdateView.as_view(), name='meetplan-termdate-update')
]
