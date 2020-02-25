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
]

urlpatterns += [
    path('meetplan_all/', meetplan_view.MeetPlanView.as_view(), name='meetplan_all'),
    path('meetplan_create/', meetplan_view.MeetPlanCreateView.as_view(), name='meetplan_create'),
    path('meetplan_update/<int:pk>/', meetplan_view.MeetPlanUpdateView.as_view(), name='meetplan_update'),
    path('meetplanorder_all/', meetplan_view.MeetPlanOrderView.as_view(), name='meetplanorder_all'),
    path('meetplanorder_create/', meetplan_view.MeetPlanOrderCreateView.as_view(), name='meetplanorder_create'),
    path('meetplanorder_update/<int:pk>/', meetplan_view.UpdateMeetPlanOrderView.as_view(), name='meetplanorder_update'),
    path('meetplanfeedback_all/', meetplan_view.FeedBackListView.as_view(), name='feedback_all'),
    path('meetplanfeedback_update/<int:pk>/', meetplan_view.FeedBackUpdateView.as_view(), name='feedback_update'),
]