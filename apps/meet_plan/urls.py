from django.urls import path
from . import views, tea_views, stu_views

app_name = 'meet_plan'

urlpatterns = [
    path('', views.noindex),
    path('index/', views.IndexView.as_view(), name='index'),

]

# 学生相关
urlpatterns += [
    path('teacherlist/', stu_views.TeacherListView.as_view(), name='stu-all-tea-all-plan-list'),
    path('teacherplanlist/<int:tea_id>/', stu_views.TeacherPlanListView.as_view(), name='stu-tea-plan-list'),
    path('planordercreate/<int:mp_id>/', stu_views.MeetPlanOrderCreateView.as_view(), name='stu-order-create'),
]

# 教师相关
urlpatterns += [
    path('myplanall/', tea_views.MeetPlanListView.as_view(), name='tea-plan-list-all'),
    path('myplancreate/', tea_views.MeetPlanCreateView.as_view(), name='tea-plan-add'),
    path('myplanupdate/<int:pk>/', tea_views.MeetPlanUpdateView.as_view(), name='tea-plan-update'),
    path('myplandetail/<int:pk>/', tea_views.MeetPlanDetailView.as_view(), name='tea-plan-detail'),
    path('myplanorderconfirm/<int:pk>/', tea_views.MeetPlanOrderUpdateView.as_view(), name='tea-order-update'),
    path('myfeedbacklist/', tea_views.FeedBackListView.as_view(), name='tea-feedback-list'),
    path('myfeedbackadd/', tea_views.FeedBackCreateView.as_view(), name='tea-feedback-add'),
]
