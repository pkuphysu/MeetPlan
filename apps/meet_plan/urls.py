from django.urls import path

from . import views, tea_views, stu_views

app_name = 'meet_plan'

urlpatterns = [
    path('', views.NoIndexView.as_view()),
    path('tea_index/', views.TeaIndexView.as_view(), name='tea-index'),
    path('stu_index/', views.StuIndexView.as_view(), name='stu-index'),
]

# 学生相关
urlpatterns += [
    path('teacherlist/', stu_views.TeacherListView.as_view(), name='stu-all-tea-all-plan-list'),
    path('teacherplanlist/<int:pk>/', stu_views.TeacherPlanListView.as_view(), name='stu-tea-plan-list'),
    path('planordercreate/<int:pk>/', stu_views.MeetPlanOrderCreateView.as_view(), name='stu-order-create'),
]

# 教师相关
urlpatterns += [
    path('my_plan_all/', tea_views.MeetPlanListView.as_view(), name='tea-plan-list-all'),
    path('my_plan_create/', tea_views.MeetPlanCreateView.as_view(), name='tea-plan-add'),
    path('my_plan_create_fast/', tea_views.MeetPlanFastCreateView.as_view(), name='tea-plan-add-fast'),
    path('my_plan_update/<int:pk>/', tea_views.MeetPlanUpdateView.as_view(), name='tea-plan-update'),
    path('my_plan_detail/<int:pk>/', tea_views.MeetPlanDetailView.as_view(), name='tea-plan-detail'),
    path('my_plan_confirm_delete/<int:pk>/', tea_views.MeetPlanDeleteView.as_view(), name='tea-plan-confirm-delete'),

    path('my_planorder_confirm/<int:pk>/', tea_views.MeetPlanOrderUpdateView.as_view(), name='tea-order-update'),
    path('my_planorder_delete/<int:pk>/', tea_views.MeetPlanOrderDeleteView.as_view(), name='tea-order-delete'),
    path('my_planorder_add/', tea_views.TeacherAddMeetPlanOrderView.as_view(), name='tea-order-add'),

    path('myfeedbacklist/', tea_views.FeedBackListView.as_view(), name='tea-feedback-list'),
    path('myfeedbackadd/', tea_views.FeedBackCreateView.as_view(), name='tea-feedback-add'),
]
